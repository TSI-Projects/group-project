package handlers

import (
	"github.com/TSI-Projects/group-project/internal/auth"
	"github.com/TSI-Projects/group-project/internal/db"
	"github.com/TSI-Projects/group-project/internal/repository"
	"github.com/TSI-Projects/group-project/pkg/validation"
)

type Handler struct {
	DBClient        db.IDatabase
	OrderRepo       repository.IOrderRepository[repository.Order]
	OrderTypeRepo   repository.IRepository[repository.OrderType]
	WorkerRepo      repository.IRepository[repository.Worker]
	LanguageRepo    repository.IRepository[repository.Language]
	CustomerRepo    repository.IRepository[repository.Customer]
	OrderStatusRepo repository.IRepository[repository.OrderStatus]
	AuthClient      auth.IAuthClient
	Validator       *validation.ValidatorClient
}

func NewHandler(dbClient db.IDatabase) *Handler {
	return &Handler{
		DBClient:        dbClient,
		OrderRepo:       repository.NewOrderRepo(dbClient),
		OrderTypeRepo:   repository.NewOrderTypeRepo(dbClient),
		WorkerRepo:      repository.NewWorkerRepo(dbClient),
		LanguageRepo:    repository.NewLanguageRepo(dbClient),
		CustomerRepo:    repository.NewCustomerRepo(dbClient),
		OrderStatusRepo: repository.NewOrderStatusRepo(dbClient),
		Validator:       validation.NewValidatorClient(),
		AuthClient:      auth.NewAuthClient(dbClient),
	}
}
