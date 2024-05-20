package controllers

import (
	pg "brainhabit/data"
	"brainhabit/data/models"
	"brainhabit/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (controller *Controller) LoginUser(c *gin.Context) {
	type Credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var creds Credentials

	if err := c.BindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	var user models.User

	if err := pg.DB.Select("id,password_hash, role").Where("email = ?", creds.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, "User not found")
		return
	}

	if !utils.VerifyPassword(creds.Password, user.PasswordHash) {
		c.JSON(http.StatusUnauthorized, "Incorrect password")
		return
	}

	jwtTTL := time.Minute * 30

	token, err := utils.GenerateToken(jwtTTL, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (controller *Controller) LogoutUser(c *gin.Context) {
	//TODO remove placeholder after there is proper JWT implemented
	c.JSON(http.StatusOK, "Successfully logged out")
}
