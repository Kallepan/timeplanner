package dao

import (
	"testing"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func TestParseAdditionalFieldsFromDBRecord(t *testing.T) {
	type test struct {
		record  *neo4j.Record
		wantErr bool
	}

	tests := []test{
		{
			record: &neo4j.Record{
				Values: []any{
					[]any{
						neo4j.Node{
							Id:        1,
							ElementId: "test",
							Labels:    []string{"Workplace"},
							Props: map[string]interface{}{
								"id":            "test",
								"name":          "test",
								"department_id": "test",
							},
						},
					},
					[]any{
						neo4j.Node{
							Id:        1,
							ElementId: "test",
							Labels:    []string{"Department"},
							Props: map[string]interface{}{
								"id":   "test",
								"name": "test",
							},
						},
					},
					[]any{
						neo4j.Node{
							Id:        1,
							ElementId: "test",
							Labels:    []string{"Weekday"},
							Props: map[string]interface{}{
								"id":   *new(int64),
								"name": "Monday",
							},
						},
					},
				},
				Keys: []string{"workplaces", "departments", "weekdays"},
			},
			wantErr: false,
		},
		{
			record: &neo4j.Record{
				Values: []interface{}{
					nil,
					nil,
					nil,
				},
				Keys: []string{"workplaces", "departments", "weekdays"},
			},
			wantErr: false,
		},
		{
			record: &neo4j.Record{
				Values: []interface{}{
					[]interface{}{neo4j.Node{}},
					[]interface{}{neo4j.Node{}},
					[]interface{}{neo4j.Node{}},
				},
				Keys: []string{},
			},
			wantErr: true,
		},
	}

	for i, tc := range tests {
		t.Run("TestParseAdditionalFieldsFromDBRecord", func(t *testing.T) {
			p := Person{}
			if err := p.ParseAdditionalFieldsFromDBRecord(tc.record); (err != nil) != tc.wantErr {
				t.Errorf("Error: %d, ParseAdditionalFieldsFromDBRecord() error = %v, wantErr %v", i, err, tc.wantErr)
			}
		})
	}
}
