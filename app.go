package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load(".env." + gin.Mode())
}

func helloGet(c *gin.Context) {
	session := sessions.Default(c)

	var count int
	v := session.Get("count")

	if v == nil {
		count = 0
	} else {
		count = v.(int)
		count++
	}

	session.Set("count", count)
	session.Save()

	c.JSON(http.StatusOK, gin.H{"hello": session.Get("count")})
}

func main() {
	r := gin.Default()
	store, err := redis.NewStore(10, "tcp", os.Getenv("APP_REDIS_ADDR"), "", []byte("secret"))
	if err != nil {
		log.Fatal(err)
	}
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/hello", helloGet)

	r.Run(os.Getenv("APP_PORT"))
}
