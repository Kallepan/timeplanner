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
	AddWeekdayToTimeslot(timeslot *dao.Timeslot, weekday *dao.Weekday) ([]dao.Weekday, error)
	DeleteWeekdayFromTimeslot(timeslot *dao.Timeslot, weekday *dao.Weekday) error
	UpdateWeekdayForTimeslot(timeslot *dao.Timeslot, weekday *dao.Weekday) ([]dao.Weekday, error)
}

type WeekdayRepositoryImpl struct {
	db *neo4j.DriverWithContext
}

func (w WeekdayRepositoryImpl) AddWeekdayToTimeslot(timeslot *dao.Timeslot, weekday *dao.Weekday) ([]dao.Weekday, error) {
	/* Adds a weekday to a timeslot */
	ctx := context.Background()
	query := `
    MATCH (d:Department {name: $departmentName})-[:HAS_WORKPLACE]->(wp:Workplace {name: $workplaceName})-[:HAS_TIMESLOT]->(t:Timeslot {name: $timeslotName})
    MATCH (wd:Weekday {id: $weekdayID})
    MERGE (t)-[r:OFFERED_ON]->(wd)
    ON CREATE SET r.start_time = time($startTime), r.end_time = time($endTime)
    ON MATCH SET r.start_time = time($startTime), r.end_time = time($endTime)
    WITH t
    MATCH (t)-[r:OFFERED_ON]->(wd:Weekday)
    RETURN wd`
	params := map[string]interface{}{
		"departmentName": timeslot.DepartmentName,
		"workplaceName":  timeslot.WorkplaceName,
		"timeslotName":   timeslot.Name,
		"weekdayID":      weekday.ID,
		"startTime":      weekday.StartTime,
		"endTime":        weekday.EndTime,
	}

	result, err := neo4j.ExecuteQuery(
		ctx,
		*w.db,
		query,
		params,
		neo4j.EagerResultTransformer,
	)
	if err != nil {
		return nil, err
	}

	weekdays := []dao.Weekday{}
	for _, record := range result.Records {
		weekday := dao.Weekday{}
		if err := weekday.ParseFromDB(record); err != nil {
			return nil, err
		}

		weekdays = append(weekdays, weekday)
	}

	return weekdays, nil
}

func (w WeekdayRepositoryImpl) DeleteWeekdayFromTimeslot(timeslot *dao.Timeslot, weekday *dao.Weekday) error {
	/* Deletes a weekday from a timeslot */
	ctx := context.Background()
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
		ctx,
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

func (w WeekdayRepositoryImpl) UpdateWeekdayForTimeslot(timeslot *dao.Timeslot, weekday *dao.Weekday) ([]dao.Weekday, error) {
	/* Updates a weekday for a timeslot */
	ctx := context.Background()
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
		ctx,
		*w.db,
		query,
		params,
		neo4j.EagerResultTransformer,
	)
	if err != nil {
		return nil, err
	}

	weekdays := []dao.Weekday{}
	for _, record := range result.Records {
		weekday := dao.Weekday{}
		if err := weekday.ParseFromDB(record); err != nil {
			return nil, err
		}

		weekdays = append(weekdays, weekday)
	}

	return weekdays, nil
}

func WeekdayRepositoryInit(db *neo4j.DriverWithContext) *WeekdayRepositoryImpl {
	return &WeekdayRepositoryImpl{db: db}
}

var weekdayRepositorySet = wire.NewSet(
	WeekdayRepositoryInit,
	wire.Bind(new(WeekdayRepository), new(*WeekdayRepositoryImpl)),
)
