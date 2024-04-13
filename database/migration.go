package database

import (
	models "chatbot-backend/model"
	"fmt"
	"log"
)

func Migration() {
	err := DB.AutoMigrate(
		&models.User{},
	)
	if err != nil {
		log.Fatal("Failed to migrate...")
	}

	fmt.Println("Migrated successfully")
}
