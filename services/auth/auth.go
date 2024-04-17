package auth

import (
	"chatbot-backend/database"
	models "chatbot-backend/model"
	types "chatbot-backend/routes/v1/auth/types"
	"time"

	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type TokenDetails struct {
	Token string `json:"access_token"`
}

type AccessClaims struct {
	AccessTokenID string `json:"access_token_id"`
	ID            int64  `json:"user"`
	jwt.StandardClaims
}

// IssueAccessToken generate access tokens used for auth
func IssueAccessToken(user *models.User) (*TokenDetails, error) {
	tokenUUID := uuid.New().String()
	expireTime := time.Now().Add(time.Hour).Unix() // 1 hour
	// Create the JWT claims, which includes the user ID and expiry time
	claims := AccessClaims{
		tokenUUID,
		user.ID,
		jwt.StandardClaims{
			ExpiresAt: expireTime,
			Issuer:    "chatbot-backend-issuer",
		},
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

func Login(data types.LoginRequest) (*TokenDetails, *fiber.Error) {
	var user *models.User

	database.DB.Where("name = ?", data.Username).First(&user) //Check the email is present in the DB

	if user.ID == 0 { //If the ID return is '0' then there is no such email present in the DB
		return nil, fiber.NewError(401, "Username not found.")
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data.Password)); err != nil {
		return nil, fiber.NewError(400, "Incorrect username or password.")
	}

	token, err := IssueAccessToken(user)

	if err != nil {
		fiber.NewError(400, err.Error())
	}

	return token, nil
}
