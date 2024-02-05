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
	GetWorkdaysForDepartmentAndDate(departmentID string, date string) ([]dao.Workday, error)
	GetWorkday(departmentID string, workplaceID string, timeslotID string, date string) (dao.Workday, error)
	Save(workday *dao.Workday) error
	// UpdateWorkday()
	// DeleteWorkday()

	// Main interface to Assign people to a given workday
	AssignPersonToWorkday(personID string, departmentID string, workplaceID string, timeslotID string, date string) error
	UnassignPersonFromWorkday(personID string, departmentID string, workplaceID string, timeslotID string, date string) error
}

type WorkdayRepositoryImpl struct {
	db  *neo4j.DriverWithContext
	ctx context.Context
}

func (w WorkdayRepositoryImpl) GetWorkdaysForDepartmentAndDate(departmentID string, date string) ([]dao.Workday, error) {

	query := `
	// fetch the department
	MATCH (d:Department {id: $departmentID}) -[:HAS_WORKPLACE]-> (w:Workplace) -[:HAS_TIMESLOT]-> (t:Timeslot) <-[:IS_TIMESLOT]- (wkd:Workday {date: date($date)})
	// fetch the person assigned to the workday
	OPTIONAL MATCH (wkd)<-[:ASSIGNED_TO]-(p:Person)
	// if workday is active
	WHERE wkd.active = true
	// return the workday and the person
	RETURN wkd, p, t, w, d
	ORDER BY w.id, t.name
	`
	params := map[string]interface{}{
		"date":         date,
		"departmentID": departmentID,
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

	var workdays []dao.Workday
	for _, record := range result.Records {
		workday := dao.Workday{}
		if err := workday.ParseFromDBRecord(record, date); err != nil {
			return nil, err
		}

		workdays = append(workdays, workday)
	}

	return workdays, nil
}

func (w WorkdayRepositoryImpl) GetWorkday(departmentID string, workplaceID string, timeslotID string, date string) (dao.Workday, error) {
	query := `
	// fetch the department, workplace, timeslot, and the workday
	MATCH (d:Department {id: $departmentID}) -[:HAS_WORKPLACE]-> (w:Workplace {id: $workplaceID}) -[:HAS_TIMESLOT]-> (t:Timeslot {id: $timeslotID}) <-[:IS_TIMESLOT]- (wkd:Workday {date: date($date)})
	// fetch the person assigned to the workday
	OPTIONAL MATCH (wkd)<-[:ASSIGNED_TO]-(p:Person)
	// if workday is active
	WHERE wkd.active = true
	// return the workday and the person
	RETURN wkd, p, t, w, d`

	params := map[string]interface{}{
		"date":         date,
		"departmentID": departmentID,
		"workplaceID":  workplaceID,
		"timeslotID":   timeslotID,
	}

	result, err := neo4j.ExecuteQuery(
		w.ctx,
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
	if err := workday.ParseFromDBRecord(result.Records[0], date); err != nil {
		return dao.Workday{}, err
	}

	return workday, nil
}

func (w WorkdayRepositoryImpl) Save(workday *dao.Workday) error {
	query := `
	MATCH (d:Department {id: $departmentID}) -[:HAS_WORKPLACE]-> (w:Workplace {id: $workplaceID}) -[:HAS_TIMESLOT]-> (t:Timeslot {id: $timeslotID})
	MATCH (dt:Date {date: date($date)})
	MATCH (t) <-[:IS_TIMESLOT]- (wkd:Workday) -[:IS_DATE]-> (dt)
	SET
		wkd.start_time = time($startTime),
		wkd.end_time = time($endTime),
		wkd.duration_in_minutes = duration.between(time($startTime), time($endTime)).minutes,
		wkd.active = $active,
		wkd.updated_at = datetime(),
		wkd.comment = $comment
	// fetch the person assigned to the workday
	WITH wkd, t, d, w, dt
	OPTIONAL MATCH (wkd)<-[:ASSIGNED_TO]-(p:Person)
	RETURN wkd, t, d, w, dt, p
	`
	params := map[string]interface{}{
		"departmentID": workday.Department.ID,
		"workplaceID":  workday.Workplace.ID,
		"timeslotID":   workday.Timeslot.ID,
		"date":         workday.Date,
		"startTime":    workday.StartTime,
		"endTime":      workday.EndTime,
		"active":       workday.Active,
		"comment":      workday.Comment,
	}

	result, err := neo4j.ExecuteQuery(
		w.ctx,
		*w.db,
		query,
		params,
		neo4j.EagerResultTransformer,
	)
	if err != nil {
		return err
	}

	if len(result.Records) == 0 {
		return pkg.ErrNoRows
	}

	if err := workday.ParseFromDBRecord(result.Records[0], workday.Date); err != nil {
		return err
	}

	return nil
}

func (w WorkdayRepositoryImpl) AssignPersonToWorkday(personID string, departmentID string, workplaceID string, timeslotID string, date string) error {

	query := `
	// delete the relationship between the person and the workday
	MATCH (wkd:Workday {date: date($date), department: $departmentID, workplace: $workplaceID, timeslot: $timeslotID, active: true})
	// delete the relationship between the person and the workday
	OPTIONAL MATCH (wkd)<-[r:ASSIGNED_TO]-(:Person)
	DELETE r
	WITH wkd
	// fetch the person
	MATCH (p:Person {id: $personID})
	// create a relationship between the person and the workday
	MERGE (p)-[r:ASSIGNED_TO]->(wkd)
	ON CREATE SET r.created_at = datetime()
	RETURN p, r, wkd
	`
	params := map[string]interface{}{
		"personID":     personID,
		"date":         date,
		"departmentID": departmentID,
		"workplaceID":  workplaceID,
		"timeslotID":   timeslotID,
	}

	result, err := neo4j.ExecuteQuery(
		w.ctx,
		*w.db,
		query,
		params,
		neo4j.EagerResultTransformer,
	)
	if err != nil {
		return err
	}

	if len(result.Records) == 0 {
		return pkg.ErrDidNotCreateRelationship
	}

	return nil
}

func (w WorkdayRepositoryImpl) UnassignPersonFromWorkday(personID string, departmentID string, workplaceID string, timeslotID string, date string) error {

	query := `
	MATCH (wkd:Workday {date: date($date), department: $departmentID, workplace: $workplaceID, timeslot: $timeslotID, active: true})
	// delete the relationship between the person and the workday
	MATCH (wkd)<-[r:ASSIGNED_TO]-(p:Person {id: $personID})
	DELETE r
	`
	params := map[string]interface{}{
		"personID":     personID,
		"date":         date,
		"departmentID": departmentID,
		"workplaceID":  workplaceID,
		"timeslotID":   timeslotID,
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

func WorkdayRepositoryInit(db *neo4j.DriverWithContext, ctx context.Context) *WorkdayRepositoryImpl {
	return &WorkdayRepositoryImpl{
		db:  db,
		ctx: ctx,
	}
}

var workdayRepositorySet = wire.NewSet(
	WorkdayRepositoryInit,
	wire.Bind(new(WorkdayRepository), new(*WorkdayRepositoryImpl)),
)
