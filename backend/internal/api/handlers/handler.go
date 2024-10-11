package handlers

import (
	"github.com/TSI-Projects/group-project/internal/db"
	"github.com/TSI-Projects/group-project/internal/repository"
)

type Handler struct {
	DBClient      db.IDatabase
	OrderRepo     repository.IRepository[repository.Order]
	OrderTypeRepo repository.IRepository[repository.OrderType]
	WorkerRepo    repository.IRepository[repository.Worker]
	LanguageRepo  repository.IRepository[repository.Language]
}

func NewHandler(dbClient db.IDatabase) *Handler {
	return &Handler{
		DBClient:      dbClient,
		OrderRepo:     repository.NewOrderRepo(dbClient),
		OrderTypeRepo: repository.NewOrderTypeRepo(dbClient),
		WorkerRepo:    repository.NewWorkerRepo(dbClient),
		LanguageRepo:  repository.NewLanguageRepo(dbClient),
	}
}
