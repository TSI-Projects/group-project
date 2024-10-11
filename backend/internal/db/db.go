package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DBClient struct {
	config   *PostgresConfig
	database *sql.DB
}

func NewDBClient() (IDatabase, error) {
	conf := NewPostgresConfig()

	db, err := sql.Open("postgres", conf.ConnectionStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open connection: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	return &DBClient{
		config:   conf,
		database: db,
	}, nil
}

func (c *DBClient) Close() {
	c.database.Close()
}

func (c *DBClient) Query(query string, args ...any) (*sql.Rows, error) {
	return c.database.Query(query, args)
}

func (c *DBClient) Exec(query string, args ...any) (sql.Result, error) {
	return c.database.Exec(query, args)
}

func (c *DBClient) Begin() (*sql.Tx, error) {
	return c.database.Begin()
}
