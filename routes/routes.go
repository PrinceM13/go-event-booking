package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// Health check route
	server.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	// Event routes
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/events", createEvent)
	server.PUT("/events/:id", updateEvent)
	server.DELETE("/events/:id", deleteEvent)

	// User routes
	server.POST("/signup", signUp)
}
