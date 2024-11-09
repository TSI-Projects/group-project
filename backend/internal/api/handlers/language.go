package handlers

import (
	"encoding/json"
	"net/http"

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
		w.Write([]byte("Internal Error"))
		return
	}

	output, err := json.Marshal(languages)
	if err != nil {
		log.Errorf("Failed to marshal languages struct: %v", err)
		w.Write([]byte("Internal Error"))
		return
	}

	w.Write(output)
}

func (h *Handler) CreateLanguage(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	language := &repository.Language{}

	if err := json.NewDecoder(r.Body).Decode(language); err != nil {
		log.Errorf("failed to decode body: %v", err)
		w.Write([]byte("Internal Error"))
		return
	}

	if err := h.Validator.Validate(language); err != nil {
		log.Errorf("failed to validate language struct: %v", err)
		w.Write([]byte(utils.UppercaseFirstLetter(err.Error())))
		return
	}

	if err := h.Validator.Validate(language); err != nil {
		log.Errorf("failed to validate language struct: %v", err)
		w.Write([]byte(utils.UppercaseFirstLetter(err.Error())))
		return
	}

	if err := h.LanguageRepo.Create(language); err != nil {
		log.Errorf("failed to create language: %v", err)
		w.Write([]byte("Internal Error"))
		return
	}

	w.Write([]byte("Success"))
}

func (h *Handler) DeleteLanguage(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	rawId := mux.Vars(r)["id"]
	if len(rawId) == 0 {
		log.Errorf("failed to delete language due to empty id param")
		w.Write([]byte("Language ID is not specified"))
		return
	}

	id, err := utils.StringToUint(rawId)
	if err != nil {
		log.Errorf("failed to convert raw str id to int: %v", err)
		w.Write([]byte("Internal Error"))
		return
	}

	if err := h.LanguageRepo.Delete(id); err != nil {
		log.Errorf("failed to delete language: %v", err)
		w.Write([]byte("Internal Error"))
		return
	}

	w.Write([]byte("Success"))
}

func (h *Handler) GetLanguageByID(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	rawId := mux.Vars(r)["id"]
	if len(rawId) == 0 {
		log.Errorf("failed to delete language due to empty id param")
		w.Write([]byte("Language ID is not specified"))
		return
	}

	id, err := utils.StringToUint(rawId)
	if err != nil {
		log.Errorf("failed to convert raw str id to int: %v", err)
		w.Write([]byte("Internal Error"))
		return
	}

	language, err := h.LanguageRepo.GetByID(id)
	if err != nil {
		log.Errorf("failed to delete language: %v", err)
		w.Write([]byte("Internal Error"))
		return
	}

	output, err := json.Marshal(language)
	if err != nil {
		log.Errorf("Failed to marshal languages struct: %v", err)
		w.Write([]byte("Internal Error"))
		return
	}

	w.Write(output)
}

func (h *Handler) UpdateLanguage(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	language := &repository.Language{}

	if err := json.NewDecoder(r.Body).Decode(language); err != nil {
		log.Errorf("failed to decode body: %v", err)
		w.Write([]byte("Internal Error"))
		return
	}

	if err := h.Validator.Validate(language); err != nil {
		log.Errorf("failed to validate language struct: %v", err)
		w.Write([]byte(utils.UppercaseFirstLetter(err.Error())))
		return
	}

	if err := h.LanguageRepo.Update(language); err != nil {
		log.Errorf("failed to create language: %v", err)
		w.Write([]byte("Internal Error"))
		return
	}

	w.Write([]byte("Success"))
}
