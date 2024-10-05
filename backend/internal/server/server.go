package server

import (
	"fmt"
	"net/http"

	"github.com/TSI-Projects/group-project/internal/api/routes"
	"github.com/TSI-Projects/group-project/internal/db"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

const SERVER_PORT = "8000"

type Server struct {
	Router   *mux.Router
	Database db.IDatabase

	ServerPort string
}

func NewServer(database db.IDatabase) IServer {
	router := routes.NewRouter(database)

	return &Server{
		ServerPort: SERVER_PORT,
		Router:     router,
		Database:   database,
	}
}

func (s *Server) Start() error {
	log.Printf("Server is running on port %s", s.ServerPort)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", s.ServerPort), s.Router); err != nil {
		return fmt.Errorf("failed to start server: %v", err)
	}
	return nil
}
