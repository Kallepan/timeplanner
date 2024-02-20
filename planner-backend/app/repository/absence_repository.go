package repository

import (
	"context"
	"errors"
	"planner-backend/app/domain/dao"

	"github.com/google/wire"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type AbsenceRepository interface {
	FindAllAbsencies(departmentID string, date string) ([]dao.Absence, error)
}

type AbsenceRepositoryImpl struct {
	db  *neo4j.DriverWithContext
	ctx context.Context
}

func (a AbsenceRepositoryImpl) FindAllAbsencies(departmentID string, date string) ([]dao.Absence, error) {
	/* FindAllAbsencies is a function to get all absencies for a given department and date
	 * @param departmentID is the department id
	 * @param date is the date
	 * @return []dao.Absence, error
	 */

	query := `
    MATCH (d: Department {id: $departmentID}) <-[:WORKS_AT]- (p: Person) -[r:ABSENT_ON]-> (date: Date {date: date($date)})
    RETURN p.id, date.date, r`
	params := map[string]interface{}{
		"departmentID": departmentID,
		"date":         date,
	}

	result, err := neo4j.ExecuteQuery(
		a.ctx,
		*a.db,
		query,
		params,
		neo4j.EagerResultTransformer,
	)
	if err != nil {
		return nil, err
	}

	absences := make([]dao.Absence, 0)
	if len(result.Records) == 0 {
		return absences, nil
	}

	for _, record := range result.Records {
		dbDate, ok := record.Values[1].(neo4j.Date)
		if !ok {
			return nil, errors.New("could not parse date")
		}
		date := dbDate.Time().Format("2006-01-02")

		personID, ok := record.Values[0].(string)
		if !ok {
			return nil, errors.New("could not parse person id")
		}

		absence := dao.Absence{}
		if err := absence.ParseFromDBRecord(record, date, personID); err != nil {
			return nil, err
		}

		absences = append(absences, absence)
	}

	return absences, nil
}

func AbsenceRepositoryInit(db *neo4j.DriverWithContext, ctx context.Context) AbsenceRepositoryImpl {
	return AbsenceRepositoryImpl{
		db:  db,
		ctx: ctx,
	}
}

var absenceRepositorySet = wire.NewSet(
	AbsenceRepositoryInit,
	wire.Bind(new(AbsenceRepository), new(AbsenceRepositoryImpl)),
)
