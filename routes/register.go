package routes

import (
	"net/http"
	"strconv"

	"github.com/PrinceM13/go-event-booking/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(c *gin.Context) {
	userId := c.GetInt64("userId")
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	event, err := models.GetEventByID(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch event, please try again later."})
		return
	}
	if event == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
		return
	}

	err = event.Register(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to register for event, please try again later."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully registered for event"})
}

func cancelRegistration(c *gin.Context) {}
