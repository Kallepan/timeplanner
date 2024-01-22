package dao

import (
	"planner-backend/app/constant"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Workday struct {
	// Metadata to uniquely identify a Workday
	DepartmentID string
	WorkplaceID  string
	TimeslotName string
	Date         string

	// Assigned Person can be nil
	Person *Person

	// StartTime and EndTime
	StartTime string
	EndTime   string

	Weekday string
}

func (w *Workday) ParseFromDBRecord(record *neo4j.Record, departmentID string, date string) error {
	/* Parses a workday from a neo4j record and sets the values on this workday */

	// get wkd Node
	workdayNode, _, err := neo4j.GetRecordValue[neo4j.Node](record, "wkd")
	if err != nil {
		return err
	}

	// get wkd properties
	workplaceID, err := neo4j.GetProperty[string](workdayNode, "workplace")
	if err != nil {
		return err
	}
	timeslotName, err := neo4j.GetProperty[string](workdayNode, "timeslot")
	if err != nil {
		return err
	}

	startTime, err := neo4j.GetProperty[neo4j.Time](workdayNode, "start_time")
	if err != nil {
		return err
	}
	endTime, err := neo4j.GetProperty[neo4j.Time](workdayNode, "end_time")
	if err != nil {
		return err
	}

	weekday, err := neo4j.GetProperty[string](workdayNode, "weekday")
	if err != nil {
		return err
	}

	// get person Node
	// If the person is not assigned to the workday, the person Node will be nil
	// I am sorry for this ugly code
	person := Person{}
	if err := person.ParseFromDBRecord(record); err != nil {
		return err
	} else if person.ID != "" {
		w.Person = &person
	}

	// set values on workday
	w.DepartmentID = departmentID
	w.WorkplaceID = workplaceID
	w.TimeslotName = timeslotName
	w.Date = date
	w.StartTime = startTime.Time().Format(constant.TimeFormat)
	w.EndTime = endTime.Time().Format(constant.TimeFormat)
	w.Weekday = weekday

	return nil
}
