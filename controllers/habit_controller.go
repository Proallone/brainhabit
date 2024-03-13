package controllers

import (
	pg "brainhabit/db"
	"brainhabit/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetHabits(c *gin.Context) {
	var habits []models.Habit
	pg.DB.Find(&habits)

	c.JSON(http.StatusOK, habits)
}

func CreateHabit(c *gin.Context) {

}

func PatchHabit(c *gin.Context) {

}

func DeleteHabit(c *gin.Context) {
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

	c.JSON(http.StatusNoContent, gin.H{"res": "Habit deleted"})
}

func GetHabit(c *gin.Context) {
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
