package main

import (
	"log"
	"student-crud/config"
	"student-crud/controllers"
	"student-crud/models"

	"github.com/gin-gonic/gin"
)

func main() {
    // Initialize database connection
    config.ConnectDB()

    // Auto migrate the Student model
    config.DB.AutoMigrate(&models.Student{})

    // Set up Gin router
    router := gin.Default()
	
	router.POST("/students", controllers.CreateStudent)
    router.GET("/students", controllers.GetStudents)
    router.GET("/students/:id", controllers.GetStudent)
    router.PUT("/students/:id", controllers.UpdateStudent)
    router.DELETE("/students/:id", controllers.DeleteStudent)

    log.Println("Server running on port 8080")
    router.Run(":8080")
}
