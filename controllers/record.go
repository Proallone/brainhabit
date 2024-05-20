package controllers

import (
	pg "brainhabit/data"
	"brainhabit/data/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (controller *Controller) CreateHabitRecords(c *gin.Context) {
	var record models.HabitRecord
	if err := c.ShouldBindJSON(&record); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := pg.DB.Create(&record).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create record"})
		return
	}

	c.JSON(http.StatusCreated, record)
}

func (controller *Controller) PatchHabitRecord(c *gin.Context) {
	ID, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid record ID"})
		return
	}

	var record models.HabitRecord

	if err := pg.DB.Where("id = ?", ID).First(&record).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	if err := c.BindJSON(&record); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pg.DB.Model(&record).Updates(record)

	c.JSON(http.StatusOK, record)
}

func (controller *Controller) DeleteHabitRecord(c *gin.Context) {
	ID, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var record models.HabitRecord

	err = pg.DB.Unscoped().Delete(&record, ID).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cannot delete record"})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"response": "Record deleted"})
}
