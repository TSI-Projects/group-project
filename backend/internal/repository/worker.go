package repository

import (
	"fmt"

	"github.com/TSI-Projects/group-project/internal/db"
)

type Worker struct {
	ID        int    `db:"id"          json:"id"`
	FirstName string `db:"first_name"  json:"first_name"`
	LastName  string `db:"last_name"   json:"last_name"`
}

type WorkerRepo struct {
	DBClient db.IDatabase
}

type IWorkerRepo interface {
	CreateWorker(*Worker) error
	GetWorkers() ([]*Worker, error)
	DeleteWorker(int) error
}

func NewWorkerRepo(database db.IDatabase) IWorkerRepo {
	return &WorkerRepo{
		DBClient: database,
	}
}

func (w *WorkerRepo) CreateWorker(worker *Worker) error {
	db := w.DBClient.GetConn()

	if _, err := db.Query(
		`INSERT INTO workers
			(first_name, last_name)
		VALUES
			($1, $2)`,
		worker.FirstName,
		worker.LastName,
	); err != nil {
		return fmt.Errorf("failed to make insert query request: %v", err)
	}

	return nil
}

func (w *WorkerRepo) DeleteWorker(id int) error {
	db := w.DBClient.GetConn()

	if _, err := db.Query(`DELETE FROM workers WHERE id = $1`, id); err != nil {
		return fmt.Errorf("failed to make delete query request: %v", err)
	}

	return nil
}

func (w *WorkerRepo) GetWorkers() ([]*Worker, error) {
	db := w.DBClient.GetConn()
	var workers []*Worker

	rows, err := db.Query(
		`SELECT
			id,
			first_name,
			last_name
		FROM
			workers`,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to make get query request: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		worker := &Worker{}

		if err := rows.Scan(
			&worker.ID,
			&worker.FirstName,
			&worker.LastName,
		); err != nil {
			return nil, fmt.Errorf("failed to scan rows in workers table: %v", err)
		}
		workers = append(workers, worker)
	}

	return workers, nil
}
