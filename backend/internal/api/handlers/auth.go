package handlers

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	user := struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Errorf("Failed to decode body: %v", err)
		http.Error(w, "Incorrect payload", http.StatusForbidden)
		return
	}

	if err := h.Validator.Validate(user); err != nil {
		log.Errorf("Failed to validate user login data: %v", err)
		http.Error(w, "Incorrect payload", http.StatusForbidden)
	}

	accessToken, err := h.AuthClient.Login(user.Username, user.Password)
	if err != nil {
		log.Errorf("Failed to login using username: '%s', error: %v", user.Username, err)
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	w.Write([]byte(accessToken))
}
