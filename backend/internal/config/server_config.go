package config

import (
	"github.com/TSI-Projects/group-project/internal/api"
	"github.com/TSI-Projects/group-project/utils"
	"github.com/gorilla/mux"
)

const (
	ENV_PORT_NAME = "PORT"
)

type ServerConfig struct {
	Port   string
	Router *mux.Router
}

func NewServerConfig() ServerConfig {
	return ServerConfig{
		Port:   utils.LookupEnv(ENV_PORT_NAME),
		Router: api.NewRouter(),
	}
}
