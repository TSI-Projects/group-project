package config

import (
	"os"

	"github.com/TSI-Projects/group-project/internal/api"
	"github.com/gorilla/mux"
)

const (
	ENV_PORT_NAME = "PORT"
)

type ServerConfig struct {
	Port   string
	Router *mux.Router
}

func NewAppConfig() ServerConfig {
	return ServerConfig{
		Port:   os.Getenv(ENV_PORT_NAME),
		Router: api.NewRouter(),
	}
}
