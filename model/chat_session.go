package models

import (
	"time"
)

type ChatSession struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	UserID    int64     `json:"user_id"`
	Title     *string   `json:"title" gorm:"default:null"`
	CreatedAt time.Time `json:"created_at" gorm:"CURRENT_TIMESTAMP"`
	User      User      `json:"-"`
}
