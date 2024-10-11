package db

import "database/sql"

type IDatabase interface {
	Close()
	Begin() (*sql.Tx, error)
	Query(query string, args ...any) (*sql.Rows, error)
	Exec(query string, args ...any) (sql.Result, error)
}
