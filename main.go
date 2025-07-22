package main

import (
	"net/http"

	"github.com/PrinceM13/go-event-booking/db"
	"github.com/PrinceM13/go-event-booking/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
