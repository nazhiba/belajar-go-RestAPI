package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Membuat router default
	r := gin.Default()

	// Menambahkan endpoint GET
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Menjalankan server di port 8080
	r.Run(":8080")
}
