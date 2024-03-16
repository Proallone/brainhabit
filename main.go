package main

import (
	pg "brainhabit/db"
	"brainhabit/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("No .env file found")
	}

	r := gin.Default()
	routes.Routes(r)

	pg.Setup("user=postgres password=postgres dbname=postgres host=localhost port=5432 sslmode=disable")

	r.Run()
}
