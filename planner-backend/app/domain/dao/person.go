package dao

import (
	"time"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Weekday struct {
	ID   string
	Name string
}

func (w *Weekday) ParseFromNode(node *neo4j.Node) error {
	/**
	* Parses a weekday struct from a neo4j node and sets the value on this weekday
	 */

	id, err := neo4j.GetProperty[string](node, "id")
	if err != nil {
		return err
	}

	name, err := neo4j.GetProperty[string](node, "name")
	if err != nil {
		return err
	}

	w.ID = id
	w.Name = name

	return nil
}

type Person struct {
	Base

	ID        string
	FirstName string
	LastName  string
	Email     string
	Active    bool

	Workplaces  []Workplace
	Departments []Department
	Weekdays    []Weekday

	WorkingHours int64
}

func (p *Person) ParseAdditionalFieldsFromDBRecord(record *neo4j.Record) error {
	/**
	* Parses additional fields such as departments, workplaces, and weekdays from a neo4j record and sets the values on this person
	**/

	if workplaceNodeInteraces, isNil, err := neo4j.GetRecordValue[[]interface{}](record, "workplaces"); err != nil {
		return err
	} else if !isNil {
		for _, workplaceNodeInterface := range workplaceNodeInteraces {
			workplaceNode, ok := workplaceNodeInterface.(neo4j.Node)
			if !ok {
				continue
			}
			workplace := Workplace{}
			if err := workplace.ParseFromNode(&workplaceNode); err != nil {
				return err
			}
			p.Workplaces = append(p.Workplaces, workplace)
		}
	}

	if departments, isNil, err := neo4j.GetRecordValue[[]interface{}](record, "departments"); err != nil {
		return err
	} else if !isNil {
		for _, departmentInterface := range departments {
			departmentNode, ok := departmentInterface.(neo4j.Node)
			if !ok {
				continue
			}
			department := Department{}
			if err := department.ParseFromNode(&departmentNode); err != nil {
				return err
			}
			p.Departments = append(p.Departments, department)
		}
	}

	if weekdays, isNil, err := neo4j.GetRecordValue[[]interface{}](record, "weekdays"); err != nil {
		return err
	} else if !isNil {
		for _, weekdayInterface := range weekdays {
			weekdayNode, ok := weekdayInterface.(neo4j.Node)
			if !ok {
				continue
			}
			weekday := Weekday{}
			if err := weekday.ParseFromNode(&weekdayNode); err != nil {
				return err
			}
			p.Weekdays = append(p.Weekdays, weekday)
		}
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
