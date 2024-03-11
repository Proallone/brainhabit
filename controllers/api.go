package controllers

import (
	"net/http"

	pg "brainhabit/db"

	"github.com/gin-gonic/gin"
)

func CheckServer(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Server": "Healthy",
	})
}

func CheckPostgres(c *gin.Context) {

	dbName := pg.DB.Migrator().CurrentDatabase()

	c.JSON(http.StatusOK, gin.H{"message": "Database connection OK", "database": dbName})

}
