package controllers

import (
	"net/http"
	"student-crud/config"
	"student-crud/models"

	"github.com/gin-gonic/gin"
)

func CreateStudent(c *gin.Context) {
    var student models.Student
    if err := c.ShouldBindJSON(&student); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    config.DB.Create(&student)
    c.JSON(http.StatusCreated, student)
}

func GetStudents(c *gin.Context) {
    var students []models.Student
    config.DB.Find(&students)
    c.JSON(http.StatusOK, students)
}

func GetStudent(c *gin.Context) {
    var student models.Student
    if err := config.DB.First(&student, "id = ?", c.Param("id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
        return
    }
    c.JSON(http.StatusOK, student)
}

func UpdateStudent(c *gin.Context) {
    var student models.Student
    if err := config.DB.First(&student, "id = ?", c.Param("id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
        return
    }
    if err := c.ShouldBindJSON(&student); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    config.DB.Save(&student)
    c.JSON(http.StatusOK, student)
}

func DeleteStudent(c *gin.Context) {
    var student models.Student
    if err := config.DB.First(&student, "id = ?", c.Param("id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
        return
    }
    config.DB.Delete(&student)
    c.JSON(http.StatusNoContent, gin.H{"message": "Student deleted"})
}