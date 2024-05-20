package controllers

import (
	pg "brainhabit/data"
	"brainhabit/data/models"
	"brainhabit/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (controller *Controller) GetUsers(c *gin.Context) {
	var users []models.User
	pg.DB.Find(&users)

	c.JSON(http.StatusOK, users)
}

func (controller *Controller) RegisterUser(c *gin.Context) {
	type NewUser struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var newUser NewUser

	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	passHash, err := utils.HashPassword(newUser.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error during password hashing"})
		return
	}

	user := models.User{Email: newUser.Email, PasswordHash: passHash}
	result := pg.DB.Create(&user)

	if err := result.Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, result)
}

func (controller *Controller) PatchUser(c *gin.Context) {
	ID, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var user models.User

	if err := pg.DB.Where("id = ?", ID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found!"})
		return
	}

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pg.DB.Model(&user).Updates(user)

	c.JSON(http.StatusOK, user)
}

func (controller *Controller) DeleteUser(c *gin.Context) {
	ID, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var user models.User

	err = pg.DB.Delete(&user, ID).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cannot delete user"})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"response": "User deleted"})
}
