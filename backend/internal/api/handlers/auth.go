package handlers

import (
	"encoding/json"
	"net/http"

	response "github.com/TSI-Projects/group-project/internal/models/responses"
	log "github.com/sirupsen/logrus"
)

const (
	INVALID_AUTH_DATA_CODE    = "INVALID_AUTH_DATA"
	INCORRECT_PAYLOAD_CODE    = "INCORRECT_PAYLOAD"
	INCORRECT_PAYLOAD_MESSAGE = "The payload is invalid. Both 'password' and 'username' are required."
)

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	user := struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Errorf("Failed to decode body: %v", err)
		sendErrorResponse(w, http.StatusBadRequest, INCORRECT_PAYLOAD_CODE, INCORRECT_PAYLOAD_MESSAGE)
		return
	}

	if err := h.Validator.Validate(user); err != nil {
		log.Errorf("Failed to validate user login data: %v", err)
		sendErrorResponse(w, http.StatusBadRequest, INCORRECT_PAYLOAD_CODE, INCORRECT_PAYLOAD_MESSAGE)
		return
	}

	accessToken, err := h.AuthClient.Login(user.Username, user.Password)
	if err != nil {
		log.Errorf("Failed to login using username: '%s', error: %v", user.Username, err)
		sendErrorResponse(w, http.StatusForbidden, INVALID_AUTH_DATA_CODE, err.Error())
		return
	}

	resp, err := response.NewLoginResponseWithAccessToken(accessToken)
	if err != nil {
		log.Errorf("Failed to get login response with access token: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(resp); err != nil {
		log.Errorf("Failed to write response: %v", err)
	}
}

func sendErrorResponse(w http.ResponseWriter, statusCode int, errCode, errMsg string) {
	resp, err := response.NewLoginResponseWithError(errCode, errMsg)
	if err != nil {
		log.Errorf("Failed to get login response with error: %v", err)
		return
	}

	w.WriteHeader(statusCode)
	if _, err := w.Write(resp); err != nil {
		log.Errorf("Failed to write response: %v", err)
	}
}
