package repository

import (
	"fmt"

	"github.com/TSI-Projects/group-project/internal/db"
)

type Admin struct {
	ID       int    `db:"id"          json:"id"`
	Username string `db:"username"    json:"username"`
	Password string `db:"password"    json:"password"`
}

type AdminRepo struct {
	DBClient db.IDatabase
}

type IAdminRepository interface {
	GetByUsername(username string) (*Admin, error)
}

func NewAdminRepo(dbClient db.IDatabase) IAdminRepository {
	return &AdminRepo{
		DBClient: dbClient,
	}
}

func (a *AdminRepo) GetByUsername(username string) (*Admin, error) {
	admin := &Admin{
		Username: username,
	}

	if err := a.DBClient.QueryRow(
		`SELECT password, id
		 FROM admins
		 WHERE username = $1
		`, username,
	).Scan(
		&admin.Password,
		&admin.ID,
	); err != nil {
		return admin, fmt.Errorf("failed to execute query row: %v", err)
	}

	return admin, nil
}
