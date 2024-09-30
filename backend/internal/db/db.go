package db

import (
	"database/sql"
	"fmt"

	"github.com/TSI-Projects/group-project/internal/config"
	_ "github.com/lib/pq"
)

var Database *sql.DB

func InitDB() error {
	var err error
	conf := config.NewPostgresConfig()

	if Database, err = sql.Open("postgres", conf.ConnectionStr); err != nil {
		return fmt.Errorf("failed to open connection: %v", err)
	}

	if err = Database.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %v", err)
	}

	return nil
}
