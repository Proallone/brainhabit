package main

import (
	pg "brainhabit/data"
	"brainhabit/routes"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {

	r := setupRouter()
	routes.Routes(r)
	dns := getDatabaseConnString()
	fmt.Print(dns)
	pg.Setup(dns)
	port := getPort()

	if err := r.Run(port); err != nil {
		panic(err)
	}
}

func getPort() string {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	return ":" + port
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodDelete},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	return router
}

func getDatabaseConnString() string {
	return os.Getenv("CONNECTION_STRING")
}
