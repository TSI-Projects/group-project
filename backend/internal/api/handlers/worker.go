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

func (h *Handler) GetWorkers(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	workers, err := h.WorkerRepo.GetAll()
	if err != nil {
		log.Errorf("failed to get workers: %v", err)
		response.WriteResponseWithError(w, apiErrors.INTERNAL_ERROR_CODE, apiErrors.DEFAULT_ERROR_MESSAGE, http.StatusInternalServerError)
		return
	}

	resp, err := response.NewCustomerResponseWithDataArr("workers", workers)
	if err != nil {
		log.Errorf("Failed to create new response with customers: %v", err)
		return
	}

	response.WriteResponse(w, http.StatusOK, resp)
}

func (h *Handler) DeleteWorker(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	rawId := mux.Vars(r)["id"]
	if len(rawId) == 0 {
		log.Errorf("failed to delete worker due to empty id param")
		response.WriteResponseWithError(w, apiErrors.ID_NOT_SPECIFIED_ERROR_CODE, fmt.Sprintf("Worker %s", apiErrors.ID_NOT_SPECIFIED_ERROR_MESSAGE), http.StatusBadRequest)
		return
	}

	id, err := utils.StringToUint(rawId)
	if err != nil {
		log.Errorf("failed to convert raw str id to int: %v", err)
		response.WriteResponseWithError(w, apiErrors.NOT_FOUND_ERROR_CODE, apiErrors.NOT_FOUND_ERROR_MESSAGE, http.StatusNotFound)
		return
	}

	if err := h.WorkerRepo.Delete(id); err != nil {
		log.Errorf("failed to delete worker: %v", err)
		response.WriteResponseWithError(w, apiErrors.INTERNAL_ERROR_CODE, apiErrors.DEFAULT_ERROR_MESSAGE, http.StatusInternalServerError)
		return
	}

	response.WriteResponseWithSuccess(w)
}

func (h *Handler) CreateWorker(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	worker := &repository.Worker{}

	if err := json.NewDecoder(r.Body).Decode(worker); err != nil {
		log.Errorf("failed to decode worker: %v", err)
		response.WriteResponseWithError(w, apiErrors.INTERNAL_ERROR_CODE, apiErrors.DEFAULT_ERROR_MESSAGE, http.StatusInternalServerError)
		return
	}

	if err := h.Validator.Validate(worker); err != nil {
		log.Errorf("failed to validate worker struct: %v", err)
		response.WriteResponseWithError(w, apiErrors.FIELD_VALIDATION_ERROR_CODE, apiErrors.FIELD_VALIDATION_ERROR_MESSAGE, http.StatusBadRequest)
		return
	}

	if err := h.WorkerRepo.Create(worker); err != nil {
		log.Errorf("failed to create new worker: %v", err)
		response.WriteResponseWithError(w, apiErrors.INTERNAL_ERROR_CODE, apiErrors.DEFAULT_ERROR_MESSAGE, http.StatusInternalServerError)
		return
	}

	response.WriteResponseWithSuccess(w)
}

func (h *Handler) GetWorkerByID(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	rawId := mux.Vars(r)["id"]
	if len(rawId) == 0 {
		log.Errorf("failed to delete worker due to empty id param")
		response.WriteResponseWithError(w, apiErrors.ID_NOT_SPECIFIED_ERROR_CODE, fmt.Sprintf("Worker %s", apiErrors.ID_NOT_SPECIFIED_ERROR_MESSAGE), http.StatusBadRequest)
		return
	}

	id, err := utils.StringToUint(rawId)
	if err != nil {
		log.Errorf("failed to convert raw str id to int: %v", err)
		response.WriteResponseWithError(w, apiErrors.NOT_FOUND_ERROR_CODE, apiErrors.NOT_FOUND_ERROR_MESSAGE, http.StatusNotFound)
		return
	}

	worker, err := h.WorkerRepo.GetByID(id)
	if err != nil {
		log.Errorf("failed to delete worker: %v", err)
		response.WriteResponseWithError(w, apiErrors.NOT_FOUND_ERROR_CODE, "No customer found with the provided ID.", http.StatusNotFound)
		return
	}

	resp, err := response.NewCustomerResponseWithData("worker", worker)
	if err != nil {
		log.Errorf("Failed to create new response with customers: %v", err)
		return
	}

	response.WriteResponse(w, http.StatusOK, resp)
}

func (h *Handler) UpdateWorker(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	worker := &repository.Worker{}

	if err := json.NewDecoder(r.Body).Decode(worker); err != nil {
		log.Errorf("failed to decode body: %v", err)
		response.WriteResponseWithError(w, apiErrors.INTERNAL_ERROR_CODE, apiErrors.INTERNAL_ERROR_MESSAGE, http.StatusInternalServerError)
		return
	}

	if err := h.Validator.Validate(worker); err != nil {
		log.Errorf("failed to validate worker struct: %v", err)
		response.WriteResponseWithError(w, apiErrors.FIELD_VALIDATION_ERROR_CODE, apiErrors.FIELD_VALIDATION_ERROR_MESSAGE, http.StatusBadRequest)
		return
	}

	if err := h.WorkerRepo.Update(worker); err != nil {
		log.Errorf("failed to create worker: %v", err)
		response.WriteResponseWithError(w, apiErrors.INTERNAL_ERROR_CODE, apiErrors.INTERNAL_ERROR_MESSAGE, http.StatusInternalServerError)
		return
	}

	response.WriteResponseWithSuccess(w)
}
