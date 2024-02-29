package repository

import (
	"context"
	"planner-backend/app/domain/dao"
	"planner-backend/app/pkg"

	"github.com/google/wire"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type TimeslotRepository interface {
	FindAllTimeslots(departmentID string, workplaceID string) ([]dao.Timeslot, error)
	FindTimeslotByID(departmentID string, workplaceID string, timeslotID string) (dao.Timeslot, error)
	Save(departmentID string, workplaceID string, timeslot *dao.Timeslot) (dao.Timeslot, error)
	Delete(departmentID string, workplaceID string, timeslot *dao.Timeslot) error
}

type TimeslotRepositoryImpl struct {
	db  *neo4j.DriverWithContext
	ctx context.Context
}

func (t TimeslotRepositoryImpl) FindAllTimeslots(departmentID string, workplaceID string) ([]dao.Timeslot, error) {
	/* Returns all timeslots */

	timeslots := []dao.Timeslot{}
	query := `
    MATCH (d:Department {id: $departmentID})-[:HAS_WORKPLACE]->(wp:Workplace {id: $workplaceID})-[:HAS_TIMESLOT]->(t:Timeslot)
    OPTIONAL MATCH (t)-[r:OFFERED_ON]->(wd:Weekday)
	WHERE t.deleted_at IS NULL
    WITH t, wd, r ORDER BY wd.id
	RETURN t, COLLECT({
		id: wd.id,
		name: wd.name,
		start_time: r.start_time,
		end_time: r.end_time
	}) AS weekdays
    `
	params := map[string]interface{}{
		"departmentID": departmentID,
		"workplaceID":  workplaceID,
	}

	result, err := neo4j.ExecuteQuery(
		t.ctx,
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
		if err := timeslot.ParseFromDB(record, departmentID, workplaceID); err != nil {
			return nil, err
		}

		timeslots = append(timeslots, timeslot)
	}

	return timeslots, nil
}

func (t TimeslotRepositoryImpl) FindTimeslotByID(departmentID string, workplaceID string, timeslotID string) (dao.Timeslot, error) {
	/* Returns a timeslot by name */

	timeslot := dao.Timeslot{}
	query := `
    MATCH (d:Department {id: $departmentID})-[:HAS_WORKPLACE]->(wp:Workplace {id: $workplaceID})-[:HAS_TIMESLOT]->(t:Timeslot {id: $timeslotID})
    OPTIONAL MATCH (t)-[r:OFFERED_ON]->(wd:Weekday)
	WHERE t.deleted_at IS NULL AND d.deleted_at IS NULL AND wp.deleted_at IS NULL
    RETURN t, COLLECT({
        id: wd.id,
        name: wd.name,
        start_time: r.start_time,
        end_time: r.end_time
    }) as weekdays
    `
	params := map[string]interface{}{
		"departmentID": departmentID,
		"workplaceID":  workplaceID,
		"timeslotID":   timeslotID,
	}

	result, err := neo4j.ExecuteQuery(
		t.ctx,
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

	if err := timeslot.ParseFromDB(result.Records[0], departmentID, workplaceID); err != nil {
		return dao.Timeslot{}, err
	}

	return timeslot, nil
}

func (t TimeslotRepositoryImpl) Save(departmentID string, workplaceID string, timeslot *dao.Timeslot) (dao.Timeslot, error) {
	/* Saves a timeslot */

	query := `
	MATCH (d:Department {id: $departmentID})-[:HAS_WORKPLACE]->(wp:Workplace {id: $workplaceID})
	WHERE d.deleted_at IS NULL AND wp.deleted_at IS NULL
	MERGE (t:Timeslot {id: $timeslotID}) <-[:HAS_TIMESLOT]- (wp)
	ON CREATE SET
		t.created_at = datetime(),
		t.name = $timeslotName,
		t.updated_at = datetime(), 
		t.deleted_at = NULL
	ON MATCH SET 
		t.updated_at = datetime(), 
		t.name = $timeslotName,
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
		"departmentID": departmentID,
		"workplaceID":  workplaceID,
		"timeslotID":   timeslot.ID,
		"timeslotName": timeslot.Name,
	}

	result, err := neo4j.ExecuteQuery(
		t.ctx,
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

	if err := timeslot.ParseFromDB(result.Records[0], departmentID, workplaceID); err != nil {
		return dao.Timeslot{}, err
	}

	return *timeslot, nil
}

func (t TimeslotRepositoryImpl) Delete(departmentID string, workplaceID string, timeslot *dao.Timeslot) error {
	/* Deletes a timeslot */

	query := `
    MATCH (d:Department {id: $departmentID})-[:HAS_WORKPLACE]->(wp:Workplace {id: $workplaceID})-[:HAS_TIMESLOT]->(t:Timeslot {id: $timeslotID})
    SET t.deleted_at = datetime()
    `
	params := map[string]interface{}{
		"departmentID": departmentID,
		"workplaceID":  workplaceID,
		"timeslotID":   timeslot.ID,
	}

	_, err := neo4j.ExecuteQuery(
		t.ctx,
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

func TimeslotRepositoryInit(db *neo4j.DriverWithContext, ctx context.Context) *TimeslotRepositoryImpl {
	return &TimeslotRepositoryImpl{
		db:  db,
		ctx: ctx,
	}
}

var timeslotRepositorySet = wire.NewSet(
	TimeslotRepositoryInit,
	wire.Bind(new(TimeslotRepository), new(*TimeslotRepositoryImpl)),
)
