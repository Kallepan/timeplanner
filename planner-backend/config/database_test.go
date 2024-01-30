package config

import (
	"context"
	"planner-backend/app/mock"
	"testing"
)

func TestMigrate(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db, err := mock.SetupTestDB(ctx, t)
	if err != nil {
		t.Errorf("Failed to connect to mock database")
	}

	Migrate(ctx, db)
	Clear(ctx, db)
}
