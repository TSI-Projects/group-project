package config

import (
	"fmt"

	"github.com/TSI-Projects/group-project/utils"
	log "github.com/sirupsen/logrus"
)

const (
	ENV_DB_NAME     = "DB_DATABASE"
	ENV_DB_PASSWORD = "DB_PASSWORD"
	ENV_DB_USERNAME = "DB_USER"
	ENV_DB_PORT     = "DB_PORT"
	ENV_DB_HOST     = "DB_HOST"
)

type PostgresConfig struct {
	DBName        string
	Username      string
	Password      string
	Host          string
	Port          string
	ConnectionStr string
}

func NewPostgresConfig() *PostgresConfig {
	port := utils.LookupEnv(ENV_DB_PORT)
	host := utils.LookupEnv(ENV_DB_HOST)
	name := utils.LookupEnv(ENV_DB_NAME)
	pass := utils.LookupEnv(ENV_DB_PASSWORD)
	username := utils.LookupEnv(ENV_DB_USERNAME)

	return &PostgresConfig{
		Username:      username,
		Password:      pass,
		DBName:        name,
		Host:          host,
		Port:          port,
		ConnectionStr: fmt.Sprintf("user=%s dbname=%s password=%s port=%s host=%s sslmode=disable", username, name, pass, port, host),
	}
}

func (c *PostgresConfig) Debug() {
	log.Debugf("Database name: %s", c.DBName)
	log.Debugf("Username: %s", c.Username)
	log.Debugf("Password: %s", c.Password)
	log.Debugf("Host: %s", c.Host)
	log.Debugf("Port: %s", c.Port)
}
