package handlers

import (
	"fmt"
	"net/http"

	"github.com/goccy/go-json"

	apiErrors "github.com/TSI-Projects/group-project/internal/errors"
	response "github.com/TSI-Projects/group-project/internal/models/responses"
	"github.com/TSI-Projects/group-project/internal/repository"
	"github.com/TSI-Projects/group-project/utils"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func (h *Handler) GetLanguages(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	languages, err := h.LanguageRepo.GetAll()
	if err != nil {
		log.Errorf("Failed to get languages: %v", err)
		response.WriteResponseWithError(w, apiErrors.INTERNAL_ERROR_CODE, apiErrors.DEFAULT_ERROR_MESSAGE, http.StatusInternalServerError)
		return
	}

	resp, err := response.NewCustomerResponseWithDataArr("languages", languages)
	if err != nil {
		log.Errorf("Failed to create new response with languages: %v", err)
		return
	}

	response.WriteResponse(w, http.StatusOK, resp)
}

func (h *Handler) CreateLanguage(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	language := &repository.Language{}

	if err := json.NewDecoder(r.Body).Decode(language); err != nil {
		log.Errorf("failed to decode body: %v", err)
		response.WriteResponseWithError(w, apiErrors.INTERNAL_ERROR_CODE, apiErrors.DEFAULT_ERROR_MESSAGE, http.StatusInternalServerError)
		return
	}

	if err := h.Validator.Validate(language); err != nil {
		log.Errorf("failed to validate language struct: %v", err)
		response.WriteResponseWithError(w, apiErrors.FIELD_VALIDATION_ERROR_CODE, apiErrors.FIELD_VALIDATION_ERROR_MESSAGE, http.StatusBadRequest)
		return
	}

	if err := h.LanguageRepo.Create(language); err != nil {
		log.Errorf("failed to create language: %v", err)
		response.WriteResponseWithError(w, apiErrors.INTERNAL_ERROR_CODE, apiErrors.DEFAULT_ERROR_MESSAGE, http.StatusInternalServerError)
		return
	}

	response.WriteResponseWithSuccess(w)
}

func (h *Handler) DeleteLanguage(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	rawId := mux.Vars(r)["id"]
	if len(rawId) == 0 {
		log.Errorf("failed to delete language due to empty id param")
		response.WriteResponseWithError(w, apiErrors.ID_NOT_SPECIFIED_ERROR_CODE, fmt.Sprintf("Language %s", apiErrors.ID_NOT_SPECIFIED_ERROR_MESSAGE), http.StatusBadRequest)
		return
	}

	id, err := utils.StringToUint(rawId)
	if err != nil {
		log.Errorf("failed to convert raw str id to int: %v", err)
		response.WriteResponseWithError(w, apiErrors.NOT_FOUND_ERROR_CODE, apiErrors.NOT_FOUND_ERROR_MESSAGE, http.StatusNotFound)
		return
	}

	if err := h.LanguageRepo.Delete(id); err != nil {
		log.Errorf("failed to delete language: %v", err)
		response.WriteResponseWithError(w, apiErrors.INTERNAL_ERROR_CODE, apiErrors.DEFAULT_ERROR_MESSAGE, http.StatusInternalServerError)
		return
	}

	response.WriteResponseWithSuccess(w)
}

func (h *Handler) GetLanguageByID(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	rawId := mux.Vars(r)["id"]
	if len(rawId) == 0 {
		log.Errorf("failed to delete language due to empty id param")
		response.WriteResponseWithError(w, apiErrors.ID_NOT_SPECIFIED_ERROR_CODE, fmt.Sprintf("Language %s", apiErrors.ID_NOT_SPECIFIED_ERROR_MESSAGE), http.StatusBadRequest)
		return
	}

	id, err := utils.StringToUint(rawId)
	if err != nil {
		log.Errorf("failed to convert raw str id to int: %v", err)
		response.WriteResponseWithError(w, apiErrors.NOT_FOUND_ERROR_CODE, apiErrors.NOT_FOUND_ERROR_MESSAGE, http.StatusNotFound)
		return
	}

	language, err := h.LanguageRepo.GetByID(id)
	if err != nil {
		log.Errorf("failed to delete language: %v", err)
		response.WriteResponseWithError(w, apiErrors.NOT_FOUND_ERROR_CODE, "No language found with the provided ID.", http.StatusNotFound)
		return
	}

	resp, err := response.NewCustomerResponseWithData("language", language)
	if err != nil {
		log.Errorf("Failed to create new response with customers: %v", err)
		return
	}

	response.WriteResponse(w, http.StatusOK, resp)
}

func (h *Handler) UpdateLanguage(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	language := &repository.Language{}

	if err := json.NewDecoder(r.Body).Decode(language); err != nil {
		log.Errorf("failed to decode body: %v", err)
		response.WriteResponseWithError(w, apiErrors.INTERNAL_ERROR_CODE, apiErrors.INTERNAL_ERROR_MESSAGE, http.StatusInternalServerError)
		return
	}

	if err := h.Validator.Validate(language); err != nil {
		log.Errorf("failed to validate language struct: %v", err)
		response.WriteResponseWithError(w, apiErrors.FIELD_VALIDATION_ERROR_CODE, apiErrors.FIELD_VALIDATION_ERROR_MESSAGE, http.StatusBadRequest)
		return
	}

	if err := h.LanguageRepo.Update(language); err != nil {
		log.Errorf("failed to create language: %v", err)
		response.WriteResponseWithError(w, apiErrors.INTERNAL_ERROR_CODE, apiErrors.INTERNAL_ERROR_MESSAGE, http.StatusInternalServerError)
		return
	}

	response.WriteResponseWithSuccess(w)
}
