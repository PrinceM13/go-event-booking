package middlewares

import (
	"net/http"

	"github.com/PrinceM13/go-event-booking/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized access."})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized access.", "error": err.Error()})
		return
	}

	c.Set("userId", userId)
	c.Next()
}
