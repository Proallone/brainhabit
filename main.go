package main

import (
	pg "brainhabit/db"
	"brainhabit/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	routes.Routes(r)

	pg.Setup("user=postgres password=postgres dbname=postgres host=localhost port=5432 sslmode=disable")

	r.Run()
}
