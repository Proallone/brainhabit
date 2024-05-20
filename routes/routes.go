package routes

import (
	"brainhabit/controllers"
	"brainhabit/controllers/middlewares"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	r := router.Group("/api")
	c := controllers.NewController()

	health := r.Group("/health")
	{
		health.GET("/", c.CheckServer)
		health.GET("/pg", c.CheckPostgres)
	}

	auth := r.Group("/auth")
	{
		auth.POST("/login", c.LoginUser)
		auth.GET("/logout", c.LogoutUser)
		auth.POST("/register", c.RegisterUser)
	}

	user := r.Group("/users", middlewares.AuthMiddleware())
	{
		user.GET("/:id", c.GetUser)
		user.PATCH("/:id", c.PatchUser)
		user.DELETE("/:id", c.DeleteUser)
	}

	admin := r.Group("/admin", middlewares.AuthMiddleware())
	{
		admin.GET("/users", c.GetUsers)
		admin.GET("/habits", c.GetHabits)
		admin.GET("/records", c.GetHabitRecords)
	}

	habit := r.Group("/habits", middlewares.AuthMiddleware())
	{
		habit.POST("/", c.CreateHabit)
		habit.GET("/:id", c.GetHabit)
		habit.PATCH("/:id", c.PatchHabit)
		habit.DELETE("/:id", c.DeleteHabit)
	}

	habit_record := r.Group("/records", middlewares.AuthMiddleware())
	{
		habit_record.POST("/", c.CreateHabitRecords)
		habit_record.PATCH("/:id", c.PatchHabitRecord)
		habit_record.DELETE("/:id", c.DeleteHabitRecord)
	}

}
