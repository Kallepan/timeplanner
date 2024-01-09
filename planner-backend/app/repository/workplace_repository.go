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
	FindAllWorkplaces(departmentName string) ([]dao.Workplace, error)
	FindWorkplaceByName(departmentName string, workplaceName string) (dao.Workplace, error)
	Save(departmentName string, workplace *dao.Workplace) (dao.Workplace, error)
	Delete(departmentName string, workplace *dao.Workplace) error
}

type WorkplaceRepositoryImpl struct {
	db *neo4j.DriverWithContext
}

func (w WorkplaceRepositoryImpl) FindAllWorkplaces(departmentName string) ([]dao.Workplace, error) {
	/* Returns all workplaces */
	ctx := context.Background()
	workplaces := []dao.Workplace{}
	query := `
    MATCH (d:Department {name: $departmentName})-[:HAS_WORKPLACE]->(w:Workplace)
	WHERE w.deleted_at IS NULL
    RETURN w`
	params := map[string]interface{}{
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

	for _, record := range result.Records {
		workplace := dao.Workplace{}
		if err := workplace.ParseWorkplaceFromDBRecord(record); err != nil {
			return nil, err
		}

		workplaces = append(workplaces, workplace)
	}

	return workplaces, nil
}

func (w WorkplaceRepositoryImpl) FindWorkplaceByName(departmentName string, workplaceName string) (dao.Workplace, error) {
	/* Returns a workplace by name */
	ctx := context.Background()
	workplace := dao.Workplace{}
	query := `
    MATCH (d:Department {name: $departmentName})-[:HAS_WORKPLACE]->(w:Workplace {name: $workplaceName})
    RETURN w`
	params := map[string]interface{}{
		"departmentName": departmentName,
		"workplaceName":  workplaceName,
	}

	result, err := neo4j.ExecuteQuery(
		ctx,
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

	if err := workplace.ParseWorkplaceFromDBRecord(result.Records[0]); err != nil {
		return dao.Workplace{}, err
	}

	return workplace, nil
}

func (w WorkplaceRepositoryImpl) Save(departmentName string, workplace *dao.Workplace) (dao.Workplace, error) {
	/* Saves a workplace */
	ctx := context.Background()
	query := `
    MATCH (d:Department {name: $departmentName})
    MERGE (d)-[:HAS_WORKPLACE]->(w:Workplace {name: $workplaceName})
    ON CREATE SET
        w.created_at = datetime(),
        w.updated_at = datetime()
    RETURN w`
	params := map[string]interface{}{
		"departmentName": departmentName,
		"workplaceName":  workplace.Name,
	}

	result, err := neo4j.ExecuteQuery(
		ctx,
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

	if err := workplace.ParseWorkplaceFromDBRecord(result.Records[0]); err != nil {
		return dao.Workplace{}, err
	}

	return *workplace, nil
}

func WorkplaceRepositoryInit(db *neo4j.DriverWithContext) *WorkplaceRepositoryImpl {
	return &WorkplaceRepositoryImpl{
		db: db,
	}
}

var workplaceRepositorySet = wire.NewSet(
	WorkplaceRepositoryInit,
	wire.Bind(new(WorkplaceRepository), new(*WorkplaceRepositoryImpl)),
)
