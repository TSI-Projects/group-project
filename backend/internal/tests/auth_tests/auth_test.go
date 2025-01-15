package auth_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/TSI-Projects/group-project/internal/auth"
	"github.com/TSI-Projects/group-project/internal/repository"
)

type mockAdminRepo struct {
	adminData  map[string]*repository.Admin
	forceError bool
}

func (m *mockAdminRepo) GetByUsername(username string) (*repository.Admin, error) {
	if m.forceError {
		return nil, errors.New("forced error from repository")
	}

	admin, ok := m.adminData[username]
	if !ok {
		return nil, nil
	}
	return admin, nil
}

func (m *mockAdminRepo) CreateAdmin(admin *repository.Admin) error {
	return nil
}

func (m *mockAdminRepo) UpdateAdmin(admin *repository.Admin) error {
	return nil
}

func (m *mockAdminRepo) DeleteAdmin(username string) error {
	return nil
}

func (m *mockAdminRepo) ListAdmins() ([]*repository.Admin, error) {
	return nil, nil
}

func TestLoginSuccess(t *testing.T) {
	// Готовим данные
	mockRepo := &mockAdminRepo{
		adminData: map[string]*repository.Admin{
			"alice": {
				Username: "alice",
				Password: "123456",
			},
		},
		forceError: false,
	}

	// Создаём AuthClient с нашим мок-репо
	authClient := &auth.AuthClient{
		AdminRepo: mockRepo,
	}

	tokenString, err := authClient.Login("alice", "123456")

	// Проверяем результат
	require.NoError(t, err, "ожидается, что логин успешен, ошибка не вернётся")
	require.NotEmpty(t, tokenString, "при успешном логине токен не должен быть пустым")
}

func TestLoginWithEmptyUsername(t *testing.T) {
	mockRepo := &mockAdminRepo{
		adminData:  map[string]*repository.Admin{},
		forceError: false,
	}
	authClient := &auth.AuthClient{
		AdminRepo: mockRepo,
	}

	tokenString, err := authClient.Login("", "123456")
	require.Error(t, err, "should be an error due to empty username")
	require.Empty(t, tokenString, "token should be empty")
	require.Equal(t, auth.INVALID_AUTH_DATA, err, "error should be equal to INVALID_AUTH_DATA")
}

func TestLoginWithEmptyPassword(t *testing.T) {
	mockRepo := &mockAdminRepo{
		adminData:  map[string]*repository.Admin{},
		forceError: false,
	}
	authClient := &auth.AuthClient{
		AdminRepo: mockRepo,
	}

	tokenString, err := authClient.Login("andrew", "")
	require.Error(t, err, "should be an error due to empty password")
	require.Empty(t, tokenString, "token should be empty")
	require.Equal(t, auth.INVALID_AUTH_DATA, err, "error should be equal to INVALID_AUTH_DATA")
}

func TestLoginUserNotFound(t *testing.T) {
	mockRepo := &mockAdminRepo{
		adminData:  map[string]*repository.Admin{},
		forceError: false,
	}
	authClient := &auth.AuthClient{
		AdminRepo: mockRepo,
	}

	tokenString, err := authClient.Login("vadims_toxic_228", "123456")
	require.Error(t, err, "expected error due to user not found")
	require.Empty(t, tokenString, "token should be empty")
	require.Equal(t, auth.INVALID_AUTH_DATA, err, "error should be equal to INVALID_AUTH_DATA")
}

func TestLoginWrongPassword(t *testing.T) {
	mockRepo := &mockAdminRepo{
		adminData: map[string]*repository.Admin{
			"bob": {
				Username: "bob",
				Password: "correct_pass",
			},
		},
	}
	authClient := &auth.AuthClient{
		AdminRepo: mockRepo,
	}

	tokenString, err := authClient.Login("bob", "wrong_pass")
	require.Error(t, err, "expected error due to wrong password")
	require.Empty(t, tokenString, "token should be empty")
	require.Equal(t, auth.INVALID_AUTH_DATA, err)
}

func TestLoginRepositoryError(t *testing.T) {
	mockRepo := &mockAdminRepo{
		adminData: map[string]*repository.Admin{
			"admin": {
				Username: "admin",
				Password: "admin_pass",
			},
		},
		forceError: true,
	}
	authClient := &auth.AuthClient{
		AdminRepo: mockRepo,
	}

	tokenString, err := authClient.Login("admin", "admin_pass")
	require.Error(t, err, "Error expected when simulating repository failure (forceError=true)")
	require.Empty(t, tokenString, "token should be empty on repository error")
	require.Equal(t, auth.INVALID_AUTH_DATA, err)
}
