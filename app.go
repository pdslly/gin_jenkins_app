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
			"message": "hello world",
			"port":    os.Getenv("APP_REDIS_ADDR"),
		})
	})
	r.Run(os.Getenv("APP_PORT"))
}
