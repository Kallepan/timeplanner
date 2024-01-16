package dao

import (
	"time"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Person struct {
	Base

	ID        string
	FirstName string
	LastName  string
	Email     string
	Active    bool

	Workplaces  []string
	Departments []string
	Weekdays    []string

	WorkingHours int64
}

func (p *Person) ParseAdditionalFieldsFromDBRecord(record *neo4j.Record) error {
	/**
	* Parses additional fields such as departments, workplaces, and weekdays from a neo4j record and sets the values on this person
	**/

	if workplaces, isNil, err := neo4j.GetRecordValue[[]interface{}](record, "workplaces"); err != nil {
		return err
	} else if !isNil {
		p.Workplaces = convertInterfaceSliceToStringSlice(workplaces)
	}

	if departments, isNil, err := neo4j.GetRecordValue[[]interface{}](record, "departments"); err != nil {
		return err
	} else if !isNil {
		p.Departments = convertInterfaceSliceToStringSlice(departments)
	}

	if weekdays, isNil, err := neo4j.GetRecordValue[[]interface{}](record, "weekdays"); err != nil {
		return err
	} else if !isNil {
		p.Weekdays = convertInterfaceSliceToStringSlice(weekdays)
	}

	return nil
}

func (p *Person) ParseFromDBRecord(record *neo4j.Record) error {
	/**
	 * Parses a person from a neo4j record and sets the values on this person
	 */

	personNode, isNil, err := neo4j.GetRecordValue[neo4j.Node](record, "p")
	if err != nil {
		return err
	}
	if isNil {
		return nil
	}

	id, err := neo4j.GetProperty[string](personNode, "id")
	if err != nil {
		return err
	}
	firstName, err := neo4j.GetProperty[string](personNode, "firstName")
	if err != nil {
		return err
	}
	lastName, err := neo4j.GetProperty[string](personNode, "lastName")
	if err != nil {
		return err
	}
	email, err := neo4j.GetProperty[string](personNode, "email")
	if err != nil {
		return err
	}
	active, err := neo4j.GetProperty[bool](personNode, "active")
	if err != nil {
		return err
	}
	workingHours, err := neo4j.GetProperty[int64](personNode, "workingHours")
	if err != nil {
		return err
	}
	createdAt, err := neo4j.GetProperty[time.Time](personNode, "created_at")
	if err != nil {
		return err
	}
	updatedAt, err := neo4j.GetProperty[time.Time](personNode, "updated_at")
	if err != nil {
		return err
	}
	deletedAtInterface, _ := neo4j.GetProperty[[]any](personNode, "deleted_at")
	deletedAt, err := ConvertNullableValueToTime(deletedAtInterface)
	if err != nil {
		return err
	}

	p.ID = id
	p.FirstName = firstName
	p.LastName = lastName
	p.Email = email
	p.Active = active
	p.WorkingHours = workingHours
	p.Base.CreatedAt = createdAt
	p.Base.UpdatedAt = updatedAt
	p.Base.DeletedAt = deletedAt

	return nil
}
