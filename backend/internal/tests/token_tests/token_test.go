package token_test

import (
	"testing"

	token "github.com/TSI-Projects/group-project/internal/auth/access_token"
	"github.com/stretchr/testify/require"
)

func TestGenerateTokenWithEmptyString(t *testing.T) {
	token := token.NewTokenClient()
	accessToken, err := token.Generate("")
	require.Error(t, err)
	require.Empty(t, accessToken)
}

func TestGenerateAndValidate(t *testing.T) {
	token := token.NewTokenClient()
	accessToken, err := token.Generate("testuser")
	require.NoError(t, err)
	require.NotEmpty(t, accessToken)

	claims, err := token.Validate(accessToken)
	require.NoError(t, err)
	require.NotNil(t, claims)
	require.Equal(t, "testuser", (*claims)["username"])
}

func TestValidateInvalidToken(t *testing.T) {
	token := token.NewTokenClient()
	claims, err := token.Validate("invalid token string")
	require.Error(t, err)
	require.Nil(t, claims)
}
