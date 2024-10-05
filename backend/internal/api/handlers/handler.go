package handlers

import (
	"github.com/TSI-Projects/group-project/internal/db"
)

type Handler struct {
	DBClient db.IDatabase
}

func NewHandler(databaseClient db.IDatabase) *Handler {
	return &Handler{
		DBClient: databaseClient,
	}
}
