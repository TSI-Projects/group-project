package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/TSI-Projects/group-project/internal/repository"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func (h *Handler) GetWorkers(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	workers, err := h.WorkerRepo.GetWorkers()
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

	rawID := mux.Vars(r)["id"]
	id, err := strconv.Atoi(rawID)
	if err != nil {
		log.Errorf("failed to convert raw str id to int: %v", err)
		w.Write([]byte("Internal Error"))
		return
	}

	if err := h.WorkerRepo.DeleteWorker(id); err != nil {
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

	if err := h.WorkerRepo.CreateWorker(worker); err != nil {
		log.Errorf("failed to create new worker: %v", err)
		w.Write([]byte("Internal Error"))
		return
	}

	w.Write([]byte("Success"))
}
