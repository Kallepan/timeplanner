package dao

import (
	"time"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Department struct {
	Name string

	Base
}

func (d *Department) ParseDepartmentFromDBRecord(record *neo4j.Record) error {
	/**
	 * Parses a department from a neo4j record and sets the values on this department
	 */

	departmentNode, _, err := neo4j.GetRecordValue[neo4j.Node](record, "d")
	if err != nil {
		return err
	}

	name, err := neo4j.GetProperty[string](departmentNode, "name")
	if err != nil {
		return err
	}
	createdAt, err := neo4j.GetProperty[time.Time](departmentNode, "created_at")
	if err != nil {
		return err
	}
	updatedAt, err := neo4j.GetProperty[time.Time](departmentNode, "updated_at")
	if err != nil {
		return err
	}
	deletedAt, err := neo4j.GetProperty[time.Time](departmentNode, "deleted_at")
	if err != nil {
		return err
	}

	d.Name = name
	d.Base.CreatedAt = createdAt
	d.Base.UpdatedAt = updatedAt
	d.Base.DeletedAt = deletedAt

	return nil
}
