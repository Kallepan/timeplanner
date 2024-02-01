package repository

import (
	"context"
	"planner-backend/app/domain/dao"
	"planner-backend/app/pkg"

	"github.com/google/wire"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type PersonRelRepository interface {
	// Function Used by the service
	AddAbsencyToPerson(person dao.Person, absence dao.Absence) error
	RemoveAbsencyFromPerson(person dao.Person, absence dao.Absence) error
	FindAbsencyForPerson(personID string, date string) (dao.Absence, error)

	AddDepartmentToPerson(person dao.Person, departmentID string) error
	RemoveDepartmentFromPerson(person dao.Person, departmentID string) error

	AddWorkplaceToPerson(person dao.Person, departmentID string, workplaceID string) error
	RemoveWorkplaceFromPerson(person dao.Person, departmentID string, workplaceID string) error

	AddWeekdayToPerson(person dao.Person, weekdayID string) error
	RemoveWeekdayFromPerson(person dao.Person, weekdayID string) error
}

type PersonRelRepositoryImpl struct {
	db  *neo4j.DriverWithContext
	ctx context.Context
}

// Function Used by the service
func (p PersonRelRepositoryImpl) AddAbsencyToPerson(person dao.Person, absence dao.Absence) error {
	/* Adds an absency to a person
	   @param person: The person to add the absency to
	   @param date: The date of the absency
	*/

	// Ensure that the date exists
	if err := EnsureDateExists(p.db, context.Background(), absence.Date); err != nil {
		return err
	}

	query := `
	MATCH (p: Person {id: $personID})
	MATCH (d: Date {date: date($date)})
	MERGE (p) -[r:ABSENT_ON]-> (d)
	ON CREATE SET r.created_at = datetime()
	SET r.reason = $reason
	WITH d, p
	OPTIONAL MATCH (d) <-[:IS_DATE]- (wkd: Workday) <-[r:ASSIGNED_TO]- (p)
	DELETE r
	`

	params := map[string]interface{}{
		"personID": person.ID,
		"date":     absence.Date,
		"reason":   absence.Reason,
	}

	_, err := neo4j.ExecuteQuery(
		p.ctx,
		*p.db,
		query,
		params,
		neo4j.EagerResultTransformer,
	)
	if err != nil {
		return err
	}

	return nil
}

func (p PersonRelRepositoryImpl) RemoveAbsencyFromPerson(person dao.Person, absence dao.Absence) error {
	/* Removes an absency from a person
	   @param person: The person to remove the absency from
	   @param date: The date of the absency
	*/

	query := `
	MATCH (p: Person {id: $personID}) -[r:ABSENT_ON]-> (d: Date {date: date($date)})
	DELETE r
	`
	params := map[string]interface{}{
		"personID": person.ID,
		"date":     absence.Date,
	}

	_, err := neo4j.ExecuteQuery(
		p.ctx,
		*p.db,
		query,
		params,
		neo4j.EagerResultTransformer,
	)
	if err != nil {
		return err
	}

	return nil
}

func (p PersonRelRepositoryImpl) FindAbsencyForPerson(personID string, date string) (dao.Absence, error) {
	/* Finds an absency for a person
	   @param personID: The ID of the person to find the absency for
	   @param date: The date of the absency
	*/

	query := `
	MATCH (p: Person {id: $personID}) -[r:ABSENT_ON]-> (d: Date {date: date($date)})
	RETURN r`
	params := map[string]interface{}{
		"personID": personID,
		"date":     date,
	}

	result, err := neo4j.ExecuteQuery(
		p.ctx,
		*p.db,
		query,
		params,
		neo4j.EagerResultTransformer,
	)
	if err != nil {
		return dao.Absence{}, err
	}

	if len(result.Records) == 0 {
		return dao.Absence{}, pkg.ErrNoRows
	}

	absence := dao.Absence{}
	record := result.Records[0]
	if err := absence.ParseFromDBRecord(record, date, personID); err != nil {
		return dao.Absence{}, err
	}

	return absence, nil
}

func (p PersonRelRepositoryImpl) AddDepartmentToPerson(person dao.Person, departmentID string) error {
	/* Adds a department to a person
	   @param person: The person to add the department to
	   @param departmentID: The name of the department to add
	*/

	query := `
	MATCH (p: Person {id: $personID})
	MATCH (d: Department {id: $departmentID})
	MERGE (p) -[r:WORKS_AT]-> (d)
	ON CREATE SET r.created_at = datetime()
	`
	params := map[string]interface{}{
		"personID":     person.ID,
		"departmentID": departmentID,
	}

	_, err := neo4j.ExecuteQuery(
		p.ctx,
		*p.db,
		query,
		params,
		neo4j.EagerResultTransformer,
	)
	if err != nil {
		return err
	}

	return nil
}

func (p PersonRelRepositoryImpl) RemoveDepartmentFromPerson(person dao.Person, departmentID string) error {
	/* Removes a department from a person
	   @param person: The person to remove the department from
	   @param departmentID: The name of the department to remove
	*/

	query := `
	MATCH (p: Person {id: $personID}) -[r:WORKS_AT]-> (d: Department {id: $departmentID})
	DELETE r
	`
	params := map[string]interface{}{
		"personID":     person.ID,
		"departmentID": departmentID,
	}

	_, err := neo4j.ExecuteQuery(
		p.ctx,
		*p.db,
		query,
		params,
		neo4j.EagerResultTransformer,
	)
	if err != nil {
		return err
	}

	return nil
}

func (p PersonRelRepositoryImpl) AddWorkplaceToPerson(person dao.Person, departmentID string, workplaceID string) error {
	/* Adds a workplace to a person
	   @param person: The person to add the workplace to
	   @param workplaceID: The name of the workplace to add
	*/

	query := `
	MATCH (p: Person {id: $personID})
	MATCH (d: Department {id: $departmentID}) -[:HAS_WORKPLACE]-> (w: Workplace {id: $workplaceID})
	MATCH (p) -[:WORKS_AT]-> (d)
	MERGE (p) -[r:QUALIFIED_FOR]-> (w)
	ON CREATE SET r.created_at = datetime()
	`

	params := map[string]interface{}{
		"personID":     person.ID,
		"departmentID": departmentID,
		"workplaceID":  workplaceID,
	}

	_, err := neo4j.ExecuteQuery(
		p.ctx,
		*p.db,
		query,
		params,
		neo4j.EagerResultTransformer,
	)
	if err != nil {
		return err
	}

	return nil
}

func (p PersonRelRepositoryImpl) RemoveWorkplaceFromPerson(person dao.Person, departmentID string, workplaceID string) error {
	/* Removes a workplace from a person
	   @param person: The person to remove the workplace from
	   @param workplaceID: The name of the workplace to remove
	*/

	query := `
	MATCH (p: Person {id: $personID}) -[r:QUALIFIED_FOR]-> (w: Workplace {id: $workplaceID}) <-[:HAS_WORKPLACE]- (d: Department {id: $departmentID})
	DELETE r
	`
	params := map[string]interface{}{
		"personID":     person.ID,
		"departmentID": departmentID,
		"workplaceID":  workplaceID,
	}

	_, err := neo4j.ExecuteQuery(
		p.ctx,
		*p.db,
		query,
		params,
		neo4j.EagerResultTransformer,
	)
	if err != nil {
		return err
	}

	return nil
}

func (p PersonRelRepositoryImpl) AddWeekdayToPerson(person dao.Person, weekdayID string) error {
	/* Adds a weekday to a person
	   @param person: The person to add the weekday to
	   @param weekdayID: The ID of the weekday to add
	*/

	query := `
	MATCH (p: Person {id: $personID})
	MATCH (wd: Weekday {id: $weekdayID})
	MERGE (p) -[r:AVAILABLE_ON]-> (wd)
	ON CREATE SET r.created_at = datetime()
	`
	params := map[string]interface{}{
		"personID":  person.ID,
		"weekdayID": weekdayID,
	}

	_, err := neo4j.ExecuteQuery(
		p.ctx,
		*p.db,
		query,
		params,
		neo4j.EagerResultTransformer,
	)

	if err != nil {
		return err
	}

	return nil
}

func (p PersonRelRepositoryImpl) RemoveWeekdayFromPerson(person dao.Person, weekdayID string) error {
	/* Removes a weekday from a person
	   @param person: The person to remove the weekday from
	   @param weekdayID: The ID of the weekday to remove
	*/

	query := `
	MATCH (p: Person {id: $personID}) -[r:AVAILABLE_ON]-> (wd: Weekday {id: $weekdayID})
	DELETE r
	`
	params := map[string]interface{}{
		"personID":  person.ID,
		"weekdayID": weekdayID,
	}

	_, err := neo4j.ExecuteQuery(
		p.ctx,
		*p.db,
		query,
		params,
		neo4j.EagerResultTransformer,
	)

	if err != nil {
		return err
	}

	return nil
}

func PersonRelRepositoryInit(db *neo4j.DriverWithContext, ctx context.Context) *PersonRelRepositoryImpl {
	return &PersonRelRepositoryImpl{
		db:  db,
		ctx: ctx,
	}
}

var personRelRepositorySet = wire.NewSet(
	PersonRelRepositoryInit,
	wire.Bind(new(PersonRelRepository), new(*PersonRelRepositoryImpl)),
)
