package auth

import (
	"fmt"

	token "github.com/TSI-Projects/group-project/internal/auth/access_token"
	"github.com/TSI-Projects/group-project/internal/db"
	"github.com/TSI-Projects/group-project/internal/repository"

	log "github.com/sirupsen/logrus"
)

type Error string

func (e Error) Error() string {
	return string(e)
}

type AuthClient struct {
	AdminRepo repository.IAdminRepository
	Token     token.TokenClient
}

const (
	INVALID_AUTH_DATA Error = "Username or password is invalid"
)

func NewAuthClient(dbClient db.IDatabase, tokenClient *token.TokenClient) IAuthClient {
	return &AuthClient{
		AdminRepo: repository.NewAdminRepo(dbClient),
		Token:     *tokenClient,
	}
}

func (a *AuthClient) Login(username, password string) (string, error) {
	if len(username) == 0 {
		log.Errorln("Username is not specified")
		return "", INVALID_AUTH_DATA
	}

	if len(password) == 0 {
		log.Errorln("Password is not specified")
		return "", INVALID_AUTH_DATA
	}

	admin, err := a.AdminRepo.GetByUsername(username)
	if err != nil {
		log.Errorf("Failed to get admin from database: %v", err)
		return "", INVALID_AUTH_DATA
	}

	if admin == nil {
		log.Errorf("User %s not found", username)
		return "", INVALID_AUTH_DATA
	}

	if password != admin.Password {
		log.Errorf("Invalid password '%s' for user '%s'", password, username)
		return "", INVALID_AUTH_DATA
	}

	token, err := a.Token.Generate(username)
	if err != nil {
		return "", fmt.Errorf("failed to generate access token: %v", err)
	}
	return token, nil
}

func (a *AuthClient) Logout(accessToken string) {}
