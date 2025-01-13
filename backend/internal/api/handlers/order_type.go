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

func (h *Handler) GetOrderTypes(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	orderTypes, err := h.OrderTypeRepo.GetAll()
	if err != nil {
		log.Errorf("failed to get order types: %v", err)
		response.WriteResponseWithError(w, apiErrors.INTERNAL_ERROR_CODE, apiErrors.DEFAULT_ERROR_MESSAGE, http.StatusInternalServerError)
		return
	}

	resp, err := response.NewCustomerResponseWithDataArr("order_types", orderTypes)
	if err != nil {
		log.Errorf("Failed to create new response with customers: %v", err)
		return
	}

	response.WriteResponse(w, http.StatusOK, resp)
}

func (h *Handler) CreateOrderType(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	orderType := &repository.OrderType{}

	if err := json.NewDecoder(r.Body).Decode(orderType); err != nil {
		log.Errorf("failed to decode body: %v", err)
		response.WriteResponseWithError(w, apiErrors.INTERNAL_ERROR_CODE, apiErrors.DEFAULT_ERROR_MESSAGE, http.StatusInternalServerError)
		return
	}

	if err := h.Validator.Validate(orderType); err != nil {
		log.Errorf("failed to validate order type struct: %v", err)
		response.WriteResponseWithError(w, apiErrors.FIELD_VALIDATION_ERROR_CODE, apiErrors.FIELD_VALIDATION_ERROR_MESSAGE, http.StatusBadRequest)
		return
	}

	if err := h.OrderTypeRepo.Create(orderType); err != nil {
		log.Errorf("failed to create order type: %v", err)
		response.WriteResponseWithError(w, apiErrors.INTERNAL_ERROR_CODE, apiErrors.DEFAULT_ERROR_MESSAGE, http.StatusInternalServerError)
		return
	}

	response.WriteResponseWithSuccess(w)
}

func (h *Handler) DeleteOrderType(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	rawId := mux.Vars(r)["id"]
	if len(rawId) == 0 {
		log.Errorf("failed to delete order type due to empty id param")
		response.WriteResponseWithError(w, apiErrors.ID_NOT_SPECIFIED_ERROR_CODE, fmt.Sprintf("Order type %s", apiErrors.ID_NOT_SPECIFIED_ERROR_MESSAGE), http.StatusBadRequest)
		return
	}

	id, err := utils.StringToUint(rawId)
	if err != nil {
		log.Errorf("failed to convert str id to int: id '%s': %v ", rawId, err)
		response.WriteResponseWithError(w, apiErrors.NOT_FOUND_ERROR_CODE, apiErrors.NOT_FOUND_ERROR_MESSAGE, http.StatusNotFound)
		return
	}

	if err := h.OrderTypeRepo.Delete(id); err != nil {
		log.Errorf("failed to delete order type item: %s", err)
		response.WriteResponseWithError(w, apiErrors.INTERNAL_ERROR_CODE, apiErrors.DEFAULT_ERROR_MESSAGE, http.StatusInternalServerError)
		return
	}

	response.WriteResponseWithSuccess(w)
}

func (h *Handler) GetOrderTypeByID(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	rawId := mux.Vars(r)["id"]
	if len(rawId) == 0 {
		log.Errorf("failed to delete orderType due to empty id param")
		response.WriteResponseWithError(w, apiErrors.ID_NOT_SPECIFIED_ERROR_CODE, fmt.Sprintf("Order type %s", apiErrors.ID_NOT_SPECIFIED_ERROR_MESSAGE), http.StatusBadRequest)
		return
	}

	id, err := utils.StringToUint(rawId)
	if err != nil {
		log.Errorf("failed to convert raw str id to int: %v", err)
		response.WriteResponseWithError(w, apiErrors.NOT_FOUND_ERROR_CODE, apiErrors.NOT_FOUND_ERROR_MESSAGE, http.StatusNotFound)
		return
	}

	orderType, err := h.OrderTypeRepo.GetByID(id)
	if err != nil {
		log.Errorf("failed to delete orderType: %v", err)
		response.WriteResponseWithError(w, apiErrors.NOT_FOUND_ERROR_CODE, "No order type found with the provided ID.", http.StatusNotFound)
		return
	}

	resp, err := response.NewCustomerResponseWithData("order_type", orderType)
	if err != nil {
		log.Errorf("Failed to create new response with customers: %v", err)
		return
	}

	response.WriteResponse(w, http.StatusOK, resp)
}

func (h *Handler) UpdateOrderType(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	orderType := &repository.OrderType{}

	if err := json.NewDecoder(r.Body).Decode(orderType); err != nil {
		log.Errorf("failed to decode body: %v", err)
		response.WriteResponseWithError(w, apiErrors.INTERNAL_ERROR_CODE, apiErrors.INTERNAL_ERROR_MESSAGE, http.StatusInternalServerError)
		return
	}

	if err := h.Validator.Validate(orderType); err != nil {
		log.Errorf("failed to validate orderType struct: %v", err)
		response.WriteResponseWithError(w, apiErrors.FIELD_VALIDATION_ERROR_CODE, apiErrors.FIELD_VALIDATION_ERROR_MESSAGE, http.StatusBadRequest)
		return
	}

	if err := h.OrderTypeRepo.Update(orderType); err != nil {
		log.Errorf("failed to create orderType: %v", err)
		response.WriteResponseWithError(w, apiErrors.INTERNAL_ERROR_CODE, apiErrors.INTERNAL_ERROR_MESSAGE, http.StatusInternalServerError)
		return
	}

	response.WriteResponseWithSuccess(w)
}
