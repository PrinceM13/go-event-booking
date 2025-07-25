package routes

import (
	"net/http"

	"github.com/PrinceM13/go-event-booking/middlewares"
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

	// User routes
	server.POST("/signup", signUp)
	server.POST("/login", login)

	// Protected routes that require authentication
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)

}
