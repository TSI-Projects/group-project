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

func (h *Handler) GetOrderTypes(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	orderTypes, err := h.OrderTypeRepo.GetOrderTypes()
	if err != nil {
		log.Errorf("failed to get order types: %v", err)
		w.Write([]byte("Internal Error"))
		return
	}

	output, err := json.Marshal(orderTypes)
	if err != nil {
		log.Errorf("failed to marshal order types: %v", err)
		w.Write([]byte("Internal Error"))
		return
	}

	w.Write(output)
}

func (h *Handler) CreateOrderType(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var orderType repository.OrderType

	if err := json.NewDecoder(r.Body).Decode(&orderType); err != nil {
		log.Errorf("failed to decode body: %v", err)
		w.Write([]byte("Internal Error"))
		return
	}

	if err := h.Validator.Validate(orderType); err != nil {
		log.Errorf("failed to validate order type struct: %v", err)
		w.Write([]byte(utils.UppercaseFirstLetter(err.Error())))
		return
	}

	if err := h.OrderTypeRepo.CreateOrderType(&orderType); err != nil {
		log.Errorf("failed to create order type: %v", err)
		w.Write([]byte("Internal Error"))
		return
	}

	w.Write([]byte("Success"))
}

func (h *Handler) DeleteOrderType(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	rawId := mux.Vars(r)["id"]
	if len(rawId) == 0 {
		log.Errorf("failed to delete order type due to empty id param")
		w.Write([]byte("Order type ID is not specified"))
		return
	}

	id, err := strconv.Atoi(rawId)
	if err != nil {
		log.Errorf("failed to convert str id to int: id '%s': %v ", rawId, err)
		w.Write([]byte("Internal Error"))
		return
	}

	if err := h.OrderTypeRepo.DeleteOrderType(id); err != nil {
		log.Errorf("failed to delete order type item: %s", err)
		w.Write([]byte("Internal Error"))
		return
	}

	w.Write([]byte("Success"))
}
