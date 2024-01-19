/**
* Test using go-testcontainers to test the repository using a mock database
**/
package repository

import (
	"context"
	"fmt"
	"testing"

	"github.com/docker/go-connections/nat"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func SetupTestDB(ctx context.Context, t *testing.T) (*neo4j.DriverWithContext, error) {
	req := testcontainers.ContainerRequest{
		Image: "neo4j:latest",
		Env: map[string]string{
			"NEO4J_AUTH": "neo4j/test",
		},
		WaitingFor:   wait.ForLog("Started."),
		ExposedPorts: []string{"7687/tcp"},
	}

	neo4jContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		t.Fatal(err)
	}

	// Terminate the container when the context is done
	go func() {
		<-ctx.Done()
		if err := neo4jContainer.Terminate(ctx); err != nil {
			panic(err)
		}
	}()

	neo4jPort, err := neo4jContainer.MappedPort(ctx, nat.Port("7687"))
	if err != nil {
		return nil, err
	}

	dbSN := fmt.Sprintf("bolt://localhost:%s", neo4jPort.Port())

	db, err := neo4j.NewDriverWithContext(dbSN, neo4j.BasicAuth("neo4j", "test", ""))
	if err != nil {
		return nil, err
	}

	return &db, nil
}
