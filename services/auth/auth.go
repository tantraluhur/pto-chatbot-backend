package auth

import (
	"chatbot-backend/database"
	models "chatbot-backend/model"
	types "chatbot-backend/routes/v1/auth/types"
	"time"

	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	jwtUser "github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AccessClaims struct {
	AccessTokenID string `json:"access_token_id"`
	ID            int64  `json:"user"`
	jwt.StandardClaims
}

type RefreshClaims struct {
	RefreshTokenID string `json:"refresh_token_id"`
	ID             string `json:"access_token"`
	jwt.StandardClaims
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// IssueAccessToken generate access tokens used for auth
func IssueAccessToken(user *models.User) (*models.AccessToken, error) {
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

	accesToken := models.AccessToken{
		ID:          tokenUUID,
		AccessToken: token,
	}

	database.DB.Create(&accesToken)

	return &accesToken, nil
}

func IssueRefreshToken(accesToken models.AccessToken) (*models.RefreshToken, error) {
	expireTime := time.Now().Add((24 * time.Hour) * 14) // 14 days
	tokenUUID := uuid.New().String()

	// Generate encoded token
	claims := RefreshClaims{
		tokenUUID,
		accesToken.ID,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "chatbot-backend-issuer",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte("secret"))

	if err != nil {
		return nil, err
	}

	refreshToken := models.RefreshToken{
		ID:           tokenUUID,
		RefreshToken: token,
	}

	database.DB.Create(&refreshToken)

	return &refreshToken, nil
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

func Login(data types.LoginRequest) (*LoginResponse, *fiber.Error) {
	var user *models.User

	database.DB.Where("name = ?", data.Username).First(&user)

	if user.ID == 0 {
		return nil, fiber.NewError(401, "Username not found.")
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data.Password)); err != nil {
		return nil, fiber.NewError(400, "Incorrect username or password.")
	}

	accessToken, accessTokenErr := IssueAccessToken(user)
	if accessTokenErr != nil {
		fiber.NewError(400, accessTokenErr.Error())
	}

	refreshToken, refreshTokenErr := IssueRefreshToken(*accessToken)
	if refreshTokenErr != nil {
		fiber.NewError(400, refreshTokenErr.Error())
	}

	loginResponse := LoginResponse{
		AccessToken:  accessToken.AccessToken,
		RefreshToken: refreshToken.RefreshToken,
	}

	return &loginResponse, nil
}

func Logout(c *fiber.Ctx) string {
	var accessToken *models.AccessToken
	// Get the user from the context and return it
	user := c.Locals("user").(*jwtUser.Token)
	claims := user.Claims.(jwtUser.MapClaims)
	database.DB.Where("id = ?", claims["access_token_id"]).First(&accessToken)
	database.DB.Delete(&accessToken)

	return "Logout success."
}
