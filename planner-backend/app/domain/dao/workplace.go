package dao

import (
	"time"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Workplace struct {
	Name string

	Base
}

func (w *Workplace) ParseFromNode(node *neo4j.Node) error {
	/**
	 * Parses a workplace from a neo4j node and sets the values on this workplace
	 */

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

	w.Name = name
	w.Base.CreatedAt = createdAt
	w.Base.UpdatedAt = updatedAt
	w.Base.DeletedAt = deletedAt

	return nil
}

func (w *Workplace) ParseFromDBRecord(record *neo4j.Record) error {
	/**
	 * Parses a workplace from a neo4j record and sets the values on this workplace
	 */

	workplaceNode, _, err := neo4j.GetRecordValue[neo4j.Node](record, "w")
	if err != nil {
		return err
	}

	name, err := neo4j.GetProperty[string](workplaceNode, "name")
	if err != nil {
		return err
	}
	createdAt, err := neo4j.GetProperty[time.Time](workplaceNode, "created_at")
	if err != nil {
		return err
	}
	updatedAt, err := neo4j.GetProperty[time.Time](workplaceNode, "updated_at")
	if err != nil {
		return err
	}
	deletedAtInterface, _ := neo4j.GetProperty[[]any](workplaceNode, "deleted_at")
	deletedAt, err := ConvertNullableValueToTime(deletedAtInterface)
	if err != nil {
		return err
	}

	w.Name = name
	w.Base.CreatedAt = createdAt
	w.Base.UpdatedAt = updatedAt
	w.Base.DeletedAt = deletedAt

	return nil
}
