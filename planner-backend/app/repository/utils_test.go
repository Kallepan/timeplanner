package repository

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func TestTimeDateToWeekdayID(t *testing.T) {
	// Test the TimeDateToWeekdayID function

	tests := []struct {
		name string
		date string
		want int64
	}{
		{
			name: "TestTimeDateToWeekdayID",
			date: "2021-01-01",
			want: 5,
		},
		{
			name: "TestTimeDateToWeekdayID",
			date: "2021-01-02",
			want: 6,
		},
		{
			name: "TestTimeDateToWeekdayID",
			date: "2020-12-31",
			want: 4,
		},
		{
			name: "TestTimeDateToWeekdayID",
			date: "2024-02-29",
			want: 4,
		},
		{
			name: "TestTimeDateToWeekdayID",
			date: "2020-12-30",
			want: 3,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%s: %d", test.name, i), func(t *testing.T) {
			parsedDate, err := time.Parse("2006-01-02", test.date)
			if err != nil {
				t.Errorf("Error parsing date: %s", err)
			}

			got := TimeDateToWeekdayID(parsedDate)
			if got != test.want {
				t.Errorf("Got: %d, Want: %d", got, test.want)
			}
		})
	}
}

func TestEnsureDateExists(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	db, err := NewTestDBInstance(ctx)
	if err != nil {
		t.Errorf("Error creating test database: %v", err)
	}
	defer cancel()
	Migrate(ctx, db)

	/* Test the ensure date exists function */
	tests := []struct {
		name string
		date string
	}{
		{
			name: "TestEnsureDateExists",
			date: "2021-01-01",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%s: %d", test.name, i), func(t *testing.T) {
			err := EnsureDateExists(db, ctx, test.date)
			if err != nil {
				t.Errorf("Error ensuring date exists: %s", err)
			}

			// Ensure that the date was created
			res, err := neo4j.ExecuteQuery(
				context.Background(),
				*db,
				"MATCH (d:Date {date: date($date)}) RETURN d",
				map[string]interface{}{
					"date": test.date,
				},
				neo4j.EagerResultTransformer,
			)
			if err != nil {
				t.Errorf("Error checking if date was created: %s", err)
			}

			if len(res.Records) == 0 {
				t.Errorf("Date was not created")
			}
		})
	}
}

func TestEnsureWeekdaysExist(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	db, err := NewTestDBInstance(ctx)
	if err != nil {
		t.Errorf("Error creating test database: %v", err)
	}
	defer cancel()
	Migrate(ctx, db)

	tests := []struct {
		name string
		id   int64
	}{
		{
			name: "TestEnsureWeekdaysExist (MONDAY)",
			id:   1,
		},
		{
			name: "TestEnsureWeekdaysExist (TUESDAY)",
			id:   2,
		},
		{
			name: "TestEnsureWeekdaysExist (WEDNESDAY)",
			id:   3,
		},
		{
			name: "TestEnsureWeekdaysExist (THURSDAY)",
			id:   4,
		},
		{
			name: "TestEnsureWeekdaysExist (FRIDAY)",
			id:   5,
		},
		{
			name: "TestEnsureWeekdaysExist (SATURDAY)",
			id:   6,
		},
		{
			name: "TestEnsureWeekdaysExist (SUNDAY)",
			id:   7,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%s: %d", test.name, i), func(t *testing.T) {
			res, err := neo4j.ExecuteQuery(
				ctx,
				*db,
				"MATCH (d:Weekday {id: $id}) RETURN d",
				map[string]interface{}{
					"id": test.id,
				},
				neo4j.EagerResultTransformer,
			)
			if err != nil {
				t.Errorf("Error checking if weekday exists: %s", err)
			}

			if len(res.Records) == 0 {
				t.Errorf("Weekday %d was not created", test.id)
			}

			if len(res.Records) > 1 {
				t.Errorf("More than one weekday %d was created", test.id)
			}
		})
	}
}
