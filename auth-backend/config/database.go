package config

import (
	"auth-backend/app/domain/dao"
	"log/slog"
	"os"

	"golang.org/x/crypto/bcrypt"
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
		slog.Error("Failed to connect to database", "error", err)
		panic(err)
	}

	// Ping to database
	sqlDB, err := db.DB()
	if err != nil {
		slog.Error("Failed to connect to database", "error", err)
		panic(err)

	}
	err = sqlDB.Ping()
	if err != nil {
		slog.Error("Failed to connect to database", "error", err)
		panic(err)
	}

	return db
}

func checkIfValueExists(db *gorm.DB, model interface{}, field string, value interface{}) bool {
	/**
	 * This function is used to check if a value exists in a field
	 */

	// Check if value exists
	var count int64
	result := db.Model(model).Where(field+" = ?", value).Count(&count)
	return result.Error == nil && count > 0
}

func Migrate(db *gorm.DB) {
	/**
	 * Aside from the migrations of the models, we also need to migrate the user and initial database configuration
	 */

	// setup department
	department := dao.Department{
		Name: "IT",
	}
	if !checkIfValueExists(db, department, "name", department.Name) {
		db.Create(&department)
	} else {
		db.Where("name = ?", department.Name).First(&department)
	}

	// create user
	username := os.Getenv("AUTH_ADMIN_USERNAME")
	if username == "" {
		username = "admin"
	}
	password := os.Getenv("AUTH_ADMIN_PASSWORD")
	if password == "" {
		password = "admin"
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		slog.Error("Failed to hash password", "error", err)
		panic(err)
	}

	email := os.Getenv("AUTH_ADMIN_EMAIL")
	if email == "" {
		email = "admin@example.com"
	}

	user := dao.User{
		Username:   username,
		Password:   string(hash),
		Email:      email,
		Department: department,
		IsAdmin:    true,
	}

	// create instances
	if !checkIfValueExists(db, user, "username", user.Username) {
		db.Create(&user)
		slog.Info("Created admin user", "username", user.Username)
	} else {
		slog.Info("Admin user already exists", "username", user.Username)
	}
}
