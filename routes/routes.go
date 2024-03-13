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

	user := r.Group("/users")
	{
		user.GET("/", c.GetUsers)
		user.POST("/", c.RegisterUser)

		user.GET("/:id", c.GetUser)
		user.PATCH("/:id", c.PatchUser)
		user.DELETE("/:id", c.DeleteUser)
	}

	habit := r.Group("/habit")
	{
		habit.GET("/", c.GetHabits)
		habit.POST("/", c.CreateHabit)

		habit.GET("/:id", c.GetHabit)
		habit.PATCH("/:id", c.PatchHabit)
		habit.DELETE("/:id", c.DeleteHabit)
	}

}
