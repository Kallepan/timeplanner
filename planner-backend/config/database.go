package config

import (
	"context"
	"log/slog"
	"os"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func ConnectToDB(ctx context.Context) *neo4j.DriverWithContext {
	uri := os.Getenv("PLANNER_DB_URI")
	username := os.Getenv("PLANNER_DB_USERNAME")
	password := os.Getenv("PLANNER_DB_PASSWORD")

	driver, err := neo4j.NewDriverWithContext(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		slog.Error("Failed to connect to database", "error", err)
		panic(err)
	}

	// Verify connectivity
	if err := driver.VerifyConnectivity(ctx); err != nil {
		slog.Error("Failed to connect to database", "error", err)
		panic(err)
	}

	// Close driver on exit
	go func() {
		<-ctx.Done()
		driver.Close(ctx)
	}()

	return &driver
}
