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

func (h *Handler) GetOrders(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	log.Println("Order request received!")

	orders, err := h.OrderRepo.GetAll()
	if err != nil {
		log.Errorf("Failed to get orders: %v", err)
		response.WriteResponseWithError(w, apiErrors.INTERNAL_ERROR_CODE, apiErrors.DEFAULT_ERROR_MESSAGE, http.StatusInternalServerError)
		return
	}

	resp, err := response.NewCustomerResponseWithDataArr("orders", orders)
	if err != nil {
		log.Errorf("Failed to create new response with customers: %v", err)
		return
	}

	response.WriteResponse(w, http.StatusOK, resp)
}

func (h *Handler) GetActiveOrders(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	log.Println("Active orders request received!")

	orders, err := h.OrderRepo.GetActiveOrders()
	if err != nil {
		log.Errorf("Failed to get orders: %v", err)
		response.WriteResponseWithError(w, apiErrors.INTERNAL_ERROR_CODE, apiErrors.DEFAULT_ERROR_MESSAGE, http.StatusInternalServerError)
		return
	}

	resp, err := response.NewCustomerResponseWithDataArr("orders", orders)
	if err != nil {
		log.Errorf("Failed to create new response with customers: %v", err)
		return
	}

	response.WriteResponse(w, http.StatusOK, resp)
}

func (h *Handler) GetCompletedOrders(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	log.Println("Completed orders request received!")

	orders, err := h.OrderRepo.GetCompletedOrders()
	if err != nil {
		log.Errorf("Failed to get orders: %v", err)
		response.WriteResponseWithError(w, apiErrors.INTERNAL_ERROR_CODE, apiErrors.DEFAULT_ERROR_MESSAGE, http.StatusInternalServerError)
		return
	}

	resp, err := response.NewCustomerResponseWithDataArr("orders", orders)
	if err != nil {
		log.Errorf("Failed to create new response with customers: %v", err)
		return
	}

	response.WriteResponse(w, http.StatusOK, resp)
}

func (h *Handler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var newOrder repository.Order

	log.Debugln("Create Order request received")

	if err := json.NewDecoder(r.Body).Decode(&newOrder); err != nil {
		log.Errorf("failed to decode body: %v", err)
		response.WriteResponseWithError(w, apiErrors.INTERNAL_ERROR_CODE, apiErrors.DEFAULT_ERROR_MESSAGE, http.StatusInternalServerError)
		return
	}

	if err := h.Validator.Validate(newOrder); err != nil {
		log.Errorf("failed to validate order struct: %v", err)
		response.WriteResponseWithError(w, apiErrors.FIELD_VALIDATION_ERROR_CODE, apiErrors.FIELD_VALIDATION_ERROR_MESSAGE, http.StatusBadRequest)
		return
	}

	if err := h.OrderRepo.Create(&newOrder); err != nil {
		log.Errorf("failed to create order item: %v", err)
		response.WriteResponseWithError(w, apiErrors.INTERNAL_ERROR_CODE, apiErrors.DEFAULT_ERROR_MESSAGE, http.StatusInternalServerError)
		return
	}

	response.WriteResponseWithSuccess(w)
}

func (h *Handler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	rawId := mux.Vars(r)["id"]
	if len(rawId) == 0 {
		log.Errorf("failed to delete order due to empty id param")
		response.WriteResponseWithError(w, apiErrors.ID_NOT_SPECIFIED_ERROR_CODE, fmt.Sprintf("Order %s", apiErrors.ID_NOT_SPECIFIED_ERROR_MESSAGE), http.StatusBadRequest)
		return
	}

	id, err := utils.StringToUint(rawId)
	if err != nil {
		log.Errorf("failed to convert id url param '%s': %v", rawId, err)
		response.WriteResponseWithError(w, apiErrors.NOT_FOUND_ERROR_CODE, apiErrors.NOT_FOUND_ERROR_MESSAGE, http.StatusNotFound)
		return
	}

	if err := h.OrderRepo.Delete(id); err != nil {
		log.Errorf("failed to delete order with '%s' id: %v", rawId, err)
		response.WriteResponseWithError(w, apiErrors.INTERNAL_ERROR_CODE, apiErrors.DEFAULT_ERROR_MESSAGE, http.StatusInternalServerError)
		return
	}

	response.WriteResponseWithSuccess(w)
}

func (h *Handler) GetOrderByID(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	rawId := mux.Vars(r)["id"]
	if len(rawId) == 0 {
		log.Errorf("failed to delete order due to empty id param")
		response.WriteResponseWithError(w, apiErrors.ID_NOT_SPECIFIED_ERROR_CODE, fmt.Sprintf("Order %s", apiErrors.ID_NOT_SPECIFIED_ERROR_MESSAGE), http.StatusBadRequest)
		return
	}

	id, err := utils.StringToUint(rawId)
	if err != nil {
		log.Errorf("failed to convert raw str id to int: %v", err)
		response.WriteResponseWithError(w, apiErrors.NOT_FOUND_ERROR_CODE, apiErrors.NOT_FOUND_ERROR_MESSAGE, http.StatusNotFound)
		return
	}

	order, err := h.OrderRepo.GetByID(id)
	if err != nil {
		log.Errorf("failed to get order by id: %v", err)
		response.WriteResponseWithError(w, apiErrors.NOT_FOUND_ERROR_CODE, "No order found with the provided ID.", http.StatusNotFound)
		return
	}

	resp, err := response.NewCustomerResponseWithData("order", order)
	if err != nil {
		log.Errorf("Failed to create new response with customers: %v", err)
		return
	}

	response.WriteResponse(w, http.StatusOK, resp)
}

func (h *Handler) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	order := &repository.Order{}

	if err := json.NewDecoder(r.Body).Decode(order); err != nil {
		log.Errorf("failed to decode body: %v", err)
		response.WriteResponseWithError(w, apiErrors.INTERNAL_ERROR_CODE, apiErrors.INTERNAL_ERROR_MESSAGE, http.StatusInternalServerError)
		return
	}

	if err := h.Validator.Validate(order); err != nil {
		log.Errorf("failed to validate order struct: %v", err)
		response.WriteResponseWithError(w, apiErrors.FIELD_VALIDATION_ERROR_CODE, apiErrors.FIELD_VALIDATION_ERROR_MESSAGE, http.StatusBadRequest)
		return
	}

	if err := h.OrderRepo.Update(order); err != nil {
		log.Errorf("failed to update order: %v", err)
		response.WriteResponseWithError(w, apiErrors.INTERNAL_ERROR_CODE, apiErrors.INTERNAL_ERROR_MESSAGE, http.StatusInternalServerError)
		return
	}

	if err := h.OrderStatusRepo.Update(order.Status); err != nil {
		log.Errorf("failed to update order status: %v", err)
		response.WriteResponseWithError(w, apiErrors.INTERNAL_ERROR_CODE, apiErrors.INTERNAL_ERROR_MESSAGE, http.StatusInternalServerError)
		return
	}

	if err := h.CustomerRepo.Update(order.Customer); err != nil {
		log.Errorf("failed to update customer: %v", err)
		response.WriteResponseWithError(w, apiErrors.INTERNAL_ERROR_CODE, apiErrors.INTERNAL_ERROR_MESSAGE, http.StatusInternalServerError)
		return
	}

	response.WriteResponseWithSuccess(w)
}
