package middlewares

import (
	"brainhabit/utils"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		authSplit := strings.Split(authHeader, "Bearer ")

		if len(authSplit) != 2 { //If other than 2 no correct split match
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Incorrect token"})
			return
		}

		if isAuthenticated(authSplit[1]) {
			c.Next()
			return
		}
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access!"})
	}
}

func isAuthenticated(token string) bool {
	claims, err := utils.ValidateToken(token, os.Getenv("JWT_KEY"))
	fmt.Print(claims)
	return err == nil
}
