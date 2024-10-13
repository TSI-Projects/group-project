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

func (h *Handler) GetCustomers(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	customers, err := h.CustomerRepo.GetAll()
	if err != nil {
		log.Errorf("Failed to get customers: %v", err)
		w.Write([]byte("Internal Error"))
		return
	}

	output, err := json.Marshal(customers)
	if err != nil {
		log.Errorf("Failed to marshal customers struct: %v", err)
		w.Write([]byte("Internal Error"))
		return
	}

	w.Write(output)
}

func (h *Handler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	customer := &repository.Customer{}

	if err := json.NewDecoder(r.Body).Decode(customer); err != nil {
		log.Errorf("failed to decode body: %v", err)
		w.Write([]byte("Internal Error"))
		return
	}

	if err := h.Validator.Validate(customer); err != nil {
		log.Errorf("failed to validate customer struct: %v", err)
		w.Write([]byte(utils.UppercaseFirstLetter(err.Error())))
		return
	}

	if err := h.CustomerRepo.Create(customer); err != nil {
		log.Errorf("failed to create customer: %v", err)
		w.Write([]byte("Internal Error"))
		return
	}

	w.Write([]byte("Success"))
}

func (h *Handler) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	rawId := mux.Vars(r)["id"]
	if len(rawId) == 0 {
		log.Errorf("failed to delete customer due to empty id param")
		w.Write([]byte("Customer ID is not specified"))
		return
	}

	id, err := strconv.Atoi(rawId)
	if err != nil {
		log.Errorf("failed to convert raw str id to int: %v", err)
		w.Write([]byte("Internal Error"))
		return
	}

	if err := h.CustomerRepo.Delete(id); err != nil {
		log.Errorf("failed to delete customer: %v", err)
		w.Write([]byte("Internal Error"))
		return
	}

	w.Write([]byte("Success"))
}

func (h *Handler) GetCustomerByID(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	rawId := mux.Vars(r)["id"]
	if len(rawId) == 0 {
		log.Errorf("failed to delete customer due to empty id param")
		w.Write([]byte("Customer ID is not specified"))
		return
	}

	id, err := strconv.Atoi(rawId)
	if err != nil {
		log.Errorf("failed to convert raw str id to int: %v", err)
		w.Write([]byte("Internal Error"))
		return
	}

	customer, err := h.CustomerRepo.GetByID(id)
	if err != nil {
		log.Errorf("failed to delete customer: %v", err)
		w.Write([]byte("Internal Error"))
		return
	}

	output, err := json.Marshal(customer)
	if err != nil {
		log.Errorf("Failed to marshal customers struct: %v", err)
		w.Write([]byte("Internal Error"))
		return
	}

	w.Write(output)
}

func (h *Handler) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	customer := &repository.Customer{}

	if err := json.NewDecoder(r.Body).Decode(customer); err != nil {
		log.Errorf("failed to decode body: %v", err)
		w.Write([]byte("Internal Error"))
		return
	}

	if err := h.Validator.Validate(customer); err != nil {
		log.Errorf("failed to validate customer struct: %v", err)
		w.Write([]byte(utils.UppercaseFirstLetter(err.Error())))
		return
	}

	if err := h.CustomerRepo.Update(customer); err != nil {
		log.Errorf("failed to create customer: %v", err)
		w.Write([]byte("Internal Error"))
		return
	}

	w.Write([]byte("Success"))
}
