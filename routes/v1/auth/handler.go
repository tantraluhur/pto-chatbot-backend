package auth

import (
	types "chatbot-backend/routes/v1/auth/types"
	"chatbot-backend/services/auth"
	"chatbot-backend/services/commons"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

var (
	validator              = commons.ParseBodyAndValidate
	responseValidatorError = commons.ValidatorErrorResponse
	responseParserError    = commons.ParserErrorResponse
	responseSuccess        = commons.HTTPResponse
	responseError          = commons.HTTPErrorResponse
)

func Register(ctx *fiber.Ctx) error {
	var body types.RegisterRequest

	errParser, errValidator := validator(ctx, &body)
	if errParser != nil {
		return ctx.Status(http.StatusBadRequest).JSON(responseParserError(errParser, "Invalid Request."))
	}

	if errValidator != nil {
		return ctx.Status(http.StatusBadRequest).JSON(responseValidatorError(errValidator, "Invalid input data."))
	}

	user := auth.Register(body)

	return ctx.Status(http.StatusCreated).JSON(responseSuccess(user))

}

func Login(ctx *fiber.Ctx) error {
	var body types.LoginRequest

	errParser, errValidator := validator(ctx, &body)

	if errParser != nil {
		return ctx.Status(http.StatusBadRequest).JSON(responseParserError(errParser, "Invalid Request."))
	}

	if errValidator != nil {
		return ctx.Status(http.StatusBadRequest).JSON(responseValidatorError(errValidator, "Invalid input data."))
	}

	token, err := auth.Login(body)
	if err != nil {
		return ctx.Status(err.Code).JSON(responseError(err.Message))
	}

	return ctx.Status(200).JSON(responseSuccess(token))
}

func Logout(ctx *fiber.Ctx) error {
	logoutMessage := auth.Logout(ctx)
	return ctx.Status(200).JSON(responseSuccess(logoutMessage))
}
