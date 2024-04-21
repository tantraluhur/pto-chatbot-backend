package database

import (
	models "chatbot-backend/model"
	"fmt"
	"log"
)

func Migration() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.AccessToken{},
		&models.RefreshToken{},
		&models.Message{},
		&models.ChatSession{},
	)
	if err != nil {
		log.Fatal("Failed to migrate...")
	}

	fmt.Println("Migrated successfully")
}
