package dao

import (
	"time"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Absence struct {
	PersonID string
	Date     string // Date as string since we only need the date
	Reason   string

	CreatedAt time.Time
}

func (a *Absence) ParseFromDBRecord(relRecord *neo4j.Record, date string, personID string) error {
	/**
	 * Parses an absence from a neo4j record (relationship) and sets the values on this absence
	 */

	absenceRel, _, err := neo4j.GetRecordValue[neo4j.Relationship](relRecord, "r")
	if err != nil {
		return err
	}

	// Get the created_at property from the relationship
	createdAt, err := neo4j.GetProperty[time.Time](absenceRel, "created_at")
	if err != nil {
		return err
	}

	// Get the reason property from the relationship
	reason, err := neo4j.GetProperty[string](absenceRel, "reason")
	if err != nil {
		return err
	}

	a.PersonID = personID
	a.Date = date
	a.Reason = reason
	a.CreatedAt = createdAt

	return nil
}
