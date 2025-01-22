package middleware

import (
	"net/http"
	"strings"

	token "github.com/TSI-Projects/group-project/internal/auth/access_token"
	response "github.com/TSI-Projects/group-project/internal/models/responses"
	log "github.com/sirupsen/logrus"
)

func AuthMW(next http.HandlerFunc, token *token.TokenClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if len(authHeader) == 0 {
			log.Error("Failed to pass authorization due to missing authorization header")
			response.WriteResponseWithError(w, "MISSING_HEADER", "Missing Authorization header", http.StatusBadRequest)
			return
		}

		accessToken := strings.TrimPrefix(authHeader, "Bearer ")
		if _, err := token.Validate(accessToken); err != nil {
			log.Errorf("Access Token: '%s', Failed to validate access token:'%v'", accessToken, err)
			response.WriteResponseWithError(w, "INVALID_TOKEN", "Invalid or expired token", http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	}
}
