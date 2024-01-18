package repository

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/wire"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type SynchronizeRepository interface {
	Synchronize(datesInAdvance int) error

	createWorkday(ctx context.Context, date string, weekday string) error
}

type SynchronizeRepositoryImpl struct {
	db *neo4j.DriverWithContext
}

func (d SynchronizeRepositoryImpl) Synchronize(datesInAdvance int) error {
	/*
		Synchronize:
			- Get all dates from now to 1 week from now
			- Create Date nodes for each of them
			- Create Workday nodes for each of them
	*/
	ctx := context.Background()
	now := time.Now()

	// Get all dates from now to 1 week from now
	for n := now; n.Before(now.AddDate(0, 0, datesInAdvance)); n = n.AddDate(0, 0, 1) {
		date := n.Format("2006-01-02")
		weekday := TimeDateToWeekdayID(n)

		if err := EnsureDateExists(d.db, ctx, date); err != nil {
			return err
		}

		if err := d.createWorkday(ctx, date, weekday); err != nil {
			return err
		}

		slog.Info(fmt.Sprintf("Synchronized date %s", date))
	}
	return nil
}

func (d SynchronizeRepositoryImpl) createWorkday(ctx context.Context, date string, weekdayID string) error {
	/**
	 * Create Workday Nodes for Given Weekday and Date
	 *
	 * This Cypher query retrieves all timeslots offered on a given weekday, loops through them,
	 * and creates a Workday node for each of them. It assumes that the date and week nodes for
	 * the specified date already exist.
	 *
	 * @param {string} $weekdayID - The ID of the target weekday.
	 * @param {string} $date - The target date in "YYYY-MM-DD" format.
	 *
	 * Query Steps:
	 * 1. Matches departments having workplaces with associated ACTIVE timeslots offered on the specified weekday.
	 * 2. Collects relevant information about workplaces, departments, timeslots, and time details.
	 * 3. Unwinds the collection for further processing.
	 * 4. Matches the existing Date node for the specified date and week.
	 * 5. Creates Workday nodes for each collected data, setting properties on node creation.
	 * 6. Creates relationships between Workday nodes and Timeslot, Date nodes.
	 *
	 * Example Usage:
	 * CALL yourProcedureName($weekdayID, $date)
	 */
	query := `
	// Get all timeslots offered on the given weekday, loop through them and create a Workday node for each of them.
	MATCH  (d:Department) -[:HAS_WORKPLACE]-> (w:Workplace) -[:HAS_TIMESLOT]-> (t:Timeslot) -[r:OFFERED_ON]-> (wd:Weekday {id: $weekdayID})
	WHERE t.deleted_at IS NULL AND t.active = true

	WITH COLLECT({workplace:w.name, department: d.name, timeslot: t, start_time: r.start_time, end_time: r.end_time}) AS collection
	UNWIND collection AS c
	WITH c

	// Date should already exist by now (created in the createDate function)
	MATCH (d:Date {date: date($date), week: date($date).week})

	MERGE (wkd:Workday {date: date($date), department: c.department, workplace: c.workplace, timeslot: c.timeslot.name})
	ON CREATE SET
		wkd.start_time = c.start_time,
		wkd.end_time = c.end_time,
		wkd.active = true,
		wkd.created_at = datetime()
		
	// Create the relationships
	WITH wkd, c.timeslot AS t, d
	MERGE (wkd) -[:IS_TIMESLOT]-> (t)
	MERGE (wkd) -[:IS_DATE]-> (d)
	`

	params := map[string]interface{}{
		"date":      date,
		"weekdayID": weekdayID,
	}

	_, err := neo4j.ExecuteQuery(
		ctx,
		*d.db,
		query,
		params,
		neo4j.EagerResultTransformer,
	)
	return err
}

func SynchronizeRepositoryInit(db *neo4j.DriverWithContext) *SynchronizeRepositoryImpl {
	return &SynchronizeRepositoryImpl{
		db: db,
	}
}

var synchronizeRepositorySet = wire.NewSet(
	SynchronizeRepositoryInit,
	wire.Bind(new(SynchronizeRepository), new(*SynchronizeRepositoryImpl)),
)
