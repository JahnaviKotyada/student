package routes

import (
	"stu/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(
	schoolController *controllers.SchoolController,
	classController *controllers.ClassController,
	studentController *controllers.StudentController,
) *gin.Engine {
	r := gin.Default()

	// Schools endpoints
	schools := r.Group("/schools")
	{
		schools.GET("/", schoolController.GetAllSchools)
		schools.GET("/:id", schoolController.GetSchoolByID)
		schools.POST("/", schoolController.CreateSchool)
		schools.PUT("/:id", schoolController.UpdateSchool)
		schools.DELETE("/:id", schoolController.DeleteSchool)
	}

	// Classes endpoints
	classes := r.Group("/classes")
	{
		classes.GET("/", classController.GetAllClasses)
		classes.GET("/:id", classController.GetClassByID)
		classes.POST("/", classController.CreateClass)
		classes.PUT("/:id", classController.UpdateClass)
		classes.DELETE("/:id", classController.DeleteClass)
	}

	// Students endpoints
	students := r.Group("/students")
	{
		students.GET("/", studentController.GetAllStudents)
		students.GET("/:id", studentController.GetStudentByID)
		students.POST("/", studentController.CreateStudent)
		students.PUT("/:id", studentController.UpdateStudent)
		students.DELETE("/:id", studentController.DeleteStudent)
	}

	return r
}
