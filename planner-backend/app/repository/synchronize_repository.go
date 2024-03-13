package repository

import (
	"context"
	"fmt"
	"log/slog"
	"planner-backend/app/pkg"
	"time"

	"github.com/google/wire"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type SynchronizeRepository interface {
	Synchronize(weeksInAdvance int) error

	createWorkday(tx neo4j.ManagedTransaction, date string, weekday int64) error
	ensureDateExists(tx neo4j.ManagedTransaction, date string, weekdayID int64) error
}

type SynchronizeRepositoryImpl struct {
	db  *neo4j.DriverWithContext
	ctx context.Context
}

func (d SynchronizeRepositoryImpl) ensureDateExists(tx neo4j.ManagedTransaction, date string, weekdayID int64) error {
	/*
		* Ensures that a date exists in the database
		This function is used during the synchronization process to ensure that a date exists
		* @param tx: The transaction to use
		* @param date: The date to ensure, Format: YYYY-MM-DD
		* @param weekday: The weekday to ensure, Format: 1-7
		* @return: An error if the date could not be created
	*/
	slog.Info(fmt.Sprintf("Ensuring date %s exists", date))
	query := `
	// The weekday nodes are already created, so we can just match on them
	MATCH (w:Weekday {id: $weekdayID})
	// Create the date node if it doesn't exist yet
	MERGE (d:Date {date: date($date), week: date($date).week})
	// Create the relationship, if it doesn't exist yet
	MERGE (d) -[:IS_ON_WEEKDAY]-> (w)
	RETURN d`
	params := map[string]interface{}{
		"weekdayID": weekdayID,
		"date":      date,
	}

	res, err := tx.Run(
		d.ctx,
		query,
		params,
	)
	if err != nil {
		return err
	}

	// Check if the date was created
	if !res.Next(d.ctx) {
		return pkg.ErrNoRows
	}

	return nil
}

func (d SynchronizeRepositoryImpl) Synchronize(weeksInAdvance int) error {
	/*
	*	Synchronize:
	*	- Get monday of the current week
	*	- calculate all dates from monday to sunday * weeksInAdvance
	 */

	// Create Workday nodes for each date and weekday
	session := (*d.db).NewSession(d.ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(d.ctx)

	// Start a new transaction
	if _, err := session.ExecuteWrite(d.ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		// Get the current date
		now := time.Now()

		// Get the monday of the current week
		monday := now.AddDate(0, 0, -int(now.Weekday())+1)

		for i := 0; i < weeksInAdvance; i++ {
			for j := 0; j < 7; j++ {
				date := monday.AddDate(0, 0, 7*i+j)
				dateStr := date.Format("2006-01-02")
				weekdayID := TimeDateToWeekdayID(date)

				// Create the date node
				if err := d.ensureDateExists(tx, dateStr, weekdayID); err != nil {
					return nil, err
				}

				if err := d.createWorkday(tx, dateStr, weekdayID); err != nil {
					return nil, err
				}
			}
		}

		return nil, nil
	}); err != nil {
		return err
	}

	return nil
}

func (d SynchronizeRepositoryImpl) createWorkday(tx neo4j.ManagedTransaction, date string, weekdayID int64) error {
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
	slog.Info(fmt.Sprintf("Creating workday nodes for date %s and weekday %d", date, weekdayID))
	query := `
	// Get all timeslots offered on the given weekday, loop through them and create a Workday node for each of them.
	MATCH  (d:Department) -[:HAS_WORKPLACE]-> (w:Workplace) -[:HAS_TIMESLOT]-> (t:Timeslot) -[r:OFFERED_ON]-> (wd:Weekday {id: $weekdayID})
	MATCH (d2:Date {date: date($date), week: date($date).week}) -[:IS_ON_WEEKDAY]-> (wd {id: $weekdayID})
	WHERE t.deleted_at IS NULL AND w.deleted_at IS NULL AND d.deleted_at IS NULL
	AND NOT (d) -[:SYNCHRONIZED_AT]-> (:Date {date: date($date)})

	// Mark as synchronized
	MERGE (d) -[s:SYNCHRONIZED_AT]-> (d2)
	ON CREATE SET s.created_at = datetime()
	ON MATCH SET s.updated_at = datetime() // This should not happen, but just in case
	WITH d, w, t, r, d2
	
	// Collect relevant information about workplaces, departments, timeslots, and time details
	WITH COLLECT({workplaceID: w.id, departmentID: d.id, timeslot: t, start_time: r.start_time, end_time: r.end_time, date: d2}) AS collection
	UNWIND collection AS c
	WITH c

	// important: workday nodes should be unique for each date, department, workplace, and timeslot
	MERGE (wkd:Workday {date: date($date), department: c.departmentID, workplace: c.workplaceID, timeslot: c.timeslot.id, weekday: $weekdayID})
	ON CREATE SET
		wkd.start_time = c.start_time,
		wkd.end_time = c.end_time,
		wkd.duration_in_minutes = duration.between(c.start_time, c.end_time).minutes,
		// set active in here to avoid the merge query not matching the node
		wkd.active = true,
		wkd.comment = "",
		wkd.created_at = datetime()
	// Create the relationships
	WITH wkd, c.timeslot AS t, c.date AS d2
	MERGE (wkd) -[:IS_TIMESLOT]-> (t)
	MERGE (wkd) -[:IS_DATE]-> (d2)
	RETURN wkd
	`

	params := map[string]interface{}{
		"date":      date,
		"weekdayID": weekdayID,
	}

	result, err := tx.Run(
		d.ctx,
		query,
		params,
	)
	if err != nil {
		return err
	}

	// Check if the result is empty
	if !result.Next(d.ctx) || result.Err() != nil {
		slog.Warn(fmt.Sprintf("no workday nodes were created for date %s and weekday %d", date, weekdayID))
		return nil
	}

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
