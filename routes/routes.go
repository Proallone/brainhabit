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

	habit := r.Group("/habits")
	{
		habit.GET("/", c.GetHabits)
		habit.POST("/", c.CreateHabit)

		habit.GET("/:id", c.GetHabit)
		habit.PATCH("/:id", c.PatchHabit)
		habit.DELETE("/:id", c.DeleteHabit)
	}

	habit_record := r.Group("/records")
	{
		habit_record.GET("/", c.GetHabitRecords)
		habit_record.POST("/", c.CreateHabitRecords)

		habit_record.GET("/:id", c.GetHabitRecord)
		habit_record.PATCH("/:id", c.PatchHabitRecord)
		habit_record.DELETE("/:id", c.DeleteHabitRecord)
	}

}
