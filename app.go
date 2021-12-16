package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello worldddd",
			"port":    os.Getenv("APP_PORT"),
		})
	})
	r.Run(os.Getenv("APP_PORT"))
}
