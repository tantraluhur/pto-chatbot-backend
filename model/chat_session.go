package models

type ChatSession struct {
	ID     string `json:"id" gorm:"primaryKey"`
	UserID int64  `json:"user_id"`

	User User `json:"-"`
}
