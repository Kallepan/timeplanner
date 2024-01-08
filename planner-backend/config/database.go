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

	// Migrate
	Migrate(ctx, driver)

	return &driver
}

/* Migrations */
var queries = map[string]string{
	// TODO: Implement a better way to handle these queries
	"unique_department": `CREATE CONSTRAINT unique_department_name IF NOT EXISTS FOR (d:Department) REQUIRE d.name IS UNIQUE;`,
	"unique_person":     `CREATE CONSTRAINT unique_person_id IF NOT EXISTS FOR (p:Person) REQUIRE p.id IS UNIQUE;`,
	"create_monday":     `MERGE (:Weekday {name: 'Monday', id: "MON"});`,
	"create_tuesday":    `MERGE (:Weekday {name: 'Tuesday', id: "TUE"});`,
	"create_wednesday":  `MERGE (:Weekday {name: 'Wednesday', id: "WED"});`,
	"create_thursday":   `MERGE (:Weekday {name: 'Thursday', id: "THU"});`,
	"create_friday":     `MERGE (:Weekday {name: 'Friday', id: "FRI"});`,
	"create_saturday":   `MERGE (:Weekday {name: 'Saturday', id: "SAT"});`,
	"create_sunday":     `MERGE (:Weekday {name: 'Sunday', id: "SUN"});`,
}

func Migrate(ctx context.Context, db neo4j.DriverWithContext) {
	slog.Info("Migrating database")

	for name, query := range queries {
		slog.Info("Running query", "name", name)

		if _, err := neo4j.ExecuteQuery(
			ctx,
			db,
			query,
			map[string]interface{}{},
			neo4j.EagerResultTransformer,
		); err != nil {
			slog.Error("Failed to run query", "name", name, "error", err)
			panic(err)
		}

	}

	slog.Info("Migration complete")
}

func Clear(ctx context.Context, db neo4j.DriverWithContext) {
	/**
	* Clears the database
	 */

	slog.Info("Clearing database")

	if db == nil {
		slog.Warn("Database is nil")
		return
	}

	if _, err := neo4j.ExecuteQuery(
		ctx,
		db,
		"MATCH (n) DETACH DELETE n;",
		map[string]interface{}{},
		neo4j.EagerResultTransformer,
	); err != nil {
		slog.Error("Failed to clear database", "error", err)
		panic(err)
	}

	slog.Info("Database cleared")
}
