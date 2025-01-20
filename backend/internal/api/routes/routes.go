package routes

import (
	"net/http"

	"github.com/TSI-Projects/group-project/internal/api/handlers"
	"github.com/TSI-Projects/group-project/internal/api/middleware"
	token "github.com/TSI-Projects/group-project/internal/auth/access_token"
	"github.com/TSI-Projects/group-project/internal/db"
	"github.com/gorilla/mux"
)

func NewRouter(dbClient db.IDatabase) *mux.Router {
	router := mux.NewRouter()
	tokenClient := token.NewTokenClient()
	handler := handlers.NewHandler(dbClient, tokenClient)

	router.Use(middleware.ContentTypeJsonMW)

	router.HandleFunc("/api/orders", middleware.AuthMW(handler.GetOrders, tokenClient)).Methods(http.MethodGet)
	router.HandleFunc("/api/orders", middleware.AuthMW(handler.CreateOrder, tokenClient)).Methods(http.MethodPost)
	router.HandleFunc("/api/orders", middleware.AuthMW(handler.UpdateOrder, tokenClient)).Methods(http.MethodPut)
	router.HandleFunc("/api/order/{id}", middleware.AuthMW(handler.DeleteOrder, tokenClient)).Methods(http.MethodDelete)
	router.HandleFunc("/api/order/{id}", middleware.AuthMW(handler.GetOrderByID, tokenClient)).Methods(http.MethodGet)

	router.HandleFunc("/api/orders/active", middleware.AuthMW(handler.GetActiveOrders, tokenClient)).Methods(http.MethodGet)
	router.HandleFunc("/api/orders/completed", middleware.AuthMW(handler.GetCompletedOrders, tokenClient)).Methods(http.MethodGet)

	router.HandleFunc("/api/orders/types", middleware.AuthMW(handler.GetOrderTypes, tokenClient)).Methods(http.MethodGet)
	router.HandleFunc("/api/orders/types", middleware.AuthMW(handler.CreateOrderType, tokenClient)).Methods(http.MethodPost)
	router.HandleFunc("/api/orders/types", middleware.AuthMW(handler.UpdateOrderType, tokenClient)).Methods(http.MethodPut)
	router.HandleFunc("/api/orders/type/{id}", middleware.AuthMW(handler.DeleteOrderType, tokenClient)).Methods(http.MethodDelete)
	router.HandleFunc("/api/orders/type/{id}", middleware.AuthMW(handler.GetOrderTypeByID, tokenClient)).Methods(http.MethodGet)

	router.HandleFunc("/api/workers", middleware.AuthMW(handler.GetWorkers, tokenClient)).Methods(http.MethodGet)
	router.HandleFunc("/api/workers", middleware.AuthMW(handler.CreateWorker, tokenClient)).Methods(http.MethodPost)
	router.HandleFunc("/api/workers", middleware.AuthMW(handler.UpdateWorker, tokenClient)).Methods(http.MethodPut)
	router.HandleFunc("/api/worker/{id}", middleware.AuthMW(handler.DeleteWorker, tokenClient)).Methods(http.MethodDelete)
	router.HandleFunc("/api/worker/{id}", middleware.AuthMW(handler.GetWorkerByID, tokenClient)).Methods(http.MethodGet)

	router.HandleFunc("/api/languages", middleware.AuthMW(handler.GetLanguages, tokenClient)).Methods(http.MethodGet)
	router.HandleFunc("/api/languages", middleware.AuthMW(handler.CreateLanguage, tokenClient)).Methods(http.MethodPost)
	router.HandleFunc("/api/languages", middleware.AuthMW(handler.UpdateLanguage, tokenClient)).Methods(http.MethodPut)
	router.HandleFunc("/api/language/{id}", middleware.AuthMW(handler.DeleteLanguage, tokenClient)).Methods(http.MethodDelete)
	router.HandleFunc("/api/language/{id}", middleware.AuthMW(handler.GetLanguageByID, tokenClient)).Methods(http.MethodGet)

	router.HandleFunc("/api/customers", middleware.AuthMW(handler.GetCustomers, tokenClient)).Methods(http.MethodGet)
	router.HandleFunc("/api/customers", middleware.AuthMW(handler.CreateCustomer, tokenClient)).Methods(http.MethodPost)
	router.HandleFunc("/api/customers", middleware.AuthMW(handler.UpdateCustomer, tokenClient)).Methods(http.MethodPut)
	router.HandleFunc("/api/customer/{id}", middleware.AuthMW(handler.DeleteCustomer, tokenClient)).Methods(http.MethodDelete)
	router.HandleFunc("/api/customer/{id}", middleware.AuthMW(handler.GetCustomerByID, tokenClient)).Methods(http.MethodGet)

	router.HandleFunc("/api/login", handler.Login).Methods(http.MethodPost)
	router.HandleFunc("/ping", handler.Pong).Methods(http.MethodGet)

	return router
}
