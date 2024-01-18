package repository

import (
	"context"
	"planner-backend/app/domain/dao"
	"planner-backend/app/pkg"

	"github.com/google/wire"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type TimeslotRepository interface {
	FindAllTimeslots(departmentName string, workplaceName string) ([]dao.Timeslot, error)
	FindTimeslotByName(departmentName string, workplaceName string, timeslotName string) (dao.Timeslot, error)
	Save(departmentName string, workplaceName string, timeslot *dao.Timeslot) (dao.Timeslot, error)
	Delete(departmentName string, workplaceName string, timeslot *dao.Timeslot) error
}

type TimeslotRepositoryImpl struct {
	db *neo4j.DriverWithContext
}

func (t TimeslotRepositoryImpl) FindAllTimeslots(departmentName string, workplaceName string) ([]dao.Timeslot, error) {
	/* Returns all timeslots */
	ctx := context.Background()
	timeslots := []dao.Timeslot{}
	query := `
    MATCH (d:Department {name: $departmentName})-[:HAS_WORKPLACE]->(wp:Workplace {name: $workplaceName})-[:HAS_TIMESLOT]->(t:Timeslot)
    OPTIONAL MATCH (t)-[r:OFFERED_ON]->(wd:Weekday)
	WHERE t.deleted_at IS NULL AND t.active = true
    RETURN t,
        COLLECT({
            id: wd.id,
            name: wd.name,
            start_time: r.start_time,
            end_time: r.end_time
	}) AS weekdays
    `
	params := map[string]interface{}{
		"departmentName": departmentName,
		"workplaceName":  workplaceName,
	}

	result, err := neo4j.ExecuteQuery(
		ctx,
		*t.db,
		query,
		params,
		neo4j.EagerResultTransformer,
	)
	if err != nil {
		return nil, err
	}

	for _, record := range result.Records {
		timeslot := dao.Timeslot{}
		if err := timeslot.ParseFromDB(record, departmentName, workplaceName); err != nil {
			return nil, err
		}

		timeslots = append(timeslots, timeslot)
	}

	return timeslots, nil
}

func (t TimeslotRepositoryImpl) FindTimeslotByName(departmentName string, workplaceName string, timeslotName string) (dao.Timeslot, error) {
	/* Returns a timeslot by name */
	ctx := context.Background()
	timeslot := dao.Timeslot{}
	query := `
    MATCH (d:Department {name: $departmentName})-[:HAS_WORKPLACE]->(wp:Workplace {name: $workplaceName})-[:HAS_TIMESLOT]->(t:Timeslot {name: $timeslotName})
    OPTIONAL MATCH (t)-[r:OFFERED_ON]->(wd:Weekday)
	WHERE t.deleted_at IS NULL AND t.active = true
    RETURN t, COLLECT({
        id: wd.id,
        name: wd.name,
        start_time: r.start_time,
        end_time: r.end_time
    }) as weekdays
    `
	params := map[string]interface{}{
		"departmentName": departmentName,
		"workplaceName":  workplaceName,
		"timeslotName":   timeslotName,
	}

	result, err := neo4j.ExecuteQuery(
		ctx,
		*t.db,
		query,
		params,
		neo4j.EagerResultTransformer,
	)
	if err != nil {
		return dao.Timeslot{}, err
	}

	if len(result.Records) == 0 {
		return dao.Timeslot{}, pkg.ErrNoRows
	}

	if err := timeslot.ParseFromDB(result.Records[0], departmentName, workplaceName); err != nil {
		return dao.Timeslot{}, err
	}

	return timeslot, nil
}

func (t TimeslotRepositoryImpl) Save(departmentName string, workplaceName string, timeslot *dao.Timeslot) (dao.Timeslot, error) {
	/* Saves a timeslot */
	ctx := context.Background()
	query := `
	MATCH (d:Department {name: $departmentName})-[:HAS_WORKPLACE]->(wp:Workplace {name: $workplaceName})
	MERGE (t:Timeslot {name: $timeslotName}) <-[:HAS_TIMESLOT]- (wp)
	ON CREATE SET 
		t.created_at = datetime(), 
		t.active = $active, 
		t.updated_at = datetime(), 
		t.deleted_at = NULL
	ON MATCH SET 
		t.updated_at = datetime(), 
		t.active = $active, 
		t.deleted_at = NULL
	WITH t
	OPTIONAL MATCH (t)-[r:OFFERED_ON]->(wd:Weekday)
	RETURN t, COLLECT({
		id: wd.id,
		name: wd.name,
		start_time: r.start_time,
		end_time: r.end_time
	}) as weekdays
	`
	params := map[string]interface{}{
		"departmentName": departmentName,
		"workplaceName":  workplaceName,
		"timeslotName":   timeslot.Name,
		"active":         timeslot.Active,
	}

	result, err := neo4j.ExecuteQuery(
		ctx,
		*t.db,
		query,
		params,
		neo4j.EagerResultTransformer,
	)
	if err != nil {
		return dao.Timeslot{}, err
	}

	if len(result.Records) == 0 {
		return dao.Timeslot{}, pkg.ErrNoRows
	}

	if err := timeslot.ParseFromDB(result.Records[0], departmentName, workplaceName); err != nil {
		return dao.Timeslot{}, err
	}

	return *timeslot, nil
}

func (t TimeslotRepositoryImpl) Delete(departmentName string, workplaceName string, timeslot *dao.Timeslot) error {
	/* Deletes a timeslot */
	ctx := context.Background()
	query := `
    MATCH (d:Department {name: $departmentName})-[:HAS_WORKPLACE]->(wp:Workplace {name: $workplaceName})-[:HAS_TIMESLOT]->(t:Timeslot {name: $timeslotName})
    SET t.deleted_at = datetime()
    `
	params := map[string]interface{}{
		"departmentName": departmentName,
		"workplaceName":  workplaceName,
		"timeslotName":   timeslot.Name,
	}

	_, err := neo4j.ExecuteQuery(
		ctx,
		*t.db,
		query,
		params,
		neo4j.EagerResultTransformer,
	)
	if err != nil {
		return err
	}

	return nil
}

func TimeslotRepositoryInit(db *neo4j.DriverWithContext) *TimeslotRepositoryImpl {
	return &TimeslotRepositoryImpl{
		db: db,
	}
}

var timeslotRepositorySet = wire.NewSet(
	TimeslotRepositoryInit,
	wire.Bind(new(TimeslotRepository), new(*TimeslotRepositoryImpl)),
)
