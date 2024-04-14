package auth

import (
	"chatbot-backend/services/commons"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

var (
	validator              = commons.ParseBodyAndValidate
	responseValidatorError = commons.ValidatorErrorResponse
	responseParserError    = commons.ParserErrorResponse
)

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func Login(ctx *fiber.Ctx) error {
	var body LoginRequest

	errParser, errValidator := validator(ctx, &body)

	if errParser != nil {
		return ctx.Status(http.StatusBadRequest).JSON(responseParserError(errParser, "Invalid Request."))
	}

	if errValidator != nil {
		return ctx.Status(http.StatusBadRequest).JSON(responseValidatorError(errValidator, "Invalid input data."))
	}

	return ctx.Status(200).JSON(fiber.Map{
		"Message": "Hello from User!",
	})
}
