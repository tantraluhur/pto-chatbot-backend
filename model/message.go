package models

type Message struct {
	ID            string `json:"id" gorm:"primaryKey"`
	Message       string `json:"message"`
	ChatSessionId string `json:"chat_session_id"`
	UserID        int64  `json:"user_id"`

	ChatSession ChatSession `json:"-" gorm:"foreignKey:ChatSessionId"`
	User        User        `json:"-" gorm:"foreignKey:UserID"`
}
