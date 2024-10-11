package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/TSI-Projects/group-project/internal/repository"
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
	id, err := strconv.Atoi(rawId)
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
