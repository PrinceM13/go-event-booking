package main

import (
	"net/http"

	"github.com/PrinceM13/go-event-booking/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080")
}

func getEvents(c *gin.Context) {
	events := models.GetAllEvents()
	c.JSON(http.StatusOK, gin.H{"events": events})
}

func createEvent(c *gin.Context) {
	var event models.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Mock saving the event
	event.ID = 1
	event.UserID = 1

	event.Save()

	c.JSON(http.StatusCreated, gin.H{"message": "Event created successfully", "event": event})
}
