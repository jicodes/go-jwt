package initializers

import (
	"log"

	"github.com/jicodes/go-jwt/models"
)

func SyncDB () {
	DB.AutoMigrate(&models.User{})
	if err := DB.AutoMigrate(&models.User{}); err != nil {
			log.Fatal("Error migrating the database schema:", err)
	}
}