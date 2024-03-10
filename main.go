package main

import (
	"net/http"

	pg "brainhabit/db"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {

	dsn := "user=postgres password=postgres dbname=postgres host=localhost port=5432 sslmode=disable"

	var database *gorm.DB
	database, _ = pg.Connect(dsn)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/pg", func(c *gin.Context) {

		currentDB := database.Migrator().CurrentDatabase()

		c.JSON(http.StatusOK, gin.H{"message": "Database connection OK", "database": currentDB})
	})

	r.Run()
}
