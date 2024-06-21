package main

import (
	"log"
	"stu/app"
	"stu/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	schoolController, classController, studentController := app.SetupApp()

	router := gin.Default()
	router = routes.SetupRouter(schoolController, classController, studentController)
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
