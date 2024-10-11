package handlers

import (
	"github.com/TSI-Projects/group-project/internal/db"
	"github.com/TSI-Projects/group-project/internal/repository"
	"github.com/TSI-Projects/group-project/pkg/validation"
)

type Handler struct {
	DBClient      db.IDatabase
	OrderRepo     repository.IOrderRepo
	OrderTypeRepo repository.IOrderTypeRepo
	WorkerRepo    repository.IWorkerRepo
	LanguageRepo  repository.ILanguageRepo
	Validator     *validation.ValidatorClient
}

func NewHandler(dbClient db.IDatabase) *Handler {
	return &Handler{
		DBClient:      dbClient,
		OrderRepo:     repository.NewOrderRepo(dbClient),
		OrderTypeRepo: repository.NewOrderTypeRepo(dbClient),
		WorkerRepo:    repository.NewWorkerRepo(dbClient),
		LanguageRepo:  repository.NewLanguageRepo(dbClient),
		Validator:     validation.NewValidatorClient(),
	}
}
