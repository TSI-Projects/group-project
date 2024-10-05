package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DBClient struct {
	config     *PostgresConfig
	connection *sql.DB
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
		config:     conf,
		connection: db,
	}, nil
}

func (c *DBClient) Close() {
	c.connection.Close()
}
