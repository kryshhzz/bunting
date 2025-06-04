package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "room.html", nil)
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"currentTime": time.Now().UnixMilli(),
		})
	})

	// Start server on :8080
	fmt.Println("server running on https://localhost:8080/")
	r.Run() // Defaults to localhost:8080
}
