package db

import "database/sql"

type IDatabase interface {
	Close()
	GetConn() *sql.DB
}
