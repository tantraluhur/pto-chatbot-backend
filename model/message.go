package models

type Message struct {
	ID            string `json:"id" gorm:"primaryKey"`
	UserMessage   string `json:"user_message"`
	BotMessage    string `json:"bot_message"`
	ChatSessionId string `json:"chat_session_id"`
	UserID        int64  `json:"user_id"`

	ChatSession ChatSession `json:"-" gorm:"foreignKey:ChatSessionId"`
	User        User        `json:"-" gorm:"foreignKey:UserID"`
}
