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
	dns := getPgDns()
	fmt.Print(dns)
	pg.Setup(dns)
	port := getPort()
	r.Run(port)
}

func getPort() string {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	return ":" + port
}

func getPgDns() string {
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	db := os.Getenv("POSTGRES_DB")
	return fmt.Sprintf("user=%s password=%s dbname=%s host=postgres port=5432 sslmode=disable", user, password, db)

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
