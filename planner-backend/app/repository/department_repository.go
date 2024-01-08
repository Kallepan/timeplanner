package repository

import (
	"context"
	"planner-backend/app/domain/dao"
	"planner-backend/app/pkg"

	"github.com/google/wire"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type DepartmentRepository interface {
	// Function Used by the service
	FindAllDepartments() ([]dao.Department, error)
	FindDepartmentByName(departmentName string) (dao.Department, error)
	Save(department *dao.Department) (dao.Department, error)
	Delete(department *dao.Department) error
}

type DepartmentRepositoryImpl struct {
	db *neo4j.DriverWithContext
}

func (d DepartmentRepositoryImpl) FindAllDepartments() ([]dao.Department, error) {
	/* Returns all departments */
	ctx := context.Background()
	departments := []dao.Department{}
	query := `
	MATCH (d:Department)
	WHERE d.deleted_at IS NULL
	RETURN d`

	result, err := neo4j.ExecuteQuery(
		ctx,
		*d.db,
		query,
		nil,
		neo4j.EagerResultTransformer,
	)
	if err != nil {
		return nil, err
	}

	for _, record := range result.Records {
		department := dao.Department{}
		if err := department.ParseDepartmentFromDBRecord(record); err != nil {
			return nil, err
		}

		departments = append(departments, department)
	}

	return departments, nil
}

func (d DepartmentRepositoryImpl) FindDepartmentByName(departmentName string) (dao.Department, error) {
	/* Returns a department by name */
	ctx := context.Background()
	department := dao.Department{}
	query := `
	MATCH (d:Department)
	WHERE d.name = $name AND d.deleted_at IS NULL
	RETURN d`
	params := map[string]interface{}{
		"name": departmentName,
	}

	result, err := neo4j.ExecuteQuery(
		ctx,
		*d.db,
		query,
		params,
		neo4j.EagerResultTransformer,
	)
	if err != nil {
		return department, err
	}
	if len(result.Records) == 0 {
		return department, pkg.ErrNoRows
	}

	if err := department.ParseDepartmentFromDBRecord(result.Records[0]); err != nil {
		return department, err
	}

	return department, nil
}

func (d DepartmentRepositoryImpl) Save(department *dao.Department) (dao.Department, error) {
	/* Creates a department */
	ctx := context.Background()
	query := `
	MERGE (d:Department {name: $name})
	ON CREATE SET
		d.created_at = datetime(),
		d.updated_at = datetime(),
		d.deleted_at = NULL
    ON MATCH SET
        d.updated_at = datetime(),
		d.deleted_at = NULL
	RETURN d`
	params := map[string]interface{}{
		"name": department.Name,
	}

	result, err := neo4j.ExecuteQuery(
		ctx,
		*d.db,
		query,
		params,
		neo4j.EagerResultTransformer,
	)
	if err != nil {
		return *department, err
	}
	if len(result.Records) == 0 {
		return *department, pkg.ErrNoRows
	}

	if err := department.ParseDepartmentFromDBRecord(result.Records[0]); err != nil {
		return *department, err
	}

	return *department, nil
}

func (d DepartmentRepositoryImpl) Delete(department *dao.Department) error {
	/* Deletes a department */
	ctx := context.Background()
	query := `
	MATCH (d:Department)
	WHERE d.name = $name
	SET d.deleted_at = datetime()`
	params := map[string]interface{}{
		"name": department.Name,
	}

	_, err := neo4j.ExecuteQuery(
		ctx,
		*d.db,
		query,
		params,
		neo4j.EagerResultTransformer,
	)
	if err != nil {
		return err
	}

	return nil
}

func DepartmentRepositoryInit(db *neo4j.DriverWithContext) *DepartmentRepositoryImpl {
	return &DepartmentRepositoryImpl{
		db: db,
	}
}

var departmentRepositorySet = wire.NewSet(
	DepartmentRepositoryInit,
	wire.Bind(new(DepartmentRepository), new(*DepartmentRepositoryImpl)),
)
