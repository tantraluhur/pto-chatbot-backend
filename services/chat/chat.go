package chat

import (
	"bytes"
	"chatbot-backend/database"
	models "chatbot-backend/model"

	"encoding/json"
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ResponseMessage struct {
	Type    string      `json:"type"`
	Content interface{} `json:"content"`
}

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

// func CreateMessage(user *models.User, chatSessionId string, message string) *models.Message {
// 	var chatSession *models.ChatSession

// 	database.DB.Where("id = ?", chatSessionId).First(&chatSession)
// 	if chatSession.Title == nil {
// 		chatSession.Title = &message
// 		database.DB.Save(&chatSession)
// 	}

// 	botMessage := "Hello from Bot!"
// 	messageId := uuid.New().String()
// 	messageObject := &models.Message{
// 		ID:            messageId,
// 		UserMessage:   message,
// 		BotMessage:    botMessage,
// 		UserID:        user.ID,
// 		ChatSessionId: chatSessionId,
// 	}

// 	database.DB.Create(&messageObject)
// 	return messageObject
// }

func CreateMessage(user *models.User, chatSessionId string, message string) (*models.Message, *fiber.Error) {
	var chatSession *models.ChatSession

	database.DB.Where("id = ?", chatSessionId).First(&chatSession)
	if chatSession.Title == nil {
		chatSession.Title = &message
		database.DB.Save(&chatSession)
	}

	postData := map[string]interface{}{
		"question": message,
	}

	jsonData, err := json.Marshal(postData)
	if err != nil {
		return nil, fiber.NewError(500, "Failed to marshal JSON data.")
	}

	response, err := http.Post("http://192.168.153.217:8081/get_answer", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fiber.NewError(500, "Failed to make POST request.")
	}

	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fiber.NewError(500, "Failed to read response body.")
	}

	var responseData map[string][]ResponseMessage
	if err := json.Unmarshal(responseBody, &responseData); err != nil {
		return nil, fiber.NewError(500, "Failed to parse response body.")
	}

	answer := responseData["answer"]
	botMessage := ""
	imageMessage := ""

	for _, value := range answer {
		if value.Type == "text" {
			botMessage += value.Content.(string)
		}

		if (value.Type) == "image" {
			for _, value := range value.Content.([]interface{}) {
				imageMessage = imageMessage + value.(string) + ","
			}
		}

	}

	messageId := uuid.New().String()
	messageObject := &models.Message{
		ID:            messageId,
		UserMessage:   message,
		BotMessage:    botMessage,
		ImageMessage:  &imageMessage,
		UserID:        user.ID,
		ChatSessionId: chatSessionId,
	}

	database.DB.Create(&messageObject)
	return messageObject, nil
}
