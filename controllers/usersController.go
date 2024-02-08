package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jicodes/go-jwt/initializers"
	"github.com/jicodes/go-jwt/models"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	// Get the email and password from the request body
	var body struct {
		Email string `json:"email"`
		Password string `json:"password"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{ //400
			"error": "Failed to read request body",
		}) 
		return
	}
	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ //400
			"error": "Failed to hash the password",
		})
		return
	}
	// Create a new user and save it to the database
	user := models.User{Email: body.Email, Password: string(hash)}
	
	result := initializers.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{ //400
			"error": "Failed to create user",
		})
		return
	}
	// Return OK
	c.JSON(http.StatusOK, gin.H{}) //200
}






