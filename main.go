package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})
	server.Run(":8080")
}
