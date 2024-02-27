package repository

import (
	"context"
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
					id        string
					startTime time.Time
					endTime   time.Time
				}{
					{id: "MON", startTime: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC), endTime: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)},
					{id: "TUE", startTime: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC), endTime: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)},
					{id: "WED", startTime: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC), endTime: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)},
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
					id        string
					startTime time.Time
					endTime   time.Time
				}{
					{id: "THU", startTime: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC), endTime: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)},
					{id: "FRI", startTime: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC), endTime: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)},
					{id: "SAT", startTime: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC), endTime: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)},
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
					id        string
					startTime time.Time
					endTime   time.Time
				}{
					{id: "SUN", startTime: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC), endTime: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)},
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
					id        string
					startTime time.Time
					endTime   time.Time
				}{
					{id: "MON", startTime: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC), endTime: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)},
					{id: "TUE", startTime: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC), endTime: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)},
					{id: "WED", startTime: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC), endTime: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)},
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
					id        string
					startTime time.Time
					endTime   time.Time
				}{
					{id: "THU", startTime: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC), endTime: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)},
					{id: "FRI", startTime: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC), endTime: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)},
					{id: "SAT", startTime: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC), endTime: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)},
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
					id        string
					startTime time.Time
					endTime   time.Time
				}{
					{id: "SUN", startTime: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC), endTime: time.Date(2021, 1, 1, 16, 0, 0, 0, time.UTC)},
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
				for date := test.startDate; date.Before(test.endDate.AddDate(0, 0, 1)); date = date.AddDate(0, 0, 1) {
					EnsureDateExists(db, ctx, date.Format("2006-01-02"))
				}

				session := (*db).NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
				defer session.Close(ctx)

				_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
					for date := test.startDate; date.Before(test.endDate.AddDate(0, 0, 1)); date = date.AddDate(0, 0, 1) {
						if err := s.createWorkday(ctx, tx, date.Format("2006-01-02"), TimeDateToWeekdayID(date)); err != nil {
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
