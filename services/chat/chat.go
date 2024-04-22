package chat

import (
	"chatbot-backend/database"
	models "chatbot-backend/model"

	"github.com/google/uuid"
)

func GetAllChatSession(user *models.User) []*models.ChatSession {
	var chatHistoryList []*models.ChatSession

	database.DB.Where("user_id = ?", user.ID).Find(&chatHistoryList)

	return chatHistoryList
}

func GetMessageByChatSession(user *models.User, chatSessionId string) []*models.Message {
	var message []*models.Message

	database.DB.Where("user_id = ? AND chat_session_id = ?", user.ID, chatSessionId).Find(&message)

	return message
}

func CreateChatSession(user *models.User) *models.ChatSession {
	chatSessionId := uuid.New().String()
	chatSession := &models.ChatSession{
		ID:     chatSessionId,
		UserID: user.ID,
	}

	database.DB.Create(&chatSession)
	return chatSession
}

func CreateMessage(user *models.User, chatSessinId string, message string) *models.Message {
	messageId := uuid.New().String()
	messageObject := &models.Message{
		ID:            messageId,
		Message:       message,
		UserID:        user.ID,
		ChatSessionId: chatSessinId,
	}

	database.DB.Create(&messageObject)
	return messageObject
}
