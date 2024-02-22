package repository

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/docker/go-connections/nat"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func setupTestDB(ctx context.Context) (*neo4j.DriverWithContext, error) {
	/**
	* Function to start a neo4j container for testing
	* This function creates an isolated neo4j database which can be accessed by the test cases
	* The function returns a neo4j driver which can be used to interact with the database
	**/
	req := testcontainers.ContainerRequest{
		Image: "neo4j:latest",
		Env: map[string]string{
			"NEO4J_AUTH": "neo4j/test",
			"NEO4J_dbms_security_auth__minimum__password__length": "1",
		},
		WaitingFor:   wait.ForLog("Started."),
		ExposedPorts: []string{"7687/tcp"},
	}

	neo4jContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, err
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

// Setup a database to be used by all sub tests
func TestMain(m *testing.M) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db, err := setupTestDB(ctx)
	if err != nil {

		panic("Test database setup failed")
	}

	Migrate(ctx, db)

	os.Exit(m.Run())
}
