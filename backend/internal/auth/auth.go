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
}

const (
	INVALID_AUTH_DATA Error = "Username or password is invalid"
)

func NewAuthClient(dbClient db.IDatabase) IAuthClient {
	return &AuthClient{
		AdminRepo: repository.NewAdminRepo(dbClient),
	}
}

func (a *AuthClient) Login(username, password string) (string, error) {
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

	token, err := token.Generate(username)
	if err != nil {
		return "", fmt.Errorf("failed to generate access token: %v", err)
	}
	return token, nil
}

func (a *AuthClient) Logout(accessToken string) {}
