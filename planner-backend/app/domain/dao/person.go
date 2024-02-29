package dao

import (
	"time"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Weekday struct {
	ID   int64
	Name string
}

func (w *Weekday) ParseFromNode(node *neo4j.Node) error {
	/**
	* Parses a weekday struct from a neo4j node and sets the value on this weekday
	 */

	id, err := neo4j.GetProperty[int64](node, "id")
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

type DepartmentInPerson struct {
	ID   string
	Name string
}
type WorkplaceInPerson struct {
	ID string

	Name         string
	DepartmentID string
}

type Person struct {
	Base

	ID        string
	FirstName string
	LastName  string
	Email     string
	Active    bool

	Workplaces  []WorkplaceInPerson
	Departments []DepartmentInPerson
	Weekdays    []Weekday

	WorkingHours float64
}

func (p *Person) ParseAdditionalFieldsFromDBRecord(record *neo4j.Record) error {
	/**
	* Parses additional fields such as departments, workplaces, and weekdays from a neo4j record and sets the values on this person
	**/

	if departmentInterfaces, isNil, err := neo4j.GetRecordValue[[]interface{}](record, "departments"); err != nil {
		return err
	} else if !isNil {
		for _, departmentInterface := range departmentInterfaces {
			departmentMap, ok := departmentInterface.(map[string]interface{})
			if !ok {
				continue
			}

			var department DepartmentInPerson
			if id, ok := departmentMap["id"].(string); !ok {
				continue
			} else if name, ok := departmentMap["name"].(string); !ok {
				continue
			} else {
				department.ID = id
				department.Name = name
			}

			p.Departments = append(p.Departments, department)
		}
	}

	if workplaceInterfaces, isNil, err := neo4j.GetRecordValue[[]interface{}](record, "workplaces"); err != nil {
		return err
	} else if !isNil {
		for _, workplaceInterface := range workplaceInterfaces {
			workplaceNode, ok := workplaceInterface.(map[string]interface{})
			if !ok {
				continue
			}

			var workplace WorkplaceInPerson
			if id, ok := workplaceNode["id"].(string); !ok {
				continue
			} else if name, ok := workplaceNode["name"].(string); !ok {
				continue
			} else if departmentID, ok := workplaceNode["department_id"].(string); !ok {
				continue
			} else {
				workplace.ID = id
				workplace.Name = name
				workplace.DepartmentID = departmentID
			}

			p.Workplaces = append(p.Workplaces, workplace)
		}
	}

	if weekdays, isNil, err := neo4j.GetRecordValue[[]interface{}](record, "weekdays"); err != nil {
		return err
	} else if !isNil {
		for _, weekdayInterface := range weekdays {
			weekdayNode, ok := weekdayInterface.(map[string]interface{})
			if !ok {
				continue
			}

			var weekday Weekday
			if id, ok := weekdayNode["id"].(int64); !ok {
				continue
			} else if name, ok := weekdayNode["name"].(string); !ok {
				continue
			} else {
				weekday.ID = id
				weekday.Name = name
			}

			p.Weekdays = append(p.Weekdays, weekday)
		}
	}

	return nil
}

func (p *Person) ParseFromNode(personNode *neo4j.Node) error {
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
	workingHours, err := neo4j.GetProperty[float64](personNode, "workingHours")
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

	if err := p.ParseFromNode(&personNode); err != nil {
		return err
	}

	return nil
}
