package api

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func GetOrder(w http.ResponseWriter, r *http.Request) {
	log.Println("Order request received!")
	w.Write([]byte("Hello World!"))
}
