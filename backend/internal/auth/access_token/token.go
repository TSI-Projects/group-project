package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type TokenClient struct {
	SecretKey []byte
}

func NewTokenClient() *TokenClient {
	return &TokenClient{
		// Generates a new temporary secret key on each program restart
		SecretKey: []byte(uuid.New().String()),
	}
}

func (t *TokenClient) Generate(username string) (string, error) {
	if len(username) == 0 {
		return "", fmt.Errorf("failed to generate token: username is empty")
	}

	claims := jwt.MapClaims{
		"username": username,
		"iat":      time.Now().Unix(),
		// The token will not have an expiration time as long as the "exp" field remains commented out.
		// Uncomment the "exp" field to set the token to expire in 24 hours.
		// "exp":      time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(t.SecretKey)
	if err != nil {
		return "", fmt.Errorf("failed to signed string: %v", err)
	}
	return tokenString, nil
}

func (t *TokenClient) Validate(accessToken string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return t.SecretKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse jwt: %v", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}
