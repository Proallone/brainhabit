package main

import (
	"brainhabit/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	routes.Routes(r)
	r.Run()
}
