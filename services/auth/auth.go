package auth

import (
	user "chatbot-backend/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenDetails struct {
	Token string
}

// IssueAccessToken generate access tokens used for auth
func IssueAccessToken(user user.User) (*TokenDetails, error) {

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
