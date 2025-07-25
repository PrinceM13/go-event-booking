package main

import (
	"github.com/PrinceM13/go-event-booking/db"
	"github.com/PrinceM13/go-event-booking/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
