package auth

import (
	"chatbot-backend/database"
	models "chatbot-backend/model"
	types "chatbot-backend/routes/v1/auth/types"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type TokenDetails struct {
	Token string
}

// IssueAccessToken generate access tokens used for auth
func IssueAccessToken(user models.User) (*TokenDetails, error) {

	expireTime := time.Now().Add(time.Hour).Unix() // 1 hour
	// Create the JWT claims, which includes the user ID and expiry time
	claims := jwt.MapClaims{
		"ID":    user.ID,
		"email": user.Email,
		"exp":   expireTime,
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte("secret"))

	if err != nil {
		return nil, err
	}

	TokenDetails := TokenDetails{
		Token: token,
	}

	return &TokenDetails, nil
}

func Register(data types.RegisterRequest) models.User {
	password, _ := bcrypt.GenerateFromPassword([]byte(data.Password), 14)

	user := models.User{
		Name:     data.Username,
		Email:    data.Email,
		Password: password,
	}

	database.DB.Create(&user)

	return user
}

func Login(data types.LoginRequest) (*models.User, *fiber.Error) {
	var user *models.User

	database.DB.Where("name = ?", data.Username).First(&user) //Check the email is present in the DB

	if user.ID == 0 { //If the ID return is '0' then there is no such email present in the DB
		return nil, fiber.NewError(401, "Username not found.")
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data.Password)); err != nil {
		return nil, fiber.NewError(400, "Incorrect username or password.")
	}

	return user, nil
}
