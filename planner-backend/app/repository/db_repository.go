package repository

import (
	"context"
	"log/slog"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

/* Migrations */
var queries = map[string]string{
	// TODO: Implement a better way to handle these queries
	"unique_department":        `CREATE CONSTRAINT unique_department_id IF NOT EXISTS FOR (d:Department) REQUIRE d.id IS UNIQUE;`,
	"unique_person":            `CREATE CONSTRAINT unique_person_id IF NOT EXISTS FOR (p:Person) REQUIRE p.id IS UNIQUE;`,
	"create_monday":            `MERGE (:Weekday {name: 'Monday', id: "MON"});`,
	"create_tuesday":           `MERGE (:Weekday {name: 'Tuesday', id: "TUE"});`,
	"create_wednesday":         `MERGE (:Weekday {name: 'Wednesday', id: "WED"});`,
	"create_thursday":          `MERGE (:Weekday {name: 'Thursday', id: "THU"});`,
	"create_friday":            `MERGE (:Weekday {name: 'Friday', id: "FRI"});`,
	"create_saturday":          `MERGE (:Weekday {name: 'Saturday', id: "SAT"});`,
	"create_sunday":            `MERGE (:Weekday {name: 'Sunday', id: "SUN"});`,
	"index_workday_date":       `CREATE INDEX workday_date IF NOT EXISTS FOR (w:Workday) ON (w.date);`,
	"index_workday_department": `CREATE INDEX workday_department IF NOT EXISTS FOR (w:Workday) ON (w.department);`,
	"index_workday_workplace":  `CREATE INDEX workday_workplace IF NOT EXISTS FOR (w:Workday) ON (w.workplace);`,
	"index_workday_timeslot":   `CREATE INDEX workday_timeslot IF NOT EXISTS FOR (w:Workday) ON (w.timeslot);`,
}

func Migrate(ctx context.Context, db *neo4j.DriverWithContext) {
	slog.Info("Migrating database")

	for name, query := range queries {
		slog.Info("Running query", "name", name)

		if _, err := neo4j.ExecuteQuery(
			ctx,
			*db,
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

func Clear(ctx context.Context, db *neo4j.DriverWithContext) {
	/**
	* Clears the database
	 */
	slog.Info("Clearing database")

	if db == nil {
		slog.Warn("Database is nil. Not continuing...")
		return
	}

	if _, err := neo4j.ExecuteQuery(
		ctx,
		*db,
		"MATCH (n) DETACH DELETE n;",
		map[string]interface{}{},
		neo4j.EagerResultTransformer,
	); err != nil {
		slog.Error("Failed to clear database", "error", err)
		panic(err)
	}

	slog.Info("Database cleared")
}
