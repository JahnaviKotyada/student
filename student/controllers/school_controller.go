package controllers

import (
	"net/http"
	"strconv"
	"stu/models"
	"stu/services"
	"testing"

	"github.com/gin-gonic/gin"
)

type SchoolController struct {
	service services.SchoolService
}

func NewSchoolController(service services.SchoolService) *SchoolController {
	return &SchoolController{
		service: service,
	}
}

func (sc *SchoolController) GetAllSchools(c *gin.Context) {
	schools, err := sc.service.GetAllSchools()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, schools)
}

func (sc *SchoolController) GetSchoolByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid school ID"})
		return
	}
	school, err := sc.service.GetSchoolByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "School not found"})
		return
	}
	c.JSON(http.StatusOK, school)
}

func (sc *SchoolController) CreateSchool(c *gin.Context) {
	var newSchool models.School
	if err := c.ShouldBindJSON(&newSchool); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	err := sc.service.CreateSchool(&newSchool)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newSchool)
}

func (sc *SchoolController) UpdateSchool(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid school ID"})
		return
	}
	var updatedSchool models.School
	if err := c.ShouldBindJSON(&updatedSchool); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	updatedSchool.ID = uint(id)
	err = sc.service.UpdateSchool(&updatedSchool)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedSchool)
}

func (sc *SchoolController) DeleteSchool(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid school ID"})
		return
	}
	err = sc.service.DeleteSchool(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "School deleted successfully"})
}

func RunTests(m *testing.M) {
	m.Run()
}
