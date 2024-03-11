package controllers

import (
	"net/http"

	pg "brainhabit/db"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CheckServer(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"Server": "Healthy",
	})

}

func CheckPostgres(c *gin.Context) {

	dsn := "user=postgres password=postgres dbname=postgres host=localhost port=5432 sslmode=disable"

	var database *gorm.DB
	database, _ = pg.Connect(dsn)

	currentDB := database.Migrator().CurrentDatabase()

	c.JSON(http.StatusOK, gin.H{"message": "Database connection OK", "database": currentDB})

}
