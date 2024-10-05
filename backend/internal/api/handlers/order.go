package handlers

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func (h *Handler) GetOrder(w http.ResponseWriter, r *http.Request) {
	log.Println("Order request received!")
	w.Write([]byte("Hello World!"))
}
