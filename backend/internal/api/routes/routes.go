package routes

import (
	"net/http"

	"github.com/TSI-Projects/group-project/internal/api/handlers"
	"github.com/TSI-Projects/group-project/internal/db"
	"github.com/gorilla/mux"
)

func NewRouter(dbClient db.IDatabase) *mux.Router {
	router := mux.NewRouter()
	handler := handlers.NewHandler(dbClient)

	router.HandleFunc("/api/order", handler.GetOrder).Methods(http.MethodGet)
	router.HandleFunc("/", handler.GetOrder).Methods(http.MethodGet)
	return router
}
