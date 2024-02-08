package initializers

import "github.com/jicodes/go-jwt/models"

func SyncDB () {
	DB.AutoMigrate(&models.User{})
}