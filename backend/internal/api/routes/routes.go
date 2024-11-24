package routes

import (
	"net/http"

	"github.com/TSI-Projects/group-project/internal/api/handlers"
	"github.com/TSI-Projects/group-project/internal/api/middleware"
	"github.com/TSI-Projects/group-project/internal/db"
	"github.com/gorilla/mux"
)

func NewRouter(dbClient db.IDatabase) *mux.Router {
	router := mux.NewRouter()
	handler := handlers.NewHandler(dbClient)

	router.Use(middleware.ContentTypeJsonMW)

	router.HandleFunc("/api/orders", handler.GetOrders).Methods(http.MethodGet)
	router.HandleFunc("/api/orders", handler.CreateOrder).Methods(http.MethodPost)
	router.HandleFunc("/api/orders", handler.UpdateOrder).Methods(http.MethodPut)
	router.HandleFunc("/api/order/{id}", handler.DeleteOrder).Methods(http.MethodDelete)
	router.HandleFunc("/api/order/{id}", handler.GetOrderByID).Methods(http.MethodGet)

	router.HandleFunc("/api/orders/types", handler.GetOrderTypes).Methods(http.MethodGet)
	router.HandleFunc("/api/orders/types", handler.CreateOrderType).Methods(http.MethodPost)
	router.HandleFunc("/api/orders/types", handler.UpdateOrderType).Methods(http.MethodPut)
	router.HandleFunc("/api/orders/type/{id}", handler.DeleteOrderType).Methods(http.MethodDelete)
	router.HandleFunc("/api/orders/type/{id}", handler.GetOrderTypeByID).Methods(http.MethodGet)

	router.HandleFunc("/api/workers", handler.GetWorkers).Methods(http.MethodGet)
	router.HandleFunc("/api/workers", handler.CreateWorker).Methods(http.MethodPost)
	router.HandleFunc("/api/workers", handler.UpdateWorker).Methods(http.MethodPut)
	router.HandleFunc("/api/worker/{id}", handler.DeleteWorker).Methods(http.MethodDelete)
	router.HandleFunc("/api/worker/{id}", handler.GetWorkerByID).Methods(http.MethodGet)

	router.HandleFunc("/api/languages", handler.GetLanguages).Methods(http.MethodGet)
	router.HandleFunc("/api/languages", handler.CreateLanguage).Methods(http.MethodPost)
	router.HandleFunc("/api/languages", handler.UpdateLanguage).Methods(http.MethodPut)
	router.HandleFunc("/api/language/{id}", handler.DeleteLanguage).Methods(http.MethodDelete)
	router.HandleFunc("/api/language/{id}", handler.GetLanguageByID).Methods(http.MethodGet)

	router.HandleFunc("/api/customers", handler.GetCustomers).Methods(http.MethodGet)
	router.HandleFunc("/api/customers", handler.CreateCustomer).Methods(http.MethodPost)
	router.HandleFunc("/api/customers", handler.UpdateCustomer).Methods(http.MethodPut)
	router.HandleFunc("/api/customer/{id}", handler.DeleteCustomer).Methods(http.MethodDelete)
	router.HandleFunc("/api/customer/{id}", handler.GetCustomerByID).Methods(http.MethodGet)

	router.HandleFunc("/", handler.GetOrders).Methods(http.MethodGet)
	return router
}
