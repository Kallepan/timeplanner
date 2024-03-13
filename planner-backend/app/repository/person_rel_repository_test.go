package repository

import (
	"context"
	"fmt"
	"planner-backend/app/domain/dao"
	"testing"
)

func TestAddAndRemoveAbsencyToPerson(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	db, err := NewTestDBInstance(ctx)
	if err != nil {
		t.Errorf("Error creating test database: %v", err)
	}
	defer cancel()
	Migrate(ctx, db)

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
		weekdayIDs: []int64{1, 2, 3, 4, 5, 6, 7},
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

	p := PersonRelRepositoryImpl{
		db:  db,
		ctx: ctx,
	}

	tests := []struct {
		name          string
		person        dao.Person
		absency       dao.Absence
		expectedError bool
		expectedSave  bool
	}{
		{
			name: "Add Absency no person",
			person: dao.Person{
				ID: "noop",
			},
			absency: dao.Absence{
				Date:   "2021-01-01",
				Reason: "Test 1 Reason",
			},
			expectedError: false,
			expectedSave:  false,
		},
		{
			name: "Add Absency wrong date",
			person: dao.Person{
				ID: "person1",
			},
			absency: dao.Absence{
				Date:   "2021-011-01",
				Reason: "Test 1 Reason",
			},
			expectedError: true,
			expectedSave:  false,
		},
		{
			name: "Add Absency normal",
			person: dao.Person{
				ID: "person1",
			},
			absency: dao.Absence{
				Date:   "2021-01-01",
				Reason: "Test 1 Reason",
			},
			expectedError: false,
			expectedSave:  true,
		},
		{
			name: "Add Absency no reason",
			person: dao.Person{
				ID: "person1",
			},
			absency: dao.Absence{
				Date:   "2021-01-01",
				Reason: "",
			},
			expectedError: false,
			expectedSave:  true,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Test %s: %d", test.name, i), func(t *testing.T) {

			err := p.AddAbsencyToPerson(test.person, test.absency)
			if err != nil && !test.expectedError {
				t.Errorf("Expected no error, got %v", err)
			}

			if err == nil && test.expectedError {
				t.Errorf("Expected error, got no error")
			}

			if test.expectedError || !test.expectedSave {
				return
			}

			// check if the absency was added
			absency, err := p.FindAbsencyForPerson(test.person.ID, test.absency.Date)
			if err != nil {
				t.Errorf("Error finding absencies: %v", err)
			}

			if absency.Date != test.absency.Date {
				t.Errorf("Expected %s, got %s", test.absency.Date, absency.Date)
			}

			if absency.Reason != test.absency.Reason {
				t.Errorf("Expected %s, got %s", test.absency.Reason, absency.Reason)
			}
		})
	}
}

func TestRemoveAbsencyFromPerson(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	db, err := NewTestDBInstance(ctx)
	if err != nil {
		t.Errorf("Error creating test database: %v", err)
	}
	defer cancel()
	Migrate(ctx, db)

	p := PersonRelRepositoryImpl{
		db:  db,
		ctx: ctx,
	}

	// tests
	tests := []struct {
		name          string
		person        dao.Person
		absency       dao.Absence
		expectedError bool
		addAbsency    bool
	}{
		{
			name: "Delete Absency normal",
			person: dao.Person{
				ID: "p1",
			},
			absency: dao.Absence{
				Date:   "2021-01-01",
				Reason: "Test 1 Reason",
			},
			expectedError: false,
			addAbsency:    true,
		},
		{
			name: "Delete Absency no reason",
			person: dao.Person{
				ID: "p1",
			},
			absency: dao.Absence{
				Date:   "2021-01-01",
				Reason: "",
			},
			expectedError: false,
			addAbsency:    true,
		},
		{
			name: "Delete Absency not present",
			person: dao.Person{
				ID: "p2",
			},
			absency: dao.Absence{
				Date:   "2021-01-01",
				Reason: "Test 1 Reason",
			},
			expectedError: false,
			addAbsency:    false,
		},
		{
			name: "Delete Absency no person",
			person: dao.Person{
				ID: "noop",
			},
			absency: dao.Absence{
				Date:   "2021-01-01",
				Reason: "Test 1 Reason",
			},
			expectedError: false,
			addAbsency:    false,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Test %s: %d", test.name, i), func(t *testing.T) {
			// Add the absency
			if test.addAbsency {
				err := p.AddAbsencyToPerson(test.person, test.absency)
				if err != nil {
					t.Errorf("Error adding absency: %v", err)
				}
			}

			err := p.RemoveAbsencyFromPerson(test.person, test.absency)
			if err != nil && !test.expectedError {
				t.Errorf("Expected no error, got %v", err)
			}

			if err == nil && test.expectedError {
				t.Errorf("Expected error, got no error")
			}

			if test.expectedError || !test.addAbsency {
				return
			}

			// check if the absency was removed
			absency, err := p.FindAbsencyForPerson(test.person.ID, test.absency.Date)
			if err == nil {
				t.Errorf("Expected error, got no error")
			}

			if absency.Date != "" {
				t.Errorf("Expected empty, got %s", absency.Date)
			}
			if absency.Reason != "" {
				t.Errorf("Expected empty, got %s", absency.Reason)
			}
		})
	}
}
