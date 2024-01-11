package dao

import (
	"errors"
	"time"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Weekday struct {
	ID        string
	Name      string
	StartTime time.Time
	EndTime   time.Time
}

type Timeslot struct {
	Name           string
	Active         bool
	DepartmentName string
	WorkplaceName  string
	Weekdays       []Weekday
	Base
}

func (w *Weekday) ParseFromMap(data map[string]interface{}) error {
	id, ok := data["id"].(string)
	if !ok {
		return errors.New("could not parse id")
	}

	name, ok := data["name"].(string)
	if !ok {
		return errors.New("could not parse name")
	}

	startTime, ok := data["start_time"].(neo4j.Time)
	if !ok {
		return errors.New("could not parse start_time")
	}

	endTime, ok := data["end_time"].(neo4j.Time)
	if !ok {
		return errors.New("could not parse end_time")
	}

	w.ID = id
	w.Name = name
	w.StartTime = startTime.Time()
	w.EndTime = endTime.Time()

	return nil
}

func (w *Weekday) ParseFromDB(record *neo4j.Record) error {
	weekdayNode, _, err := neo4j.GetRecordValue[neo4j.Node](record, "w")
	if err != nil {
		return err
	}

	id, err := neo4j.GetProperty[string](weekdayNode, "id")
	if err != nil {
		return err
	}

	name, err := neo4j.GetProperty[string](weekdayNode, "name")
	if err != nil {
		return err
	}

	startTime, err := neo4j.GetProperty[neo4j.Time](weekdayNode, "start_time")
	if err != nil {
		return err
	}
	startTimeFormatted := startTime.Time()

	endTime, err := neo4j.GetProperty[neo4j.Time](weekdayNode, "end_time")
	if err != nil {
		return err
	}
	endTimeFormatted := endTime.Time()

	w.ID = id
	w.Name = name
	w.StartTime = startTimeFormatted
	w.EndTime = endTimeFormatted

	return nil
}

func (t *Timeslot) ParseFromDB(record *neo4j.Record, departmentName string, workplaceName string) error {
	timelotNode, _, err := neo4j.GetRecordValue[neo4j.Node](record, "t")
	if err != nil {
		return err
	}

	name, err := neo4j.GetProperty[string](timelotNode, "name")
	if err != nil {
		return err
	}

	active, err := neo4j.GetProperty[bool](timelotNode, "active")
	if err != nil {
		return err
	}

	createdAt, err := neo4j.GetProperty[time.Time](timelotNode, "created_at")
	if err != nil {
		return err
	}

	updatedAt, err := neo4j.GetProperty[time.Time](timelotNode, "updated_at")
	if err != nil {
		return err
	}

	deletedAtInterface, _ := neo4j.GetProperty[[]any](timelotNode, "deleted_at")
	deletedAt, err := ConvertNullableValueToTime(deletedAtInterface)
	if err != nil {
		return err
	}

	t.Name = name
	t.Active = active
	t.DepartmentName = departmentName
	t.WorkplaceName = workplaceName
	t.Base.CreatedAt = createdAt
	t.Base.UpdatedAt = updatedAt
	t.Base.DeletedAt = deletedAt

	weekdays, ok, err := neo4j.GetRecordValue[[]any](record, "weekdays")
	if err != nil {
		return err
	}
	if !ok {
		return nil
	}

	var weekdaysParsed []Weekday
	for _, weekdayInterface := range weekdays {
		weekday, ok := weekdayInterface.(map[string]interface{})
		if !ok {
			return errors.New("could not parse weekday")
		}

		var weekdayParsed Weekday
		err := weekdayParsed.ParseFromMap(weekday)
		if err != nil {
			// This is due to the structure of the database
			// If the weekday is not found, it will be null
			// So we just skip it
			continue
		}
		weekdaysParsed = append(weekdaysParsed, weekdayParsed)
	}
	t.Weekdays = weekdaysParsed

	return nil
}
