package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jicodes/go-jwt/initializers"
	"github.com/jicodes/go-jwt/controllers"
)

func init () {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
	initializers.SyncDB()
}

func main() {
	r := gin.Default()

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	
	r.Run()

}