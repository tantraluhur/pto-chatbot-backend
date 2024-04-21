package user

import (
	"chatbot-backend/database"
	models "chatbot-backend/model"

	"github.com/gofiber/fiber/v2"
	jwtUser "github.com/golang-jwt/jwt/v4"
)

func GetLoggedInUser(c *fiber.Ctx) (*models.User, *fiber.Error) {
	var userObject *models.User
	var accessToken *models.AccessToken
	// Get the user from the context and return it
	user := c.Locals("user").(*jwtUser.Token)
	claims := user.Claims.(jwtUser.MapClaims)
	database.DB.Where("id = ?", claims["user"]).First(&userObject)
	database.DB.Where("id = ?", claims["access_token_id"]).First(&accessToken)

	if accessToken.ID == "" {
		return nil, fiber.NewError(401, "Invalid access token.")
	}

	if userObject.ID == 0 {
		return nil, fiber.NewError(401, "Invalid access token.")
	}

	return userObject, nil
}
