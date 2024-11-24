package middleware

import (
	"net/http"
	"strings"

	token "github.com/TSI-Projects/group-project/internal/auth/access_token"
	log "github.com/sirupsen/logrus"
)

func AuthMW(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if len(authHeader) == 0 {
			log.Errorln("Failed to pass authorization due to missing authorization header")
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		accessToken := strings.TrimPrefix(authHeader, "Bearer ")
		if _, err := token.Validate(accessToken); err != nil {
			log.Errorf("Access Token: '%s', Failed to validate access token:'%v'", accessToken, err)
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}
