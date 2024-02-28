package repository

import (
	"context"
	"planner-backend/app/domain/dao"
	"planner-backend/app/pkg"

	"github.com/google/wire"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type WorkplaceRepository interface {
	// Function Used by the service
	FindAllWorkplaces(departmentID string) ([]dao.Workplace, error)
	FindWorkplaceByID(departmentID string, workplaceID string) (dao.Workplace, error)
	Save(departmentID string, workplace *dao.Workplace) (dao.Workplace, error)
	Delete(departmentID string, workplace *dao.Workplace) error
}

type WorkplaceRepositoryImpl struct {
	db  *neo4j.DriverWithContext
	ctx context.Context
}

func (w WorkplaceRepositoryImpl) FindAllWorkplaces(departmentID string) ([]dao.Workplace, error) {
	/* Returns all workplaces */
	workplaces := []dao.Workplace{}
	query := `
    MATCH (d:Department {id: $departmentID})-[:HAS_WORKPLACE]->(w:Workplace)
	WHERE w.deleted_at IS NULL AND d.deleted_at IS NULL
    RETURN w`
	params := map[string]interface{}{
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

	for _, record := range result.Records {
		workplace := dao.Workplace{}
		if err := workplace.ParseFromDBRecord(record, departmentID); err != nil {
			return nil, err
		}

		workplaces = append(workplaces, workplace)
	}

	return workplaces, nil
}

func (w WorkplaceRepositoryImpl) FindWorkplaceByID(departmentID string, workplaceID string) (dao.Workplace, error) {
	workplace := dao.Workplace{}
	query := `
	MATCH (d:Department {id: $departmentID})-[:HAS_WORKPLACE]->(w:Workplace {id: $workplaceID})
	WHERE w.deleted_at IS NULL AND d.deleted_at IS NULL
	RETURN w`
	params := map[string]interface{}{
		"departmentID": departmentID,
		"workplaceID":  workplaceID,
	}

	result, err := neo4j.ExecuteQuery(
		w.ctx,
		*w.db,
		query,
		params,
		neo4j.EagerResultTransformer,
	)
	if err != nil {
		return dao.Workplace{}, err
	}

	if len(result.Records) == 0 {
		return dao.Workplace{}, pkg.ErrNoRows
	}

	if err := workplace.ParseFromDBRecord(result.Records[0], departmentID); err != nil {
		return dao.Workplace{}, err
	}

	return workplace, nil
}

func (w WorkplaceRepositoryImpl) Save(departmentID string, workplace *dao.Workplace) (dao.Workplace, error) {
	/* Saves a workplace */
	query := `
    MATCH (d:Department {id: $departmentID})
	WHERE d.deleted_at IS NULL
	MERGE (w:Workplace {id: $workplaceID}) <-[:HAS_WORKPLACE]- (d)
    ON CREATE SET
		w.name = $workplaceName,
        w.created_at = datetime(),
        w.updated_at = datetime(),
		w.deleted_at = NULL
	ON MATCH SET
		w.name = $workplaceName,
		w.updated_at = datetime(),
		w.deleted_at = NULL
    RETURN w`
	params := map[string]interface{}{
		"departmentID":  departmentID,
		"workplaceID":   workplace.ID,
		"workplaceName": workplace.Name,
	}

	result, err := neo4j.ExecuteQuery(
		w.ctx,
		*w.db,
		query,
		params,
		neo4j.EagerResultTransformer,
	)
	if err != nil {
		return dao.Workplace{}, err
	}

	if len(result.Records) == 0 {
		return dao.Workplace{}, pkg.ErrNoRows
	}

	if err := workplace.ParseFromDBRecord(result.Records[0], departmentID); err != nil {
		return dao.Workplace{}, err
	}

	return *workplace, nil
}

func (w WorkplaceRepositoryImpl) Delete(departmentID string, workplace *dao.Workplace) error {
	/* Deletes a department */
	query := `
	MATCH  (d:Department {id: $departmentID})-[:HAS_WORKPLACE]->(w:Workplace {id: $workplaceID})
	SET w.deleted_at = datetime()`
	params := map[string]interface{}{
		"departmentID": departmentID,
		"workplaceID":  workplace.ID,
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

func WorkplaceRepositoryInit(db *neo4j.DriverWithContext, ctx context.Context) *WorkplaceRepositoryImpl {
	return &WorkplaceRepositoryImpl{
		db:  db,
		ctx: ctx,
	}
}

var workplaceRepositorySet = wire.NewSet(
	WorkplaceRepositoryInit,
	wire.Bind(new(WorkplaceRepository), new(*WorkplaceRepositoryImpl)),
)
