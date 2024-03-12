package routes

import (
	c "brainhabit/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	r := router.Group("/api")

	health := r.Group("/health")
	{
		health.GET("/", c.CheckServer)
		health.GET("/pg", c.CheckPostgres)
	}

	users := r.Group("/users")
	{
		users.GET("/", c.GetUsers)
		users.POST("/", c.RegisterUser)

		users.GET("/:id", c.GetUser)
		users.PATCH("/:id", c.PatchUser)
		users.DELETE("/:id", c.DeleteUser)

	}

}
