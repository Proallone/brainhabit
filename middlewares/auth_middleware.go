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

		userClaims, err := isAuthenticated((authSplit[1]))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access!"})
			return
		}

		fmt.Println(userClaims)
		requestedPath := c.FullPath()

		if strings.Contains(requestedPath, "/admin/") && userClaims["role"] != "admin" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Insufficient priviliges!"})
			return
		}
		c.Next()
	}
}

func isAuthenticated(token string) (map[string]interface{}, error) {
	claims, err := utils.ValidateToken(token, os.Getenv("JWT_KEY"))
	if err != nil {
		return nil, err
	}
	return claims, nil
}
