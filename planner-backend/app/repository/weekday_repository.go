package repository

import (
	"context"
	"planner-backend/app/domain/dao"

	"github.com/google/wire"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

/**
    * This file contains the WeekdayRepository interface and its implementation.
    * The interface is used to add, delete or update a weekday for a give timeslot.
    * The implementation is used to interact with the database.
**/

type WeekdayRepository interface {
	DeleteAllWeekdaysFromTimeslot(timeslot *dao.Timeslot) error
	AddWeekdaysToTimeslot(timeslot *dao.Timeslot, weekdays []dao.OnWeekday) ([]dao.OnWeekday, error)

	AddWeekdayToTimeslot(timeslot *dao.Timeslot, weekday *dao.OnWeekday) ([]dao.OnWeekday, error)
	DeleteWeekdayFromTimeslot(timeslot *dao.Timeslot, weekday *dao.OnWeekday) error
}

type WeekdayRepositoryImpl struct {
	db  *neo4j.DriverWithContext
	ctx context.Context
}

func (w WeekdayRepositoryImpl) AddWeekdaysToTimeslot(timeslot *dao.Timeslot, weekdays []dao.OnWeekday) ([]dao.OnWeekday, error) {
	/* Adds a list of weekdays to a timeslot */
	query := `
	MATCH (d:Department {name: $departmentName})-[:HAS_WORKPLACE]->(wp:Workplace {name: $workplaceName})-[:HAS_TIMESLOT]->(t:Timeslot {name: $timeslotName})
	UNWIND $weekdays AS weekday
	MATCH (wd:Weekday {id: weekday.id})
	MERGE (t)-[r:OFFERED_ON]->(wd)
	ON CREATE SET
		r.start_time = time(weekday.start_time),
		r.end_time = time(weekday.end_time)
	ON MATCH SET
		r.start_time = time(weekday.start_time),
		r.end_time = time(weekday.end_time)
	RETURN COLLECT({
		id: wd.id,
		name: wd.name,
		start_time: r.start_time,
		end_time: r.end_time
	}) AS weekdays`

	// convert weekdays to a list of maps
	weekdaysMap := []map[string]interface{}{}
	for _, weekday := range weekdays {
		weekdaysMap = append(weekdaysMap, weekday.ToMap())
	}

	params := map[string]interface{}{
		"departmentName": timeslot.DepartmentName,
		"workplaceName":  timeslot.WorkplaceName,
		"timeslotName":   timeslot.Name,
		"weekdays":       weekdaysMap,
	}

	result, err := neo4j.ExecuteQuery(
		w.ctx,
		*w.db,
		query,
		params,
		neo4j.EagerResultTransformer,
	)
	if err != nil {
		return nil, err
	}

	var weekdaysResponse []dao.OnWeekday
	// get the returned record
	record := result.Records[0]

	// get the weekdays collection
	weekdaysCollection, _, err := neo4j.GetRecordValue[[]any](record, "weekdays")
	if err != nil {
		return nil, err
	}

	// parse the collection into a list of weekdays
	for _, weekdayInterface := range weekdaysCollection {
		weekdayMap, ok := weekdayInterface.(map[string]interface{})
		if !ok {
			return nil, err
		}

		weekday := dao.OnWeekday{}
		if err := weekday.ParseFromMap(weekdayMap); err != nil {
			return nil, err
		}

		weekdaysResponse = append(weekdaysResponse, weekday)
	}

	return weekdaysResponse, nil
}

func (w WeekdayRepositoryImpl) DeleteAllWeekdaysFromTimeslot(timeslot *dao.Timeslot) error {
	/* Deletes all weekdays from a timeslot */
	query := `
	MATCH (d:Department {name: $departmentName})-[:HAS_WORKPLACE]->(wp:Workplace {name: $workplaceName})-[:HAS_TIMESLOT]->(t:Timeslot {name: $timeslotName})
	MATCH (t)-[r:OFFERED_ON]->(wd:Weekday)
	DELETE r
	`

	params := map[string]interface{}{
		"departmentName": timeslot.DepartmentName,
		"workplaceName":  timeslot.WorkplaceName,
		"timeslotName":   timeslot.Name,
	}

	_, err := neo4j.ExecuteQuery(
		w.ctx,
		*w.db,
		query,
		params,
		neo4j.EagerResultTransformer,
	)

	if err != nil {
		return err
	}

	return nil
}

func (w WeekdayRepositoryImpl) AddWeekdayToTimeslot(timeslot *dao.Timeslot, weekday *dao.OnWeekday) ([]dao.OnWeekday, error) {
	/* Adds a weekday to a timeslot */

	query := `
    MATCH (d:Department {name: $departmentName})-[:HAS_WORKPLACE]->(wp:Workplace {name: $workplaceName})-[:HAS_TIMESLOT]->(t:Timeslot {name: $timeslotName})
    MATCH (wd:Weekday {id: $weekdayID})
    MERGE (t)-[r:OFFERED_ON]->(wd)
    ON CREATE SET 
		r.start_time = time($startTime), 
		r.end_time = time($endTime)
    ON MATCH SET 
		r.start_time = time($startTime), 
		r.end_time = time($endTime)
    WITH t
    MATCH (t)-[r:OFFERED_ON]->(wd:Weekday)
    RETURN COLLECT({
		id: wd.id,
		name: wd.name,
		start_time: r.start_time,
		end_time: r.end_time
	}) AS weekdays`
	params := map[string]interface{}{
		"departmentName": timeslot.DepartmentName,
		"workplaceName":  timeslot.WorkplaceName,
		"timeslotName":   timeslot.Name,
		"weekdayID":      weekday.ID,
		"startTime":      weekday.StartTime,
		"endTime":        weekday.EndTime,
	}

	result, err := neo4j.ExecuteQuery(
		w.ctx,
		*w.db,
		query,
		params,
		neo4j.EagerResultTransformer,
	)
	if err != nil {
		return nil, err
	}

	var weekdays []dao.OnWeekday
	// get the returned record
	record := result.Records[0]

	// get the weekdays collection
	weekdaysCollection, _, err := neo4j.GetRecordValue[[]any](record, "weekdays")
	if err != nil {
		return nil, err
	}

	// parse the collection into a list of weekdays
	for _, weekdayInterface := range weekdaysCollection {
		weekdayMap, ok := weekdayInterface.(map[string]interface{})
		if !ok {
			return nil, err
		}

		weekday := dao.OnWeekday{}
		if err := weekday.ParseFromMap(weekdayMap); err != nil {
			return nil, err
		}

		weekdays = append(weekdays, weekday)
	}

	return weekdays, nil
}

func (w WeekdayRepositoryImpl) DeleteWeekdayFromTimeslot(timeslot *dao.Timeslot, weekday *dao.OnWeekday) error {
	/* Deletes a weekday from a timeslot */

	query := `
    MATCH (d:Department {name: $departmentName})-[:HAS_WORKPLACE]->(wp:Workplace {name: $workplaceName})-[:HAS_TIMESLOT]->(t:Timeslot {name: $timeslotName})
    MATCH (wd:Weekday {id: $weekdayID})
    MATCH (t)-[r:OFFERED_ON]->(wd)
    DELETE r
    `
	params := map[string]interface{}{
		"departmentName": timeslot.DepartmentName,
		"workplaceName":  timeslot.WorkplaceName,
		"timeslotName":   timeslot.Name,
		"weekdayID":      weekday.ID,
	}

	_, err := neo4j.ExecuteQuery(
		w.ctx,
		*w.db,
		query,
		params,
		neo4j.EagerResultTransformer,
	)
	if err != nil {
		return err
	}

	return nil
}

/*
func (w WeekdayRepositoryImpl) UpdateWeekdayForTimeslot(timeslot *dao.Timeslot, weekday *dao.OnWeekday) ([]dao.OnWeekday, error) {

		query := `
	    MATCH (d:Department {name: $departmentName})-[:HAS_WORKPLACE]->(wp:Workplace {name: $workplaceName})-[:HAS_TIMESLOT]->(t:Timeslot {name: $timeslotName})
	    MATCH (wd:Weekday {id: $weekdayID})
	    MATCH (t)-[r:OFFERED_ON]->(wd)
	    SET r.start_time = time($startTime), r.end_time = time($endTime)
	    WITH t
	    MATCH (t)-[r:OFFERED_ON]->(wd:Weekday)
	    RETURN wd
	    `
		params := map[string]interface{}{
			"departmentName": timeslot.DepartmentName,
			"workplaceName":  timeslot.WorkplaceName,
			"timeslotName":   timeslot.Name,
			"weekdayID":      weekday.ID,
			"startTime":      weekday.StartTime,
			"endTime":        weekday.EndTime,
		}

		result, err := neo4j.ExecuteQuery(
			w.ctx,
			*w.db,
			query,
			params,
			neo4j.EagerResultTransformer,
		)
		if err != nil {
			return nil, err
		}

		weekdays := []dao.OnWeekday{}
		for _, record := range result.Records {
			weekday := dao.OnWeekday{}
			if err := weekday.ParseFromDB(record); err != nil {
				return nil, err
			}

			weekdays = append(weekdays, weekday)
		}

		return weekdays, nil
	}

*/

func WeekdayRepositoryInit(db *neo4j.DriverWithContext, ctx context.Context) *WeekdayRepositoryImpl {
	return &WeekdayRepositoryImpl{
		db:  db,
		ctx: ctx,
	}
}

var weekdayRepositorySet = wire.NewSet(
	WeekdayRepositoryInit,
	wire.Bind(new(WeekdayRepository), new(*WeekdayRepositoryImpl)),
)
