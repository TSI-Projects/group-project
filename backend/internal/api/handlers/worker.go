package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/TSI-Projects/group-project/internal/repository"
	"github.com/TSI-Projects/group-project/utils"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func (h *Handler) GetWorkers(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	workers, err := h.WorkerRepo.GetAll()
	if err != nil {
		log.Errorf("failed to get workers: %v", err)
		w.Write([]byte("Internal Error"))
		return
	}

	output, err := json.Marshal(workers)
	if err != nil {
		log.Errorf("failed to marshal workers struct: %v", err)
		w.Write([]byte("Internal Error"))
		return
	}

	w.Write(output)
}

func (h *Handler) DeleteWorker(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	rawId := mux.Vars(r)["id"]
	if len(rawId) == 0 {
		log.Errorf("failed to delete worker due to empty id param")
		w.Write([]byte("Worker ID is not specified"))
		return
	}

	id, err := strconv.Atoi(rawId)
	if err != nil {
		log.Errorf("failed to convert raw str id to int: %v", err)
		w.Write([]byte("Internal Error"))
		return
	}

	if err := h.WorkerRepo.Delete(id); err != nil {
		log.Errorf("failed to delete worker: %v", err)
		w.Write([]byte("Internal Error"))
		return
	}

	w.Write([]byte("Success"))
}

func (h *Handler) CreateWorker(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	worker := &repository.Worker{}

	if err := json.NewDecoder(r.Body).Decode(worker); err != nil {
		log.Errorf("failed to decode worker: %v", err)
		w.Write([]byte("Internal Error"))
		return
	}

	if err := h.Validator.Validate(worker); err != nil {
		log.Errorf("failed to validate worker struct: %v", err)
		w.Write([]byte(utils.UppercaseFirstLetter(err.Error())))
		return
	}

	if err := h.Validator.Validate(worker); err != nil {
		log.Errorf("failed to validate worker struct: %v", err)
		w.Write([]byte(utils.UppercaseFirstLetter(err.Error())))
		return
	}

	if err := h.WorkerRepo.Create(worker); err != nil {
		log.Errorf("failed to create new worker: %v", err)
		w.Write([]byte("Internal Error"))
		return
	}

	w.Write([]byte("Success"))
}
