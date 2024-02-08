package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jicodes/go-jwt/initializers"
)

func init () {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
	initializers.SyncDB()
}

func main() {
	r := gin.Default()
	r.GET("/login", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})
	r.Run()

}