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

	router.HandleFunc("/api/orders", middleware.AuthMW(handler.GetOrders)).Methods(http.MethodGet)
	router.HandleFunc("/api/orders", middleware.AuthMW(handler.CreateOrder)).Methods(http.MethodPost)
	router.HandleFunc("/api/orders", middleware.AuthMW(handler.UpdateOrder)).Methods(http.MethodPut)
	router.HandleFunc("/api/order/{id}", middleware.AuthMW(handler.DeleteOrder)).Methods(http.MethodDelete)
	router.HandleFunc("/api/order/{id}", middleware.AuthMW(handler.GetOrderByID)).Methods(http.MethodGet)

	router.HandleFunc("/api/orders/active", middleware.AuthMW(handler.GetActiveOrders)).Methods(http.MethodGet)
	router.HandleFunc("/api/orders/completed", middleware.AuthMW(handler.GetCompletedOrders)).Methods(http.MethodGet)

	router.HandleFunc("/api/orders/types", middleware.AuthMW(handler.GetOrderTypes)).Methods(http.MethodGet)
	router.HandleFunc("/api/orders/types", middleware.AuthMW(handler.CreateOrderType)).Methods(http.MethodPost)
	router.HandleFunc("/api/orders/types", middleware.AuthMW(handler.UpdateOrderType)).Methods(http.MethodPut)
	router.HandleFunc("/api/orders/type/{id}", middleware.AuthMW(handler.DeleteOrderType)).Methods(http.MethodDelete)
	router.HandleFunc("/api/orders/type/{id}", middleware.AuthMW(handler.GetOrderTypeByID)).Methods(http.MethodGet)

	router.HandleFunc("/api/workers", middleware.AuthMW(handler.GetWorkers)).Methods(http.MethodGet)
	router.HandleFunc("/api/workers", middleware.AuthMW(handler.CreateWorker)).Methods(http.MethodPost)
	router.HandleFunc("/api/workers", middleware.AuthMW(handler.UpdateWorker)).Methods(http.MethodPut)
	router.HandleFunc("/api/worker/{id}", middleware.AuthMW(handler.DeleteWorker)).Methods(http.MethodDelete)
	router.HandleFunc("/api/worker/{id}", middleware.AuthMW(handler.GetWorkerByID)).Methods(http.MethodGet)

	router.HandleFunc("/api/languages", middleware.AuthMW(handler.GetLanguages)).Methods(http.MethodGet)
	router.HandleFunc("/api/languages", middleware.AuthMW(handler.CreateLanguage)).Methods(http.MethodPost)
	router.HandleFunc("/api/languages", middleware.AuthMW(handler.UpdateLanguage)).Methods(http.MethodPut)
	router.HandleFunc("/api/language/{id}", middleware.AuthMW(handler.DeleteLanguage)).Methods(http.MethodDelete)
	router.HandleFunc("/api/language/{id}", middleware.AuthMW(handler.GetLanguageByID)).Methods(http.MethodGet)

	router.HandleFunc("/api/customers", middleware.AuthMW(handler.GetCustomers)).Methods(http.MethodGet)
	router.HandleFunc("/api/customers", middleware.AuthMW(handler.CreateCustomer)).Methods(http.MethodPost)
	router.HandleFunc("/api/customers", middleware.AuthMW(handler.UpdateCustomer)).Methods(http.MethodPut)
	router.HandleFunc("/api/customer/{id}", middleware.AuthMW(handler.DeleteCustomer)).Methods(http.MethodDelete)
	router.HandleFunc("/api/customer/{id}", middleware.AuthMW(handler.GetCustomerByID)).Methods(http.MethodGet)

	router.HandleFunc("/api/login", handler.Login).Methods(http.MethodPost)
	router.HandleFunc("/ping", handler.Pong).Methods(http.MethodGet)

	return router
}
