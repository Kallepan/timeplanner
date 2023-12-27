package config

import (
	"log/slog"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB() *gorm.DB {
	dsn := os.Getenv("AUTH_DB_DSN")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// Do not use transaction for each writes
		SkipDefaultTransaction: true,

		// Prepare statement before executing and cache them
		PrepareStmt: true,

	})
	if err != nil {
		slog.Error("Failed to connect to database")
		panic(err)
	}

	return db
}