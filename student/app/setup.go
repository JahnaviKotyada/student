package app

import (
	"stu/config"
	"stu/controllers"
	"stu/repository"
	"stu/services"
)

// SetupApp initializes the database and controllers
func SetupApp() (*controllers.SchoolController, *controllers.ClassController, *controllers.StudentController) {
	db := config.SetupDatabase()

	// Initialize repositories
	schoolRepo := repository.NewSchoolRepository(db)
	classRepo := repository.NewClassRepository(db)
	studentRepo := repository.NewStudentRepository(db)

	// Initialize services
	schoolService := services.NewSchoolService(schoolRepo)
	classService := services.NewClassService(classRepo)
	studentService := services.NewStudentService(studentRepo)

	// Initialize controllers
	schoolController := controllers.NewSchoolController(schoolService)
	classController := controllers.NewClassController(classService)
	studentController := controllers.NewStudentController(studentService)

	return schoolController, classController, studentController
}
