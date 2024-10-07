package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/TSI-Projects/group-project/internal/repository"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func (h *Handler) GetOrders(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	log.Println("Order request received!")

	orders, err := h.OrderRepo.GetOrders()
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

	if err := h.OrderRepo.CreateOrder(&newOrder); err != nil {
		log.Errorf("failed to create order item: %v", err)
		w.Write([]byte("Internal Error"))
		return
	}

	w.Write([]byte("Success"))
}

func (h *Handler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	rawId := mux.Vars(r)["id"]
	id, err := strconv.Atoi(rawId)
	if err != nil {
		log.Errorf("failed to convert id url param '%s': %v", rawId, err)
		w.Write([]byte("Internal Error"))
		return
	}

	if err := h.OrderRepo.DeleteOrder(id); err != nil {
		log.Errorf("failed to delete order with '%s' id: %v", rawId, err)
		w.Write([]byte("Internal Error"))
		return
	}

	w.Write([]byte("Success"))
}
