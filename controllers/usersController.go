package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
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

func Login(c *gin.Context) {
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

	// look up requeted user in the database
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{ //401
			"error": "Invalid credentials",
		})
		return
	}
	// compare sent in password with the hashed password in the database
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{ //401
			"error": "Invalid credentials",
	})
		return
	}
	// if the password is correct, generate a JWT 
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID, // subject
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(), // expiration time
	
	})
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ //400
			"error": "Failed to generate token",
		})
		return
	}
	
	// set the token in a cookie
	c.SetSameSite(http.SameSiteLaxMode) // set the SameSite attribute to Lax, which means the cookie will be sent with same-site requests, and with cross-site top-level navigations, to protect against cross-site request forgery attacks
	c.SetCookie("Authorization", tokenString, 60*60*24*30, "", "", false, true)

	// send it back to the client in the response 
	c.JSON(http.StatusOK, gin.H{"token": tokenString}) //200
}

func Validate(c *gin.Context) {
	// Get the user from the context
	user, _ := c.Get("user")

	// Get the user's email
	// user.(models.User).Email

	c.JSON(http.StatusOK, gin.H{  //200
		"logged in user" : user,
	})
}
