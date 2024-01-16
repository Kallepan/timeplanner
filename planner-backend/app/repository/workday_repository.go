package repository

import (
	"context"
	"planner-backend/app/domain/dao"
	"planner-backend/app/pkg"

	"github.com/google/wire"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type WorkdayRepository interface {
	/*
	 * Gets all Workdays along with the (if present) assigned user
	 * for a given date
	 */
	GetWorkdaysForDepartmentAndDate(departmentName string, date string) ([]dao.Workday, error)
	GetWorkday(departmentName string, workplaceName string, timeslotName string, date string) (dao.Workday, error)
	// TODO:
	// UpdateWorkday()

	// Main interface to Assign people to a given workday
	AssignPersonToWorkday(personID string, departmentName string, workplaceName string, timeslotName string, date string) error
	UnassignPersonFromWorkday(personID string, departmentName string, workplaceName string, timeslotName string, date string) error
}

type WorkdayRepositoryImpl struct {
	db *neo4j.DriverWithContext
}

func (w WorkdayRepositoryImpl) GetWorkdaysForDepartmentAndDate(departmentName string, date string) ([]dao.Workday, error) {
	ctx := context.Background()
	query := `
	// fetch the workdays for a given date
	MATCH (wkd:Workday {date: date($date), department: $departmentName})
	// fetch the person assigned to the workday
	OPTIONAL MATCH (wkd)<-[:ASSIGNED_TO]-(p:Person)
	// if workday is active
	WHERE wkd.active = true
	// return the workday and the person
	RETURN wkd, p
	`
	params := map[string]interface{}{
		"date":           date,
		"departmentName": departmentName,
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

	var workdays []dao.Workday
	for _, record := range result.Records {
		workday := dao.Workday{}
		if err := workday.ParseFromDBRecord(record, departmentName, date); err != nil {
			return nil, err
		}

		workdays = append(workdays, workday)
	}

	return workdays, nil
}

func (w WorkdayRepositoryImpl) GetWorkday(departmentName string, workplaceName string, timeslotName string, date string) (dao.Workday, error) {
	ctx := context.Background()
	query := `
	// fetch the workday
	MATCH (wkd:Workday {date: date($date), department: $departmentName, workplace: $workplaceName, timeslot: $timeslotName})
	// fetch the person assigned to the workday
	OPTIONAL MATCH (wkd)<-[:ASSIGNED_TO]-(p:Person)
	// if workday is active
	WHERE wkd.active = true
	// return the workday and the person
	RETURN wkd, p`

	params := map[string]interface{}{
		"date":           date,
		"departmentName": departmentName,
		"workplaceName":  workplaceName,
		"timeslotName":   timeslotName,
	}

	result, err := neo4j.ExecuteQuery(
		ctx,
		*w.db,
		query,
		params,
		neo4j.EagerResultTransformer,
	)
	if err != nil {
		return dao.Workday{}, err
	}

	if len(result.Records) == 0 {
		return dao.Workday{}, pkg.ErrNoRows
	}

	workday := dao.Workday{}
	if err := workday.ParseFromDBRecord(result.Records[0], departmentName, date); err != nil {
		return dao.Workday{}, err
	}

	return workday, nil
}

func (w WorkdayRepositoryImpl) AssignPersonToWorkday(personID string, departmentName string, workplaceName string, timeslotName string, date string) error {
	ctx := context.Background()
	query := `
	// fetch the person
	MATCH (p:Person {id: $personID})
	// fetch the workday
	MATCH (wkd:Workday {date: date($date), department: $departmentName, workplace: $workplaceName, timeslot: $timeslotName})
	// create a relationship between the person and the workday
	MERGE (p)-[r:ASSIGNED_TO]->(wkd)
	ON CREATE SET r.created_at = timestamp()
	`
	params := map[string]interface{}{
		"personID":       personID,
		"date":           date,
		"departmentName": departmentName,
		"workplaceName":  workplaceName,
		"timeslotName":   timeslotName,
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

func (w WorkdayRepositoryImpl) UnassignPersonFromWorkday(personID string, departmentName string, workplaceName string, timeslotName string, date string) error {
	ctx := context.Background()
	query := `
	MATCH (wkd:Workday {date: date($date), department: $departmentName, workplace: $workplaceName, timeslot: $timeslotName})
	// delete the relationship between the person and the workday
	MATCH (wkd)<-[r:ASSIGNED_TO]-(p:Person {id: $personID})
	DELETE r
	`
	params := map[string]interface{}{
		"personID":       personID,
		"date":           date,
		"departmentName": departmentName,
		"workplaceName":  workplaceName,
		"timeslotName":   timeslotName,
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

func WorkdayRepositoryInit(db *neo4j.DriverWithContext) *WorkdayRepositoryImpl {
	return &WorkdayRepositoryImpl{
		db: db,
	}
}

var workdayRepositorySet = wire.NewSet(
	WorkdayRepositoryInit,
	wire.Bind(new(WorkdayRepository), new(*WorkdayRepositoryImpl)),
)