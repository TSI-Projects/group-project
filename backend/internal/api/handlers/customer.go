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

func (h *Handler) GetCustomers(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	customers, err := h.CustomerRepo.GetAll()
	if err != nil {
		log.Errorf("Failed to get customers: %v", err)
		response.WriteResponseWithError(w, apiErrors.INTERNAL_ERROR_CODE, apiErrors.DEFAULT_ERROR_MESSAGE, http.StatusInternalServerError)
		return
	}

	resp, err := response.NewCustomerResponseWithDataArr("customers", customers)
	if err != nil {
		log.Errorf("Failed to create new response with customers: %v", err)
		return
	}

	response.WriteResponse(w, http.StatusOK, resp)
}

func (h *Handler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	customer := &repository.Customer{}

	if err := json.NewDecoder(r.Body).Decode(customer); err != nil {
		log.Errorf("failed to decode body: %v", err)
		response.WriteResponseWithError(w, apiErrors.INTERNAL_ERROR_CODE, apiErrors.DEFAULT_ERROR_MESSAGE, http.StatusInternalServerError)
		return
	}

	if err := h.Validator.Validate(customer); err != nil {
		log.Errorf("failed to validate customer struct: %v", err)
		response.WriteResponseWithError(w, apiErrors.FIELD_VALIDATION_ERROR_CODE, apiErrors.FIELD_VALIDATION_ERROR_MESSAGE, http.StatusBadRequest)
		return
	}

	if err := h.CustomerRepo.Create(customer); err != nil {
		log.Errorf("failed to create customer: %v", err)
		response.WriteResponseWithError(w, apiErrors.INTERNAL_ERROR_CODE, apiErrors.DEFAULT_ERROR_MESSAGE, http.StatusInternalServerError)
		return
	}

	response.WriteResponseWithSuccess(w)
}

func (h *Handler) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	rawId := mux.Vars(r)["id"]
	if len(rawId) == 0 {
		log.Errorf("failed to delete customer due to empty id param")
		response.WriteResponseWithError(w, apiErrors.ID_NOT_SPECIFIED_ERROR_CODE, fmt.Sprintf("Customer %s", apiErrors.ID_NOT_SPECIFIED_ERROR_MESSAGE), http.StatusBadRequest)
		return
	}

	id, err := utils.StringToUint(rawId)
	if err != nil {
		log.Errorf("failed to convert raw str id to int: %v", err)
		response.WriteResponseWithError(w, apiErrors.NOT_FOUND_ERROR_CODE, apiErrors.NOT_FOUND_ERROR_MESSAGE, http.StatusNotFound)
		return
	}

	if err := h.CustomerRepo.Delete(id); err != nil {
		log.Errorf("failed to delete customer: %v", err)
		response.WriteResponseWithError(w, apiErrors.INTERNAL_ERROR_CODE, apiErrors.DEFAULT_ERROR_MESSAGE, http.StatusInternalServerError)
		return
	}

	response.WriteResponseWithSuccess(w)
}

func (h *Handler) GetCustomerByID(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	rawId := mux.Vars(r)["id"]
	if len(rawId) == 0 {
		log.Errorf("failed to delete customer due to empty id param")
		response.WriteResponseWithError(w, apiErrors.ID_NOT_SPECIFIED_ERROR_CODE, fmt.Sprintf("Customer %s", apiErrors.ID_NOT_SPECIFIED_ERROR_MESSAGE), http.StatusBadRequest)
		return
	}

	id, err := utils.StringToUint(rawId)
	if err != nil {
		log.Errorf("failed to convert raw str id to int: %v", err)
		response.WriteResponseWithError(w, apiErrors.NOT_FOUND_ERROR_CODE, apiErrors.NOT_FOUND_ERROR_MESSAGE, http.StatusNotFound)
		return
	}

	customer, err := h.CustomerRepo.GetByID(id)
	if err != nil {
		log.Errorf("failed to delete customer: %v", err)
		response.WriteResponseWithError(w, apiErrors.NOT_FOUND_ERROR_CODE, "No customer found with the provided ID.", http.StatusNotFound)
		return
	}

	resp, err := response.NewCustomerResponseWithData("customer", customer)
	if err != nil {
		log.Errorf("Failed to create new response with customers: %v", err)
		return
	}

	response.WriteResponse(w, http.StatusOK, resp)
}

func (h *Handler) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	customer := &repository.Customer{}

	if err := json.NewDecoder(r.Body).Decode(customer); err != nil {
		log.Errorf("failed to decode body: %v", err)
		response.WriteResponseWithError(w, apiErrors.INTERNAL_ERROR_CODE, apiErrors.INTERNAL_ERROR_MESSAGE, http.StatusInternalServerError)
		return
	}

	if err := h.Validator.Validate(customer); err != nil {
		log.Errorf("failed to validate customer struct: %v", err)
		response.WriteResponseWithError(w, apiErrors.FIELD_VALIDATION_ERROR_CODE, apiErrors.FIELD_VALIDATION_ERROR_MESSAGE, http.StatusBadRequest)
		return
	}

	if err := h.CustomerRepo.Update(customer); err != nil {
		log.Errorf("failed to create customer: %v", err)
		response.WriteResponseWithError(w, apiErrors.INTERNAL_ERROR_CODE, apiErrors.INTERNAL_ERROR_MESSAGE, http.StatusInternalServerError)
		return
	}

	response.WriteResponseWithSuccess(w)
}
