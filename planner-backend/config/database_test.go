package config

import (
	"context"
	"os"
	"testing"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func TestMigrate(t *testing.T) {
	ctx := context.Background()

	uri := os.Getenv("PLANNER_DB_URI")
	username := os.Getenv("PLANNER_DB_USERNAME")
	password := os.Getenv("PLANNER_DB_PASSWORD")

	d, err := neo4j.NewDriverWithContext(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		t.Fatal(err)
	}

	defer d.Close(ctx)

	Migrate(ctx, d)
	Clear(ctx, d)
}
