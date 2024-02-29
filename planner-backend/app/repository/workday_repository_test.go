package repository

import (
	"context"
	"errors"
	"fmt"
	"planner-backend/app/domain/dao"
	"testing"
	"time"
)

func TestWorkdayRepositorySave(t *testing.T) {
	timeslotCreatorOne := TimeslotCreatorImpl{
		departmentID:   "dept1",
		departmentName: "Department 1",
		workplaceID:    "wp1",
		workplaceName:  "Workplace 1",
		id:             "ts1",
		name:           "Timeslot 1",
		weekdays: []struct {
			id        int64
			startTime time.Time
			endTime   time.Time
		}{
			{id: 1, startTime: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC), endTime: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)},
			{id: 2, startTime: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC), endTime: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)},
			{id: 3, startTime: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC), endTime: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)},
			{id: 4, startTime: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC), endTime: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)},
			{id: 5, startTime: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC), endTime: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)},
		},
	}
	timeslotCreatorTwo := TimeslotCreatorImpl{
		departmentID:   "dept1",
		departmentName: "Department 1",
		workplaceID:    "wp1",
		workplaceName:  "Workplace 1",
		id:             "ts2",
		name:           "Timeslot 2",
		weekdays: []struct {
			id        int64
			startTime time.Time
			endTime   time.Time
		}{
			{id: 1, startTime: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC), endTime: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)},
			{id: 2, startTime: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC), endTime: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)},
			{id: 3, startTime: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC), endTime: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)},
			{id: 4, startTime: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC), endTime: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)},
			{id: 5, startTime: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC), endTime: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)},
			{id: 6, startTime: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC), endTime: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)},
			{id: 7, startTime: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC), endTime: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)},
		},
	}
	timeslotCreatorThree := TimeslotCreatorImpl{
		departmentID:   "dept1",
		departmentName: "Department 1",
		workplaceID:    "wp2",
		workplaceName:  "Workplace 2",
		id:             "ts1",
		name:           "Timeslot 1",
		weekdays: []struct {
			id        int64
			startTime time.Time
			endTime   time.Time
		}{
			{id: 1, startTime: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC), endTime: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)},
			{id: 2, startTime: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC), endTime: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)},
			{id: 3, startTime: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC), endTime: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)},
			{id: 4, startTime: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC), endTime: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)},
			{id: 5, startTime: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC), endTime: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)},
			{id: 6, startTime: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC), endTime: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)},
			{id: 7, startTime: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC), endTime: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)},
		},
	}
	tests := []struct {
		name          string
		workday       *dao.Workday
		expectedError error
	}{
		{
			name: "Date does not exist",
			workday: &dao.Workday{
				Department: dao.Department{
					ID: "dept1",
				},
				Workplace: dao.Workplace{
					ID: "wp1",
				},
				Timeslot: dao.Timeslot{
					ID: "ts1",
				},
				Date:      "1999-03-22",
				StartTime: "08:00",
				EndTime:   "16:00",
				Active:    true,
			},
			expectedError: errors.New("Workplace does not exist"),
		},
		{
			name: "Timeslot does not exist",
			workday: &dao.Workday{
				Department: dao.Department{
					ID: "dept1",
				},
				Workplace: dao.Workplace{
					ID: "wp1",
				},
				Timeslot: dao.Timeslot{
					ID: "badTs",
				},
				Date:      time.Now().Format("2006-01-02"), // time.Now().Format("2006-01-02") returns the current date in the format "YYYY-MM-DD"
				StartTime: "08:00",
				EndTime:   "16:00",
				Active:    true,
			},
			expectedError: errors.New("Timeslot does not exist"),
		},
		{
			name: "Workplace does not exist",
			workday: &dao.Workday{
				Department: dao.Department{
					ID: "dept1",
				},
				Workplace: dao.Workplace{
					ID: "badWp",
				},
				Timeslot: dao.Timeslot{
					ID: "ts1",
				},
				Date:      time.Now().Format("2006-01-02"), // time.Now().Format("2006-01-02") returns the current date in the format "YYYY-MM-DD"
				StartTime: "08:00",
				EndTime:   "16:00",
				Active:    true,
			},
			expectedError: errors.New("Workplace does not exist"),
		},
		{
			name: "Department does not exist",
			workday: &dao.Workday{
				Department: dao.Department{
					ID: "badDept",
				},
				Workplace: dao.Workplace{
					ID: "wp1",
				},
				Timeslot: dao.Timeslot{
					ID: "ts1",
				},
				Date:      time.Now().Format("2006-01-02"), // time.Now().Format("2006-01-02") returns the current date in the format "YYYY-MM-DD"
				StartTime: "08:00",
				EndTime:   "16:00",
				Active:    true,
			},
			expectedError: errors.New("Department does not exist"),
		},
		{
			name: "Should save workday with active false",
			workday: &dao.Workday{
				Department: dao.Department{
					ID: "dept1",
				},
				Workplace: dao.Workplace{
					ID: "wp1",
				},
				Timeslot: dao.Timeslot{
					ID: "ts1",
				},
				Date:      time.Now().Format("2006-01-02"), // time.Now().Format("2006-01-02") returns the current date in the format "YYYY-MM-DD"
				StartTime: "08:00",
				EndTime:   "16:00",
				Active:    false,
			},
			expectedError: nil,
		},
		{
			name: "Should save workday without comment",
			workday: &dao.Workday{
				Department: dao.Department{
					ID: "dept1",
				},
				Workplace: dao.Workplace{
					ID: "wp1",
				},
				Timeslot: dao.Timeslot{
					ID: "ts1",
				},
				Date:      time.Now().Format("2006-01-02"), // time.Now().Format("2006-01-02") returns the current date in the format "YYYY-MM-DD"
				StartTime: "08:00",
				EndTime:   "16:00",
				Active:    true,
			},
			expectedError: nil,
		},
		{
			name: "Should save workday with comment",
			workday: &dao.Workday{
				Department: dao.Department{
					ID: "dept1",
				},
				Workplace: dao.Workplace{
					ID: "wp1",
				},
				Timeslot: dao.Timeslot{
					ID: "ts1",
				},
				Date:      time.Now().Format("2006-01-02"), // time.Now().Format("2006-01-02") returns the current date in the format "YYYY-MM-DD"
				StartTime: "08:00",
				EndTime:   "16:00",
				Active:    true,
				Comment:   "comment",
			},
			expectedError: nil,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d: %s", i, test.name), func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			db, err := NewTestDBInstance(ctx)
			if err != nil {
				t.Errorf("Error creating test database: %v", err)
			}

			// setup initial state
			Migrate(ctx, db)
			timeslotCreatorOne.Create(db, ctx)
			timeslotCreatorTwo.Create(db, ctx)
			timeslotCreatorThree.Create(db, ctx)

			s := SynchronizeRepositoryImpl{
				ctx: ctx,
				db:  db,
			}
			if err := s.Synchronize(2); err != nil {
				t.Errorf("Error synchronizing database: %v", err)
			}

			if err := timeslotCreatorOne.Create(db, ctx); err != nil {
				t.Errorf("Error creating timeslots: %v", err)
			}

			// Run the test
			w := WorkdayRepositoryImpl{
				db:  db,
				ctx: ctx,
			}

			err = w.Save(test.workday)

			// Check the result
			if test.expectedError == nil && err != nil {
				t.Errorf("Expected no error, but got %v", err)
			}

			if test.expectedError != nil && err == nil {
				t.Errorf("Expected error %v, but got nil", test.expectedError)
			}
			if test.expectedError != nil {
				return
			}

			res, err := w.GetWorkday(test.workday.Department.ID, test.workday.Workplace.ID, test.workday.Timeslot.ID, test.workday.Date)
			if err != nil {
				t.Errorf("Error getting workday: %v", err)
			}

			if res.StartTime != test.workday.StartTime {
				t.Errorf("Expected start time %s, but got %s", test.workday.StartTime, res.StartTime)
			}

			if res.EndTime != test.workday.EndTime {
				t.Errorf("Expected end time %s, but got %s", test.workday.EndTime, res.EndTime)
			}

			if res.Active != test.workday.Active {
				t.Errorf("Expected active %t, but got %t", test.workday.Active, res.Active)
			}

			if res.Comment != test.workday.Comment {
				t.Errorf("Expected comment %s, but got %s", test.workday.Comment, res.Comment)
			}
		})
	}
}

func TestAssignOnePersonToWorkday(t *testing.T) {
	timeslotCreatorOne := TimeslotCreatorImpl{
		departmentID:   "dept1",
		departmentName: "Department 1",
		workplaceID:    "wp1",
		workplaceName:  "Workplace 1",
		id:             "ts1",
		name:           "Timeslot 1",
		weekdays: []struct {
			id        int64
			startTime time.Time
			endTime   time.Time
		}{
			{id: 1, startTime: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC), endTime: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)},
			{id: 2, startTime: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC), endTime: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)},
			{id: 3, startTime: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC), endTime: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)},
			{id: 4, startTime: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC), endTime: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)},
			{id: 5, startTime: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC), endTime: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)},
		},
	}
	timeslotCreatorTwo := TimeslotCreatorImpl{
		departmentID:   "dept1",
		departmentName: "Department 1",
		workplaceID:    "wp1",
		workplaceName:  "Workplace 1",
		id:             "ts2",
		name:           "Timeslot 2",
		weekdays: []struct {
			id        int64
			startTime time.Time
			endTime   time.Time
		}{
			{id: 1, startTime: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC), endTime: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)},
			{id: 2, startTime: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC), endTime: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)},
			{id: 3, startTime: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC), endTime: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)},
			{id: 4, startTime: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC), endTime: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)},
			{id: 5, startTime: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC), endTime: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)},
			{id: 6, startTime: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC), endTime: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)},
			{id: 7, startTime: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC), endTime: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)},
		},
	}
	timeslotCreatorThree := TimeslotCreatorImpl{
		departmentID:   "dept1",
		departmentName: "Department 1",
		workplaceID:    "wp2",
		workplaceName:  "Workplace 2",
		id:             "ts1",
		name:           "Timeslot 1",
		weekdays: []struct {
			id        int64
			startTime time.Time
			endTime   time.Time
		}{
			{id: 1, startTime: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC), endTime: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)},
			{id: 2, startTime: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC), endTime: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)},
			{id: 3, startTime: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC), endTime: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)},
			{id: 4, startTime: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC), endTime: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)},
			{id: 5, startTime: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC), endTime: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)},
			{id: 6, startTime: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC), endTime: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)},
			{id: 7, startTime: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC), endTime: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)}},
	}

	personToCreateOne := PersonCreatorImpl{
		departments: []struct {
			id   string
			name string
		}{
			{
				id:   "dept1",
				name: "Department 1",
			},
		},
		workplaces: []struct {
			id           string
			name         string
			departmentID string
		}{
			{
				id:           "wp1",
				name:         "Workplace 1",
				departmentID: "dept1",
			},
			{
				id:           "wp2",
				name:         "Workplace 2",
				departmentID: "dept1",
			},
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
			email:        "person@example.com",
			active:       true,
			lastName:     "Person",
			firstName:    "1",
			workingHours: 40,
		},
	}
	personToCreateTwo := PersonCreatorImpl{
		departments: []struct {
			id   string
			name string
		}{
			{
				id:   "dept1",
				name: "Department 1",
			},
		},
		workplaces: []struct {
			id           string
			name         string
			departmentID string
		}{
			{
				id:           "wp1",
				name:         "Workplace 1",
				departmentID: "dept1",
			},
			{
				id:           "wp2",
				name:         "Workplace 2",
				departmentID: "dept1",
			},
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
			id:           "person2",
			email:        "test@example.com",
			active:       true,
			lastName:     "Person",
			firstName:    "2",
			workingHours: 40,
		},
	}

	personToCreateThree := PersonCreatorImpl{
		departments: []struct {
			id   string
			name string
		}{
			{
				id:   "dept2",
				name: "Department 2",
			},
		},
		workplaces: []struct {
			id           string
			name         string
			departmentID string
		}{
			{
				id:           "wp3",
				name:         "Workplace 3",
				departmentID: "dept2",
			},
			{
				id:           "wp4",
				name:         "Workplace 4",
				departmentID: "dept2",
			},
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
			id:           "person3",
			email:        "person3@example.com",
			active:       true,
			lastName:     "Person",
			firstName:    "3",
			workingHours: 40,
		},
	}

	tests := []struct {
		name            string
		workday         *dao.Workday
		personsToAssign []struct {
			ID          string
			expectError bool
		}
		expectedError    error
		checkValues      bool
		shouldGetWorkday bool
	}{
		{
			name: "Should assign person to workday",
			workday: &dao.Workday{
				Department: dao.Department{
					ID: "dept1",
				},
				Workplace: dao.Workplace{
					ID: "wp1",
				},
				Timeslot: dao.Timeslot{
					ID: "ts1",
				},
				Date:      time.Now().Format("2006-01-02"), // time.Now().Format("2006-01-02") returns the current date in the format "YYYY-MM-DD"
				StartTime: "08:00",
				EndTime:   "16:00",
				Active:    false,
			},
			personsToAssign: []struct {
				ID          string
				expectError bool
			}{
				{
					ID:          "person1",
					expectError: false,
				},
			},
			checkValues:      true,
			shouldGetWorkday: true,
		},
		{
			name: "Should not assign non-existing person to workday",
			workday: &dao.Workday{
				Department: dao.Department{
					ID: "dept1",
				},
				Workplace: dao.Workplace{
					ID: "wp1",
				},
				Timeslot: dao.Timeslot{
					ID: "ts1",
				},
				Date:      time.Now().Format("2006-01-02"), // time.Now().Format("2006-01-02") returns the current date in the format "YYYY-MM-DD"
				StartTime: "08:00",
				EndTime:   "16:00",
				Active:    true,
			},
			personsToAssign: []struct {
				ID          string
				expectError bool
			}{
				{
					ID:          "badPerson",
					expectError: true,
				},
			},
			checkValues:      false,
			shouldGetWorkday: true,
		},
		{
			name: "Should not assign person to non-existing workday",
			workday: &dao.Workday{
				Department: dao.Department{
					ID: "dept1",
				},
				Workplace: dao.Workplace{
					ID: "wp1",
				},
				Timeslot: dao.Timeslot{
					ID: "ts1",
				},
				Date:      "1999-03-22",
				StartTime: "08:00",
				EndTime:   "16:00",
				Active:    true,
			},
			personsToAssign: []struct {
				ID          string
				expectError bool
			}{
				{
					ID:          "person1",
					expectError: true,
				},
			},
			checkValues:      false,
			shouldGetWorkday: false,
		},
		{
			name: "Should not assign one person from another department to workday",
			workday: &dao.Workday{
				Department: dao.Department{
					ID: "dept1",
				},
				Workplace: dao.Workplace{
					ID: "wp1",
				},
				Timeslot: dao.Timeslot{
					ID: "ts1",
				},
				Date:      time.Now().Format("2006-01-02"), // time.Now().Format("2006-01-02") returns the current date in the format "YYYY-MM-DD"
				StartTime: "08:00",
				EndTime:   "16:00",
				Active:    true,
			},
			personsToAssign: []struct {
				ID          string
				expectError bool
			}{
				{
					ID:          "person3",
					expectError: true,
				},
			},
			checkValues:      false,
			shouldGetWorkday: true,
		},
		{
			name: "should assign multiple persons to workday",
			workday: &dao.Workday{
				Department: dao.Department{
					ID: "dept1",
				},
				Workplace: dao.Workplace{
					ID: "wp1",
				},
				Timeslot: dao.Timeslot{
					ID: "ts1",
				},
				Date:      time.Now().Format("2006-01-02"), // time.Now().Format("2006-01-02") returns the current date in the format "YYYY-MM-DD"
				StartTime: "08:00",
				EndTime:   "16:00",
				Active:    true,
			},
			personsToAssign: []struct {
				ID          string
				expectError bool
			}{
				{
					ID:          "person1",
					expectError: false,
				},
				{
					ID:          "person2",
					expectError: false,
				},
			},
			checkValues:      true,
			shouldGetWorkday: true,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d: %s", i, test.name), func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			db, err := NewTestDBInstance(ctx)
			if err != nil {
				t.Errorf("Error creating test database: %v", err)
			}

			// setup initial state
			Migrate(ctx, db)
			timeslotCreatorOne.Create(db, ctx)
			timeslotCreatorTwo.Create(db, ctx)
			timeslotCreatorThree.Create(db, ctx)
			personToCreateOne.Create(db, ctx)
			personToCreateTwo.Create(db, ctx)
			personToCreateThree.Create(db, ctx)

			s := SynchronizeRepositoryImpl{
				ctx: ctx,
				db:  db,
			}
			if err := s.Synchronize(2); err != nil {
				t.Errorf("Error synchronizing database: %v", err)
			}

			if err := timeslotCreatorOne.Create(db, ctx); err != nil {
				t.Errorf("Error creating timeslots: %v", err)
			}

			// Run the test
			w := WorkdayRepositoryImpl{
				db:  db,
				ctx: ctx,
			}

			// Assigned person on workday should not exist
			wd, err := w.GetWorkday(test.workday.Department.ID, test.workday.Workplace.ID, test.workday.Timeslot.ID, test.workday.Date)
			if err != nil && test.shouldGetWorkday {
				t.Errorf("Error getting workday: %v", err)
			} else {
				return
			}
			if len(wd.Persons) > 0 {
				t.Errorf("Expected no person assigned to workday, but got %d", len(wd.Persons))
			}

			for _, person := range test.personsToAssign {
				err = w.AssignPersonToWorkday(person.ID, test.workday.Department.ID, test.workday.Workplace.ID, test.workday.Timeslot.ID, test.workday.Date)
				if person.expectError && err == nil {
					t.Errorf("Expected no error, but got %v", err)
				}
				if !person.expectError && err != nil {
					t.Errorf("Expected error %v, but got nil", test.expectedError)
				}
			}

			if !test.checkValues {
				return
			}

			// Check the result
			res, err := w.GetWorkday(test.workday.Department.ID, test.workday.Workplace.ID, test.workday.Timeslot.ID, test.workday.Date)
			if err != nil {
				t.Errorf("Error getting workday: %v", err)
			}

			if len(res.Persons) != len(test.personsToAssign) {
				t.Errorf("Expected %d persons assigned to workday, but got %d", len(test.personsToAssign), len(res.Persons))
			}

			for _, person := range test.personsToAssign {
				found := false
				for _, p := range res.Persons {
					if p.ID == person.ID {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("Expected person %s assigned to workday, but not found", person.ID)
				}
			}
		})
	}
}
