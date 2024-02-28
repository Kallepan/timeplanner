package repository

import (
	"context"
	"planner-backend/app/domain/dao"
	"planner-backend/app/pkg"

	"github.com/google/wire"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type PersonRepository interface {
	// Function Used by the service
	FindAllPersons(departmentID string) ([]dao.Person, error)
	FindAllPersonsBy(departmentID string, workplaceID string, weekdayID string, notAbsentOn string) ([]dao.Person, error)
	FindPersonByID(personID string) (dao.Person, error)
	Save(person *dao.Person) (dao.Person, error)
	Delete(person *dao.Person) error
}

type PersonRepositoryImpl struct {
	db  *neo4j.DriverWithContext
	ctx context.Context
}

func (p PersonRepositoryImpl) FindAllPersonsBy(departmentID string, workplaceID string, weekdayID string, notAbsentDate string) ([]dao.Person, error) {
	/* Returns all persons in a department, qualified for a workplace that are present on a weekday and not absent on a date
	   @param departmentID: The name of the department the persons should be in
	   @param presentOnWeekdayID: The ID of the weekday the persons should be present on
	   @param workplaceID: The name of the workplace the person should be qualified for
	   @param notAbsentOn: The date the person should not be absent on
	*/

	persons := []dao.Person{}

	// Build dynamic query depending on which param was given
	params := map[string]interface{}{
		"departmentID": departmentID,
	}

	query := `MATCH (p: Person) 
    MATCH (p)-[:WORKS_AT]->(d: Department {id: $departmentID})`

	if workplaceID != "" {
		query += ` MATCH (p) -[:QUALIFIED_FOR]-> (w: Workplace {id: $workplaceID})`
		params["workplaceID"] = workplaceID
	}
	if weekdayID != "" {
		query += ` MATCH (p) -[:AVAILABLE_ON]->(wd: Weekday {id: $weekdayID})`
		params["weekdayID"] = weekdayID
	}

	query += ` WHERE p.deleted_at IS NULL AND p.active = true`
	if notAbsentDate != "" {
		query += ` AND NOT EXISTS((p) -[:ABSENT_ON]-> (d:Date {date($notAbsentDate)}))`
		params["notAbsentDate"] = notAbsentDate
	}

	query += ` RETURN p`

	result, err := neo4j.ExecuteQuery(
		p.ctx,
		*p.db,
		query,
		params,
		neo4j.EagerResultTransformer,
	)
	if err != nil {
		return nil, err
	}

	for _, record := range result.Records {
		person := dao.Person{}
		if err := person.ParseFromDBRecord(record); err != nil {
			return nil, err
		}

		persons = append(persons, person)
	}

	return persons, nil
}

func (p PersonRepositoryImpl) FindAllPersons(departmentID string) ([]dao.Person, error) {
	/* Returns all persons
	   @param departmentID: The name of the department the persons should be in
	*/

	persons := []dao.Person{}
	params := map[string]interface{}{
		"departmentID": departmentID,
	}
	query := `
    MATCH (p: Person)
    MATCH (p)-[:WORKS_AT]->(d: Department {id: $departmentID})
	OPTIONAL MATCH (p)-[:QUALIFIED_FOR]->(w: Workplace) <-[:HAS_WORKPLACE]-(d2: Department)
	OPTIONAL MATCH (p)-[:AVAILABLE_ON]->(wd: Weekday)
    WHERE p.deleted_at IS NULL AND p.active = true
    RETURN 
		p, 
		COLLECT(DISTINCT { 
			id: d.id, 
			name: d.name
		}) AS departments, 
		COLLECT(DISTINCT {
			id: w.id,
			name: w.name,
			department_id: d2.id
		}) AS workplaces, COLLECT(DISTINCT {
			id: wd.id,
			name: wd.name
		}) AS weekdays`

	result, err := neo4j.ExecuteQuery(
		p.ctx,
		*p.db,
		query,
		params,
		neo4j.EagerResultTransformer,
	)
	if err != nil {
		return nil, err
	}

	for _, record := range result.Records {
		person := dao.Person{}
		if err := person.ParseFromDBRecord(record); err != nil {
			return nil, err
		}

		if err := person.ParseAdditionalFieldsFromDBRecord(record); err != nil {
			return nil, err
		}

		persons = append(persons, person)
	}

	return persons, nil
}

func (p PersonRepositoryImpl) FindPersonByID(personID string) (dao.Person, error) {
	/* Returns a person by name */

	person := dao.Person{}
	query := `
	MATCH (p:Person {id: $personID})
	OPTIONAL MATCH (p)-[:WORKS_AT]->(d: Department)
	OPTIONAL MATCH (p)-[:QUALIFIED_FOR]->(w: Workplace) <-[:HAS_WORKPLACE]-(d2: Department)
	OPTIONAL MATCH (p)-[:AVAILABLE_ON]->(wd: Weekday)
	WHERE p.deleted_at IS NULL AND p.active = true
	RETURN 
	p, 
	COLLECT(DISTINCT { 
		id: d.id, 
		name: d.name
	}) AS departments, 
	COLLECT(DISTINCT {
		id: w.id,
		name: w.name,
		department_id: d2.id
	}) AS workplaces, COLLECT(DISTINCT {
		id: wd.id,
		name: wd.name
	}) AS weekdays`
	params := map[string]interface{}{
		"personID": personID,
	}

	result, err := neo4j.ExecuteQuery(
		p.ctx,
		*p.db,
		query,
		params,
		neo4j.EagerResultTransformer,
	)
	if err != nil {
		return dao.Person{}, err
	}

	if len(result.Records) == 0 {
		return dao.Person{}, pkg.ErrNoRows
	}

	if err := person.ParseFromDBRecord(result.Records[0]); err != nil {
		return dao.Person{}, err
	}
	if err := person.ParseAdditionalFieldsFromDBRecord(result.Records[0]); err != nil {
		return dao.Person{}, err
	}

	return person, nil
}

func (p PersonRepositoryImpl) Save(person *dao.Person) (dao.Person, error) {
	/* Saves a person */

	query := `
    MERGE (p:Person {id: $personID})
    ON CREATE SET
        p.firstName = $firstName,
        p.lastName = $lastName,
        p.email = $email,
        p.active = $active,
        p.workingHours = $workingHours,
        p.created_at = datetime(),
        p.updated_at = datetime()
    ON MATCH SET
        p.firstName = $firstName,
        p.lastName = $lastName,
        p.email = $email,
        p.active = $active,
        p.workingHours = $workingHours,
        p.updated_at = datetime(),
		p.deleted_at = NULL
    RETURN p`
	params := map[string]interface{}{
		"personID":     person.ID,
		"firstName":    person.FirstName,
		"lastName":     person.LastName,
		"email":        person.Email,
		"active":       person.Active,
		"workingHours": person.WorkingHours,
	}

	result, err := neo4j.ExecuteQuery(
		p.ctx,
		*p.db,
		query,
		params,
		neo4j.EagerResultTransformer,
	)
	if err != nil {
		return dao.Person{}, err
	}

	if len(result.Records) == 0 {
		return dao.Person{}, pkg.ErrNoRows
	}

	if err := person.ParseFromDBRecord(result.Records[0]); err != nil {
		return dao.Person{}, err
	}

	return *person, nil
}

func (p PersonRepositoryImpl) Delete(person *dao.Person) error {
	/* Deletes a person */

	query := `
    MATCH  (p:Person {id: $personID})
    SET p.deleted_at = datetime()`
	params := map[string]interface{}{
		"personID": person.ID,
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

func PersonRepositoryInit(db *neo4j.DriverWithContext, ctx context.Context) *PersonRepositoryImpl {
	return &PersonRepositoryImpl{
		db:  db,
		ctx: ctx,
	}
}

var personRepositorySet = wire.NewSet(
	PersonRepositoryInit,
	wire.Bind(new(PersonRepository), new(*PersonRepositoryImpl)),
)
