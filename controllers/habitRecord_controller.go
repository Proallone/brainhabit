package controllers

import (
	pg "brainhabit/db"
	"brainhabit/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHabitRecords(c *gin.Context) {
	var habits_records []models.HabitRecord
	pg.DB.Find(&habits_records)

	c.JSON(http.StatusOK, habits_records)
}

func CreateHabitRecords(c *gin.Context) {

}

func PatchHabitRecord(c *gin.Context) {

}

func DeleteHabitRecord(c *gin.Context) {

}

func GetHabitRecord(c *gin.Context) {

}
