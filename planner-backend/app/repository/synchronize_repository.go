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
	Synchronize(weeksInAdvance int) error

	createWorkday(ctx context.Context, date string, weekday string) error
}

type SynchronizeRepositoryImpl struct {
	db  *neo4j.DriverWithContext
	ctx context.Context
}

func (d SynchronizeRepositoryImpl) Synchronize(weeksInAdvance int) error {
	/*
		Synchronize:
			- Get monday of the current week
			- calculate all dates from monday to sunday * weeksInAdvance
	*/

	// Get the current date
	now := time.Now()

	// Get the monday of the current week
	monday := now.AddDate(0, 0, -int(now.Weekday())+1)

	for i := 0; i < weeksInAdvance; i++ {
		for j := 0; j < 7; j++ {
			date := monday.AddDate(0, 0, 7*i+j)
			dateStr := date.Format("2006-01-02")
			weekday := TimeDateToWeekdayID(date)

			// Create the date node
			if err := EnsureDateExists(d.db, d.ctx, dateStr); err != nil {
				return err
			}

			// Create the workday nodes
			if err := d.createWorkday(d.ctx, dateStr, weekday); err != nil {
				return fmt.Errorf("error creating workday nodes: %w", err)
			}

			slog.Info(fmt.Sprintf("Synchronized date %s", dateStr))
		}
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
	WHERE t.deleted_at IS NULL AND d.deleted_at IS NULL AND w.deleted_at IS NULL
	
	WITH COLLECT({workplaceID:w.id, departmentID: d.id, timeslot: t, start_time: r.start_time, end_time: r.end_time}) AS collection
	UNWIND collection AS c
	WITH c

	// Date should already exist by now (created in the createDate function)
	MATCH (d:Date {date: date($date), week: date($date).week})
	MATCH (d) -[:IS_ON_WEEKDAY]-> (wd:Weekday {id: $weekdayID})

	MERGE (wkd:Workday {date: date($date), department: c.departmentID, workplace: c.workplaceID, timeslot: c.timeslot.id, weekday: wd.id})
	ON CREATE SET
		wkd.start_time = c.start_time,
		wkd.end_time = c.end_time,
		wkd.duration_in_minutes = duration.between(c.start_time, c.end_time).minutes,
		wkd.active = true,
		wkd.comment = "",
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

func SynchronizeRepositoryInit(db *neo4j.DriverWithContext, ctx context.Context) *SynchronizeRepositoryImpl {
	return &SynchronizeRepositoryImpl{
		db:  db,
		ctx: ctx,
	}
}

var synchronizeRepositorySet = wire.NewSet(
	SynchronizeRepositoryInit,
	wire.Bind(new(SynchronizeRepository), new(*SynchronizeRepositoryImpl)),
)
