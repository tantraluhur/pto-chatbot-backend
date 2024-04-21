package models

type Message struct {
	ID            string `json:"id" gorm:"primaryKey"`
	Message       string `json:"message"`
	ChatSessionId string `json:"chat_session"`
	UserID        int64  `json:"user"`

	ChatSession ChatSession
	User        User
}
