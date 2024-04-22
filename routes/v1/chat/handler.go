package chat

import (
	"chatbot-backend/routes/v1/chat/types"
	"chatbot-backend/services/chat"
	"chatbot-backend/services/commons"
	"chatbot-backend/services/user"

	"github.com/gofiber/fiber/v2"
)

func GetAllChatSession(ctx *fiber.Ctx) error {
	user, err := user.GetLoggedInUser(ctx)

	if err != nil {
		return ctx.Status(err.Code).JSON(commons.HTTPErrorResponse(err.Message))
	}

	chatHistoryList := chat.GetAllChatSession(user)

	return ctx.Status(200).JSON(commons.HTTPResponse(chatHistoryList))
}

func GetMessageByChatSession(ctx *fiber.Ctx) error {
	user, err := user.GetLoggedInUser(ctx)

	if err != nil {
		return ctx.Status(err.Code).JSON(commons.HTTPErrorResponse(err.Message))
	}

	chatSessionId := ctx.Params("id")
	messageList := chat.GetMessageByChatSession(user, chatSessionId)

	return ctx.Status(200).JSON(commons.HTTPResponse(messageList))
}

func CreateChatSession(ctx *fiber.Ctx) error {
	user, err := user.GetLoggedInUser(ctx)

	if err != nil {
		return ctx.Status(err.Code).JSON(commons.HTTPErrorResponse(err.Message))
	}

	chatSession := chat.CreateChatSession(user)

	return ctx.Status(201).JSON(commons.HTTPResponse(chatSession))
}

func CreateMessage(ctx *fiber.Ctx) error {
	var body types.SubmitMessageRequest
	user, err := user.GetLoggedInUser(ctx)

	if err != nil {
		return ctx.Status(err.Code).JSON(commons.HTTPErrorResponse(err.Message))
	}

	errParser, errValidator := commons.ParseBodyAndValidate(ctx, &body)
	if errParser != nil {
		return ctx.Status(400).JSON(commons.ParserErrorResponse(errParser, "Invalid Request."))
	}

	if errValidator != nil {
		return ctx.Status(400).JSON(commons.ValidatorErrorResponse(errValidator, "Invalid input data."))
	}

	chatSessionId := ctx.Params("id")
	message := chat.CreateMessage(user, chatSessionId, body.Message)
	return ctx.Status(201).JSON(commons.HTTPResponse(message))
}
