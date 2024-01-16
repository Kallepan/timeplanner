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
				Values: []interface{}{
					[]interface{}{"workplaces", "workplace2"},
					[]interface{}{"departments", "department2"},
					[]interface{}{"monday", "tuesday"},
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
					[]interface{}{},
					[]interface{}{},
					[]interface{}{},
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
