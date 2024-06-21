package config

import (
	"log"
	"stu/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabase() *gorm.DB {
	dsn := "host=localhost user=postgres password=jahnavi@2003 dbname=stu port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = db.AutoMigrate(&models.School{}, &models.Class{}, &models.Student{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	return db
}
