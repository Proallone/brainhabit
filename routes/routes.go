package routes

import (
	c "brainhabit/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	r := router.Group("/api")

	r.GET("/health", c.CheckServer)
	r.GET("/pghealth", c.CheckPostgres)

}
