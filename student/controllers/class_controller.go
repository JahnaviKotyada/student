package controllers

import (
	"net/http"
	"strconv"
	"stu/models"
	"stu/services"

	"github.com/gin-gonic/gin"
)

type ClassController struct {
	service services.ClassService
}

func NewClassController(service services.ClassService) *ClassController {
	return &ClassController{
		service: service,
	}
}

func (cc *ClassController) GetAllClasses(c *gin.Context) {
	classes, err := cc.service.GetAllClasses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, classes)
}

func (cc *ClassController) GetClassByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid class ID"})
		return
	}
	class, err := cc.service.GetClassByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Class not found"})
		return
	}
	c.JSON(http.StatusOK, class)
}

func (cc *ClassController) CreateClass(c *gin.Context) {
	var newClass models.Class
	if err := c.ShouldBindJSON(&newClass); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	err := cc.service.CreateClass(&newClass)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newClass)
}

func (cc *ClassController) UpdateClass(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid class ID"})
		return
	}
	var updatedClass models.Class
	if err := c.ShouldBindJSON(&updatedClass); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	updatedClass.ID = uint(id)
	err = cc.service.UpdateClass(&updatedClass)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedClass)
}

func (cc *ClassController) DeleteClass(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid class ID"})
		return
	}
	err = cc.service.DeleteClass(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Class deleted successfully"})
}
