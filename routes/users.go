package routes

import (
	"net/http"

	"github.com/PrinceM13/go-event-booking/models"
	"github.com/PrinceM13/go-event-booking/utils"
	"github.com/gin-gonic/gin"
)

type userResponse struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
}

func signUp(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "error": err.Error()})
		return
	}

	if err := user.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create user, please try again later."})
		return
	}

	response := userResponse{
		ID:    user.ID,
		Email: user.Email,
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": response})
}

func login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "error": err.Error()})
		return
	}

	err := user.ValidateCredentials()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Could authenticate user", "error": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could authenticate user", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}
