package repository

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func TestSynchronize(t *testing.T) {
	tests := []struct {
		name                  string
		weeksInAdvance        int
		expectedWorkdaysCount int
	}{
		{
			name:                  "Sync 1 week",
			weeksInAdvance:        1,
			expectedWorkdaysCount: 7,
		},
		{
			name:                  "Sync 2 weeks",
			weeksInAdvance:        2,
			expectedWorkdaysCount: 14,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d: %s", i, test.name), func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			db, err := NewTestDBInstance(ctx)
			if err != nil {
				t.Errorf("Error creating test database: %v", err)
			}
			defer cancel()
			// setup initial state
			Migrate(ctx, db)

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
				},
			}
			if err := timeslotCreatorOne.Create(db, ctx); err != nil {
				t.Errorf("Error creating timeslots: %v", err)
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
					{id: 4, startTime: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC), endTime: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)},
					{id: 5, startTime: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC), endTime: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)},
					{id: 6, startTime: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC), endTime: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)},
				},
			}
			if err := timeslotCreatorTwo.Create(db, ctx); err != nil {
				t.Errorf("Error creating timeslots: %v", err)
			}
			timeslotCreatorThree := TimeslotCreatorImpl{
				departmentID:   "dept2",
				departmentName: "Department 2",
				workplaceID:    "wp2",
				workplaceName:  "Workplace 2",
				id:             "ts3",
				name:           "Timeslot 3",

				weekdays: []struct {
					id        int64
					startTime time.Time
					endTime   time.Time
				}{
					{id: 7, startTime: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC), endTime: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)},
				},
			}
			if err := timeslotCreatorThree.Create(db, ctx); err != nil {
				t.Errorf("Error creating timeslots: %v", err)
			}

			// begin the test
			s := SynchronizeRepositoryImpl{
				db:  db,
				ctx: ctx,
			}

			if err := s.Synchronize(test.weeksInAdvance); err != nil {
				t.Errorf("Error synchronizing: %v", err)
			}

			// check if the workdays were created
			results, err := neo4j.ExecuteQuery(
				ctx,
				*db,
				"MATCH (w:Workday) RETURN w",
				nil,
				neo4j.EagerResultTransformer,
			)
			if err != nil {
				t.Errorf("Error querying workdays: %v", err)
			}
			if len(results.Records) != test.expectedWorkdaysCount {
				t.Errorf("Expected %d workdays, got %d", test.expectedWorkdaysCount, len(results.Records))
			}
		})
	}
}

func TestEnsureDateExistsInService(t *testing.T) {
	tests := []struct {
		name      string
		date      time.Time
		weekdayID int64
	}{
		{
			name:      "Ensure date exists",
			date:      time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			weekdayID: 5,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d: %s", i, test.name), func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			db, err := NewTestDBInstance(ctx)
			if err != nil {
				t.Errorf("Error creating test database: %v", err)
			}
			defer cancel()
			// setup initial state
			Migrate(ctx, db)

			// begin the test
			s := SynchronizeRepositoryImpl{
				db:  db,
				ctx: ctx,
			}

			session := (*db).NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
			defer session.Close(ctx)

			// Start a new transaction
			_, err = session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
				// Create the date node
				if err := s.ensureDateExists(tx, test.date.Format("2006-01-02"), TimeDateToWeekdayID(test.date)); err != nil {
					return nil, err
				}
				return nil, nil
			})
			if err != nil {
				t.Errorf("Error ensuring date exists: %v", err)
			}

			// check if the date was created
			results, err := neo4j.ExecuteQuery(
				ctx,
				*db,
				"MATCH (d:Date {date: date($date)}) RETURN d",
				map[string]interface{}{
					"date": test.date.Format("2006-01-02"),
				},
				neo4j.EagerResultTransformer,
			)
			if err != nil {
				t.Errorf("Error querying date: %v", err)
			}
			if len(results.Records) != 1 {
				t.Errorf("Expected 1 date, got %d", len(results.Records))
			}

			// compare weekdayID
			results, err = neo4j.ExecuteQuery(
				ctx,
				*db,
				"MATCH (d:Date {date: date($date)}) -[:IS_ON_WEEKDAY]-> (w:Weekday) RETURN d, w",
				map[string]interface{}{
					"date": test.date.Format("2006-01-02"),
				},
				neo4j.EagerResultTransformer,
			)
			if err != nil {
				t.Errorf("Error querying date: %v", err)
			}
			if len(results.Records) != 1 {
				t.Errorf("Expected 1 date, got %d", len(results.Records))
			}

			// compare weekdayID
			weekdayID, isNil, err := neo4j.GetRecordValue[neo4j.Node](results.Records[0], "w")
			if err != nil {
				t.Errorf("Error getting weekdayID: %v", err)
			}

			if isNil {
				t.Errorf("Expected weekdayID, got nil")
			}

			if weekdayID.Props["id"] != test.weekdayID {
				t.Errorf("Expected weekdayID %d, got %d", test.weekdayID, weekdayID.Props["id"])
			}
		})
	}
}

func TestCreateWorkDay(t *testing.T) {
	tests := []struct {
		name                    string
		create                  bool
		startDate               time.Time
		endDate                 time.Time
		expectedWorkdaysPerDate map[string]int
	}{
		{
			name:      "Sync workdays for 1 week",
			create:    true,
			startDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			endDate:   time.Date(2021, 1, 7, 0, 0, 0, 0, time.UTC),
			expectedWorkdaysPerDate: map[string]int{
				"2021-01-01": 1,
				"2021-01-02": 1,
				"2021-01-03": 1,
				"2021-01-04": 1,
				"2021-01-05": 1,
				"2021-01-06": 1,
				"2021-01-07": 1,
			},
		},
		{
			name:      "Sync workdays for 1 day",
			create:    true,
			startDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			endDate:   time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			expectedWorkdaysPerDate: map[string]int{
				"2021-01-01": 1,
			},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d: %s", i, test.name), func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			db, err := NewTestDBInstance(ctx)
			if err != nil {
				t.Errorf("Error creating test database: %v", err)
			}
			defer cancel()
			// setup initial state
			Migrate(ctx, db)

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
				},
			}
			if err := timeslotCreatorOne.Create(db, ctx); err != nil {
				t.Errorf("Error creating timeslots: %v", err)
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
					{id: 4, startTime: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC), endTime: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)},
					{id: 5, startTime: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC), endTime: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)},
					{id: 6, startTime: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC), endTime: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)},
				},
			}
			if err := timeslotCreatorTwo.Create(db, ctx); err != nil {
				t.Errorf("Error creating timeslots: %v", err)
			}
			timeslotCreatorThree := TimeslotCreatorImpl{
				departmentID:   "dept2",
				departmentName: "Department 2",
				workplaceID:    "wp2",
				workplaceName:  "Workplace 2",
				id:             "ts3",
				name:           "Timeslot 3",

				weekdays: []struct {
					id        int64
					startTime time.Time
					endTime   time.Time
				}{
					{id: 7, startTime: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC), endTime: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)},
				},
			}
			if err := timeslotCreatorThree.Create(db, ctx); err != nil {
				t.Errorf("Error creating timeslots: %v", err)
			}

			// begin the test
			s := SynchronizeRepositoryImpl{
				db:  db,
				ctx: ctx,
			}

			if test.create {
				session := (*db).NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
				defer session.Close(ctx)

				// Start a new transaction
				_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {

					// Create Workday nodes for each date and weekday
					for date := test.startDate; date.Before(test.endDate.AddDate(0, 0, 1)); date = date.AddDate(0, 0, 1) {
						// Create the date node
						if err := s.ensureDateExists(tx, date.Format("2006-01-02"), TimeDateToWeekdayID(date)); err != nil {
							return nil, err
						}
						// Create the workday node
						if err := s.createWorkday(tx, date.Format("2006-01-02"), TimeDateToWeekdayID(date)); err != nil {
							return nil, err
						}
					}
					return nil, nil
				})
				if err != nil {
					t.Errorf("Error creating workdays: %v", err)
				}
			}
			// check if the workdays were created
			results, err := neo4j.ExecuteQuery(
				ctx,
				*db,
				"MATCH (w:Workday) RETURN w",
				nil,
				neo4j.EagerResultTransformer,
			)
			if err != nil {
				t.Errorf("Error querying workdays: %v", err)
			}

			if len(results.Records) != len(test.expectedWorkdaysPerDate) {
				t.Errorf("Expected %d workdays, got %d", len(test.expectedWorkdaysPerDate), len(results.Records))
			}

			// Check if the workdays were created
			for date, expectedCount := range test.expectedWorkdaysPerDate {
				results, err := neo4j.ExecuteQuery(
					ctx,
					*db,
					"MATCH (w:Workday {date: date($date)}) RETURN w",
					map[string]interface{}{
						"date": date,
					},
					neo4j.EagerResultTransformer,
				)
				if err != nil {
					t.Errorf("Error querying workdays: %v", err)
				}

				if len(results.Records) != expectedCount {
					t.Errorf("Expected %d workdays, got %d", expectedCount, len(results.Records))
				}
			}
		})
	}
}

func TestEnsureSynchronizationRunsOncePerDateAndDepartment(t *testing.T) {
	/*
	* This test is supposed to ensure that the synchronization runs only once per date and department.
	 */
	tests := []struct {
		name                    string
		startDate               time.Time
		endDate                 time.Time
		expectedWorkdaysPerDate map[string]int
	}{
		{
			name:      "Sync workdays for 1 week",
			startDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			endDate:   time.Date(2021, 1, 7, 0, 0, 0, 0, time.UTC),
			expectedWorkdaysPerDate: map[string]int{
				"2021-01-01": 1,
				"2021-01-02": 1,
				"2021-01-03": 1,
				"2021-01-04": 1,
				"2021-01-05": 1,
				"2021-01-06": 1,
				"2021-01-07": 1,
			},
		},
		{
			name:      "Sync workdays for 1 day",
			startDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			endDate:   time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			expectedWorkdaysPerDate: map[string]int{
				"2021-01-01": 1,
			},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d: %s", i, test.name), func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			db, err := NewTestDBInstance(ctx)
			if err != nil {
				t.Errorf("Error creating test database: %v", err)
			}
			defer cancel()
			// setup initial state
			Migrate(ctx, db)

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
				},
			}
			if err := timeslotCreatorOne.Create(db, ctx); err != nil {
				t.Errorf("Error creating timeslots: %v", err)
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
					{id: 4, startTime: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC), endTime: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)},
					{id: 5, startTime: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC), endTime: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)},
					{id: 6, startTime: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC), endTime: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)},
				},
			}
			if err := timeslotCreatorTwo.Create(db, ctx); err != nil {
				t.Errorf("Error creating timeslots: %v", err)
			}
			timeslotCreatorThree := TimeslotCreatorImpl{
				departmentID:   "dept2",
				departmentName: "Department 2",
				workplaceID:    "wp2",
				workplaceName:  "Workplace 2",
				id:             "ts3",
				name:           "Timeslot 3",

				weekdays: []struct {
					id        int64
					startTime time.Time
					endTime   time.Time
				}{
					{id: 7, startTime: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC), endTime: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)},
				},
			}
			if err := timeslotCreatorThree.Create(db, ctx); err != nil {
				t.Errorf("Error creating timeslots: %v", err)
			}

			// begin the test
			s := SynchronizeRepositoryImpl{
				db:  db,
				ctx: ctx,
			}

			session := (*db).NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
			defer session.Close(ctx)

			// Start a new transaction
			if _, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {

				// Create Workday nodes for each date and weekday
				for date := test.startDate; date.Before(test.endDate.AddDate(0, 0, 1)); date = date.AddDate(0, 0, 1) {
					// Create the date node as we would normally do
					if err := s.ensureDateExists(tx, date.Format("2006-01-02"), TimeDateToWeekdayID(date)); err != nil {
						return nil, err
					}
					// Create the workday node as we would normally do
					if err := s.createWorkday(tx, date.Format("2006-01-02"), TimeDateToWeekdayID(date)); err != nil {
						return nil, err
					}

					// Run again to ensure that the synchronization runs only once
					if err := s.createWorkday(tx, date.Format("2006-01-02"), TimeDateToWeekdayID(date)); err == nil {
						// here we expect an error since no workday should be created
						t.Errorf("Expected error, got nil")
						return nil, errors.New("Expected error, got nil")
					}
				}
				return nil, nil
			}); err != nil {
				t.Errorf("Error creating workdays: %v", err)
			}

			// check if the workdays were created
			results, err := neo4j.ExecuteQuery(
				ctx,
				*db,
				"MATCH (w:Workday) RETURN w",
				nil,
				neo4j.EagerResultTransformer,
			)
			if err != nil {
				t.Errorf("Error querying workdays: %v", err)
			}

			if len(results.Records) != len(test.expectedWorkdaysPerDate) {
				t.Errorf("Expected %d workdays, got %d", len(test.expectedWorkdaysPerDate), len(results.Records))
			}

			// check if SYNCHRONIZED_AT property was set
			results, err = neo4j.ExecuteQuery(
				ctx,
				*db,
				"MATCH (d:Department) -[s:SYNCHRONIZED_AT]-> (d2:Date) RETURN d, s, d2",
				nil,
				neo4j.EagerResultTransformer,
			)
			if err != nil {
				t.Errorf("Error querying synchronized at: %v", err)
			}

			// check if s is present
			if len(results.Records) != len(test.expectedWorkdaysPerDate) {
				t.Errorf("Expected %d synchronized at, got %d", len(test.expectedWorkdaysPerDate), len(results.Records))
			}
		})
	}
}
