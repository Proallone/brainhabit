package main

import (
	"fmt"
	"net/http"
	"time"

	postgres "brainhabit/db"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {

	dsn := "user=postgres password=postgres dbname=postgres host=localhost port=5432 sslmode=disable"

	var database *gorm.DB
	database, _ = postgres.Connect(dsn)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/time", func(c *gin.Context) {

		var now time.Time
		result := database.Raw("SELECT NOW() AS now").Scan(&now)

		if result.Error != nil {
			fmt.Println("Error fetching current time:", result.Error)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"pg_time": now,
		})
	})

	r.Run()
}

func pgTest() {

}
