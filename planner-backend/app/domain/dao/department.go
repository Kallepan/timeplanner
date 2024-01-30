package dao

import (
	"time"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Department struct {
	Name string
	ID   string

	Base
}

func (d *Department) ParseFromNode(node *neo4j.Node) error {
	/**
	 * Parses a department from a neo4j node and sets the values on this department
	 */

	id, err := neo4j.GetProperty[string](node, "id")
	if err != nil {
		return err
	}
	name, err := neo4j.GetProperty[string](node, "name")
	if err != nil {
		return err
	}
	createdAt, err := neo4j.GetProperty[time.Time](node, "created_at")
	if err != nil {
		return err
	}
	updatedAt, err := neo4j.GetProperty[time.Time](node, "updated_at")
	if err != nil {
		return err
	}
	deletedAtInterface, _ := neo4j.GetProperty[[]any](node, "deleted_at")
	deletedAt, err := ConvertNullableValueToTime(deletedAtInterface)
	if err != nil {
		return err
	}

	d.ID = id
	d.Name = name
	d.Base.CreatedAt = createdAt
	d.Base.UpdatedAt = updatedAt
	d.Base.DeletedAt = deletedAt

	return nil
}

func (d *Department) ParseFromDB(record *neo4j.Record) error {
	/**
	 * Parses a department from a neo4j record and sets the values on this department
	 */

	departmentNode, _, err := neo4j.GetRecordValue[neo4j.Node](record, "d")
	if err != nil {
		return err
	}

	d.ParseFromNode(&departmentNode)

	return nil
}
