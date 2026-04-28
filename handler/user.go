package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/harshit14100/go-todo/database/dbHelper"
	"github.com/harshit14100/go-todo/models"
	"github.com/harshit14100/go-todo/utils"
)

func Register(c *gin.Context) {

	var input models.CreateUser

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	isUserExist, err := dbHelper.IsUserExist(input.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	if isUserExist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exist"})
		return
	}

	hashedPass, _ := utils.HashPassword(input.Password)
	err = dbHelper.CreateUser(input.Email, input.Username, hashedPass)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func Login(c *gin.Context) {

	var input models.LoginUser

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email and password required"})
		return
	}

	user, err := dbHelper.GetUserByEmail(input.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	if !utils.CheckPassword(user.Password, input.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User logged in successfully",
		"token":   token,
	})
}

func Logout(c *gin.Context) {
	sessionID := c.GetHeader("Authorization")
	if sessionID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No active session"})
		return
	}
	err := dbHelper.DeleteUserSession(sessionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete user session",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User logged out"})
}
