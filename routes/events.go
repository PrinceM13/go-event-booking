package routes

import (
	"net/http"
	"strconv"

	"github.com/PrinceM13/go-event-booking/models"
	"github.com/gin-gonic/gin"
)

func getEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch events, please try again later."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"events": events})
}

func getEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch event, please try again later."})
		return
	}
	if event == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"event": event})
}

func createEvent(c *gin.Context) {
	var event models.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "error": err.Error()})
		return
	}

	userId := c.GetInt64("userId")
	event.UserID = userId

	if err := event.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create event, please try again later."})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Event created successfully", "event": event})
}

func updateEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	userId := c.GetInt64("userId")
	event, err := models.GetEventByID(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch event, please try again later."})
		return
	}
	if event == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
		return
	}

	if event.UserID != userId {
		c.JSON(http.StatusForbidden, gin.H{"message": "You do not have permission to update this event."})
		return
	}

	var updatedEvent models.Event
	if err := c.ShouldBindJSON(&updatedEvent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "error": err.Error()})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update event, please try again later."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event updated successfully"})
}

func deleteEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	userId := c.GetInt64("userId")
	event, err := models.GetEventByID(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch event, please try again later."})
		return
	}
	if event == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
		return
	}

	if event.UserID != userId {
		c.JSON(http.StatusForbidden, gin.H{"message": "You do not have permission to delete this event."})
		return
	}

	err = event.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete event, please try again later."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}
