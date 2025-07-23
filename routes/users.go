package routes

import (
	"net/http"

	"github.com/PrinceM13/go-event-booking/models"
	"github.com/gin-gonic/gin"
)

type userResponse struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
}

func signUp(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
