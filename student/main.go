package main

import (
	"log"
	"stu/controllers"
	"stu/models"
	"stu/repository"
	"stu/routes"
	"stu/services"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=jahnavi@2003 dbname=stu port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}

	db.AutoMigrate(&models.School{}, &models.Class{}, &models.Student{})

	repo := repository.NewRepository(db)
	schoolService := services.NewService(repo)
	classService := services.NewService(repo)
	studentService := services.NewService(repo)

	schoolController := controllers.NewSchoolController(schoolService)
	classController := controllers.NewClassController(classService)
	studentController := controllers.NewStudentController(studentService)

	router := routes.SetupRouter(schoolController, classController, studentController)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
