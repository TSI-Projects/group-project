package handlers

import (
	"encoding/json"
	"net/http"

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
		w.Write([]byte("Internal Error"))
		return
	}

	if orders == nil {
		log.Warning("Orders is empty")
		w.Write([]byte("Orders is empty"))
		return
	}

	byteOrders, err := json.Marshal(orders)
	if err != nil {
		log.Errorf("Failed to marshal orders: %v", err)
		w.Write([]byte("Internal Error"))
		return
	}

	w.Write(byteOrders)
}

func (h *Handler) GetActiveOrders(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	log.Println("Active orders request received!")

	orders, err := h.OrderRepo.GetActiveOrders()
	if err != nil {
		log.Errorf("Failed to get orders: %v", err)
		w.Write([]byte("Internal Error"))
		return
	}

	if orders == nil {
		log.Warning("Orders is empty")
		w.Write([]byte("Orders is empty"))
		return
	}

	byteOrders, err := json.Marshal(orders)
	if err != nil {
		log.Errorf("Failed to marshal orders: %v", err)
		w.Write([]byte("Internal Error"))
		return
	}

	w.Write(byteOrders)
}

func (h *Handler) GetCompletedOrders(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	log.Println("Completed orders request received!")

	orders, err := h.OrderRepo.GetCompletedOrders()
	if err != nil {
		log.Errorf("Failed to get orders: %v", err)
		w.Write([]byte("Internal Error"))
		return
	}

	if orders == nil {
		log.Warning("Orders is empty")
		w.Write([]byte("Orders is empty"))
		return
	}

	byteOrders, err := json.Marshal(orders)
	if err != nil {
		log.Errorf("Failed to marshal orders: %v", err)
		w.Write([]byte("Internal Error"))
		return
	}

	w.Write(byteOrders)
}

func (h *Handler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var newOrder repository.Order

	log.Debugln("Create Order request received")

	if err := json.NewDecoder(r.Body).Decode(&newOrder); err != nil {
		log.Errorf("failed to decode body: %v", err)
		w.Write([]byte("Internal Error"))
		return
	}

	if err := h.Validator.Validate(newOrder); err != nil {
		log.Errorf("failed to validate order struct: %v", err)
		w.Write([]byte(utils.UppercaseFirstLetter(err.Error())))
		return
	}

	if err := h.Validator.Validate(newOrder); err != nil {
		log.Errorf("failed to validate order struct: %v", err)
		w.Write([]byte(utils.UppercaseFirstLetter(err.Error())))
		return
	}

	if err := h.OrderRepo.Create(&newOrder); err != nil {
		log.Errorf("failed to create order item: %v", err)
		w.Write([]byte("Internal Error"))
		return
	}

	w.Write([]byte("Success"))
}

func (h *Handler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	rawId := mux.Vars(r)["id"]
	if len(rawId) == 0 {
		log.Errorf("failed to delete order due to empty id param")
		w.Write([]byte("Order ID is not specified"))
		return
	}

	id, err := utils.StringToUint(rawId)
	if err != nil {
		log.Errorf("failed to convert id url param '%s': %v", rawId, err)
		w.Write([]byte("Internal Error"))
		return
	}

	if err := h.OrderRepo.Delete(id); err != nil {
		log.Errorf("failed to delete order with '%s' id: %v", rawId, err)
		w.Write([]byte("Internal Error"))
		return
	}

	w.Write([]byte("Success"))
}

func (h *Handler) GetOrderByID(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	rawId := mux.Vars(r)["id"]
	if len(rawId) == 0 {
		log.Errorf("failed to delete order due to empty id param")
		w.Write([]byte("Order ID is not specified"))
		return
	}

	id, err := utils.StringToUint(rawId)
	if err != nil {
		log.Errorf("failed to convert raw str id to int: %v", err)
		w.Write([]byte("Internal Error"))
		return
	}

	order, err := h.OrderRepo.GetByID(id)
	if err != nil {
		log.Errorf("failed to get order by id: %v", err)
		w.Write([]byte("Internal Error"))
		return
	}

	output, err := json.Marshal(order)
	if err != nil {
		log.Errorf("Failed to marshal orders struct: %v", err)
		w.Write([]byte("Internal Error"))
		return
	}

	w.Write(output)
}

func (h *Handler) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	order := &repository.Order{}

	if err := json.NewDecoder(r.Body).Decode(order); err != nil {
		log.Errorf("failed to decode body: %v", err)
		w.Write([]byte("Internal Error"))
		return
	}

	if err := h.Validator.Validate(order); err != nil {
		log.Errorf("failed to validate order struct: %v", err)
		w.Write([]byte(utils.UppercaseFirstLetter(err.Error())))
		return
	}

	if err := h.OrderRepo.Update(order); err != nil {
		log.Errorf("failed to update order: %v", err)
		w.Write([]byte("Internal Error"))
		return
	}

	if err := h.OrderStatusRepo.Update(order.Status); err != nil {
		log.Errorf("failed to update order status: %v", err)
		w.Write([]byte("Internal Error"))
		return
	}

	if err := h.CustomerRepo.Update(order.Customer); err != nil {
		log.Errorf("failed to update customer: %v", err)
		w.Write([]byte("Internal Error"))
		return
	}

	w.Write([]byte("Success"))
}
