package controllers

import (
	"net/http"

	pg "brainhabit/data"

	"github.com/gin-gonic/gin"
)

func (controller *Controller) CheckServer(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Server": "Healthy",
	})
}

func (controller *Controller) CheckPostgres(c *gin.Context) {

	dbName := pg.DB.Migrator().CurrentDatabase()

	c.JSON(http.StatusOK, gin.H{"message": "Database connection OK", "database": dbName})

}
