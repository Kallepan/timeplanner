package repository

import (
	"context"
	"planner-backend/app/pkg"
	"strings"
	"time"

	"github.com/google/wire"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

var RepositorySet = wire.NewSet(
	departmentRepositorySet,
	workplaceRepositorySet,
	timeslotRepositorySet,
	weekdayRepositorySet,
	personRepositorySet,
	personRelRepositorySet,
	synchronizeRepositorySet,
	workdayRepositorySet,
)

func TimeDateToWeekdayID(t time.Time) string {
	return strings.ToUpper(t.Weekday().String()[0:3])
}

func EnsureDateExists(db *neo4j.DriverWithContext, ctx context.Context, date string) error {
	/**
	* Ensures that a date exists in the database
	* Is idempotent to ensure that the date is only created once
	* @param date: The date to ensure, Format: YYYY-MM-DD
	* @return: An error if the date could not be created
	 */

	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return err
	}

	weekday := TimeDateToWeekdayID(parsedDate)

	query := `
	// The weekday nodes are already created, so we can just match on them
	MATCH (w:Weekday {id: $weekdayID})
	// Create the date node if it doesn't exist yet
	MERGE (d:Date {date: date($date), week: date($date).week})
	// Create the relationship, if it doesn't exist yet
	MERGE (d) -[:IS_WEEKDAY]-> (w)
	RETURN d`
	params := map[string]interface{}{
		"weekdayID": weekday,
		"date":      date,
	}

	res, err := neo4j.ExecuteQuery(
		ctx,
		*db,
		query,
		params,
		neo4j.EagerResultTransformer,
	)
	if err != nil {
		return err
	}

	// Check if the date was created
	if len(res.Records) == 0 {
		return pkg.ErrNoRows
	}

	return err
}
