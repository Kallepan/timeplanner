package dao

import (
	"planner-backend/app/constant"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Workday struct {
	// Metadata to uniquely identify a Workday
	Department Department
	Workplace  Workplace
	Timeslot   Timeslot
	Date       string

	// Assigned Person can be nil
	Person *Person

	// StartTime and EndTime
	StartTime         string
	EndTime           string
	DurationInMinutes int64

	Weekday string
}

func (w *Workday) ParseFromDBRecord(record *neo4j.Record, date string) error {
	/* Parses a workday from a neo4j record and sets the values on this workday */

	// get department Node
	department := Department{}
	departmentNode, _, err := neo4j.GetRecordValue[neo4j.Node](record, "d")
	if err != nil {
		return err
	}
	if err := department.ParseFromNode(&departmentNode); err != nil {
		return err
	}

	// get workplace Node
	workplace := Workplace{}
	workplaceNode, _, err := neo4j.GetRecordValue[neo4j.Node](record, "w")
	if err != nil {
		return err
	}
	if err := workplace.ParseFromNode(&workplaceNode); err != nil {
		return err
	}

	// get timeslot Node
	timeslot := Timeslot{}
	timeslotNode, _, err := neo4j.GetRecordValue[neo4j.Node](record, "t")
	if err := timeslot.ParseFromNode(&timeslotNode); err != nil {
		return err
	}
	if err != nil {
		return err
	}

	// get wkd Node
	workdayNode, _, err := neo4j.GetRecordValue[neo4j.Node](record, "wkd")
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
	duration, err := neo4j.GetProperty[int64](workdayNode, "duration_in_minutes")
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
	w.Department = department
	w.Workplace = workplace
	w.Timeslot = timeslot
	w.Date = date
	w.DurationInMinutes = duration
	w.StartTime = startTime.Time().Format(constant.TimeFormat)
	w.EndTime = endTime.Time().Format(constant.TimeFormat)
	w.Weekday = weekday

	return nil
}
