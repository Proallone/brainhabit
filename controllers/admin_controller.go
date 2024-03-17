package controllers

import (
	pg "brainhabit/db"
	"brainhabit/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetUser(c *gin.Context) {
	ID, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var user models.User

	if err := pg.DB.Model(&models.User{}).Preload("Habits").Preload("Habits.Records").First(&user, ID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func GetHabits(c *gin.Context) {
	var habits []models.Habit
	pg.DB.Find(&habits)

	c.JSON(http.StatusOK, habits)
}

func GetHabitRecords(c *gin.Context) {
	var habits_records []models.HabitRecord
	pg.DB.Find(&habits_records)

	c.JSON(http.StatusOK, habits_records)
}
