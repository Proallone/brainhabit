package controllers

import (
	pg "brainhabit/data"
	"brainhabit/data/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (controller *Controller) CreateHabit(c *gin.Context) {
	var habit models.Habit
	if err := c.ShouldBindJSON(&habit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := pg.DB.Create(&habit).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create habit"})
		return
	}

	c.JSON(http.StatusCreated, habit)
}

func (controller *Controller) PatchHabit(c *gin.Context) {
	ID, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid habit ID"})
		return
	}

	var habit models.Habit

	if err := pg.DB.Where("id = ?", ID).First(&habit).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Habit not found!"})
		return
	}

	if err := c.BindJSON(&habit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pg.DB.Model(&habit).Updates(habit)

	c.JSON(http.StatusOK, habit)
}

func (controller *Controller) DeleteHabit(c *gin.Context) {
	ID, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var habit models.Habit

	err = pg.DB.Delete(&habit, ID).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cannot delete habit"})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"response": "Habit deleted"})
}

func (controller *Controller) GetHabit(c *gin.Context) {
	ID, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid habit ID"})
		return
	}

	var habit models.Habit

	if err := pg.DB.First(&habit, ID).Where("deleted_at is null", ID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Habit not found"})
		return
	}

	c.JSON(http.StatusOK, habit)
}
