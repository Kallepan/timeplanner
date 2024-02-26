package repository

import (
	"context"
	"fmt"
	"planner-backend/app/domain/dao"
	"testing"
)

func TestGetAllAbsenciesFromDepartment(t *testing.T) {
	// Create a new test database instance
	ctx, cancel := context.WithCancel(context.Background())
	db, err := NewTestDBInstance(ctx)
	if err != nil {
		t.Errorf("Error creating test database: %v", err)
	}
	defer cancel()
	Migrate(ctx, db)

	// setup the initial state
	// create a person
	personCreator := PersonCreatorImpl{
		departments: []struct {
			id   string
			name string
		}{
			{id: "dept1", name: "Department 1"},
			{id: "dept2", name: "Department 2"},
		},
		workplaces: []struct {
			id           string
			name         string
			departmentID string
		}{
			{id: "wp1", name: "Workplace 1", departmentID: "dept1"},
			{id: "wp2", name: "Workplace 2", departmentID: "dept2"},
		},
		weekdayIDs: []string{"MON", "TUE", "WED", "THU", "FRI", "SAT", "SUN"},
		person: struct {
			id           string
			email        string
			active       bool
			lastName     string
			firstName    string
			workingHours float64
		}{
			id:           "person1",
			email:        "person1@example.com",
			active:       true,
			lastName:     "Doe",
			firstName:    "John",
			workingHours: 8.0,
		},
	}
	personCreator.Create(db, ctx)
	personCreator2 := PersonCreatorImpl{
		departments: []struct {
			id   string
			name string
		}{
			{id: "dept1", name: "Department 1"},
		},
		workplaces: []struct {
			id           string
			name         string
			departmentID string
		}{
			{id: "wp1", name: "Workplace 1", departmentID: "dept1"},
		},
		weekdayIDs: []string{"MON", "TUE", "WED", "THU", "FRI", "SAT", "SUN"},
		person: struct {
			id           string
			email        string
			active       bool
			lastName     string
			firstName    string
			workingHours float64
		}{
			id:           "person2",
			email:        "person2@example.com",
			active:       true,
			lastName:     "Doe",
			firstName:    "Jane",
			workingHours: 8.0,
		},
	}
	personCreator2.Create(db, ctx)
	departmentCreator := DeparmentCreatorImpl{
		name: "Department 3",
		id:   "dept3",
	}
	departmentCreator.Create(db, ctx)

	p := PersonRelRepositoryImpl{
		db:  db,
		ctx: ctx,
	}
	a := AbsenceRepositoryImpl{
		db:  db,
		ctx: ctx,
	}

	type AbsencyToCreate struct {
		Date         string
		Reason       string
		DepartmentID string
		PersonID     string
	}

	tests := []struct {
		name              string
		departmentID      string
		date              string
		expectedError     bool
		results           int
		absenciesToCreate []AbsencyToCreate
	}{
		{
			name:          "Get all absencies from department no absencies on date",
			departmentID:  "dept1",
			date:          "2021-01-06",
			expectedError: false,
			results:       0,
			absenciesToCreate: []AbsencyToCreate{
				{
					Date:         "2021-01-01",
					Reason:       "Test 1 Reason",
					DepartmentID: "dept1",
					PersonID:     "person1",
				},
				{
					Date:         "2021-01-05",
					Reason:       "Test 2 Reason",
					DepartmentID: "dept1",
					PersonID:     "person1",
				},
				{
					Date:         "2021-01-05",
					Reason:       "Test 3 Reason",
					DepartmentID: "dept1",
					PersonID:     "person1",
				},
			},
		},
		{
			name:          "Get all absencies from department with multiple absencies with multiple persons",
			departmentID:  "dept1",
			date:          "2021-01-05",
			expectedError: false,
			results:       2,
			absenciesToCreate: []AbsencyToCreate{
				{
					Date:         "2021-01-05",
					DepartmentID: "dept2",
					PersonID:     "person1",
				},
				{
					Date:         "2021-01-01",
					DepartmentID: "dept1",
					PersonID:     "person2",
				},
				{
					Date:         "2021-01-01",
					DepartmentID: "dept1",
					PersonID:     "person1",
				},
				{
					Date:         "2021-01-05",
					DepartmentID: "dept1",
					PersonID:     "person2",
				},
				{
					Date:         "2021-01-05",
					DepartmentID: "dept1",
					PersonID:     "person1",
				},
			},
		},
		{
			name:          "Get all absencies from department with multiple absencies",
			departmentID:  "dept1",
			date:          "2021-01-02",
			expectedError: false,
			results:       1,
			absenciesToCreate: []AbsencyToCreate{
				{
					Date:         "2021-01-01",
					Reason:       "Test 1 Reason",
					DepartmentID: "dept1",
					PersonID:     "person1",
				},
				{
					Date:         "2021-01-02",
					Reason:       "Test 2 Reason",
					DepartmentID: "dept1",
					PersonID:     "person1",
				},
				{
					Date:         "2021-01-02",
					Reason:       "Test 3 Reason",
					DepartmentID: "dept1",
					PersonID:     "person1",
				},
			},
		},
		{
			name:          "Get all absencies from department no absencies",
			departmentID:  "dept3",
			date:          "2021-01-01",
			expectedError: false,
			results:       0,
		},
		{
			name:          "Get all absencies from department no department",
			departmentID:  "noop",
			date:          "2021-01-01",
			expectedError: false,
			results:       0,
		},
	}

	// create some absencies
	for i, test := range tests {
		t.Run(fmt.Sprintf("Test %s: %d", test.name, i), func(t *testing.T) {
			for _, absency := range test.absenciesToCreate {
				person := dao.Person{
					ID: absency.PersonID,
				}
				if err := p.AddAbsencyToPerson(person, dao.Absence{
					Date:   absency.Date,
					Reason: absency.Reason,
				}); err != nil {
					t.Errorf("Error adding absency: %v", err)
				}
			}

			absencies, err := a.FindAllAbsencies(test.departmentID, test.date)
			if err != nil && !test.expectedError {
				t.Errorf("Expected no error, got %v", err)
			}

			if err == nil && test.expectedError {
				t.Errorf("Expected error, got no error")
			}

			if len(absencies) != test.results {
				t.Errorf("Expected %d results, got %d", test.results, len(absencies))
			}
		})
	}
}
