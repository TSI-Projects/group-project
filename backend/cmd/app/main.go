package main

import (
	"fmt"
	"net/http"

	"github.com/TSI-Projects/group-project/internal/config"
	"github.com/TSI-Projects/group-project/internal/db"
	log "github.com/sirupsen/logrus"
)

func main() {
	if err := db.InitDB(); err != nil {
		log.Fatalf("Failed to init database: %v", err)
	}
	defer db.Database.Close()

	serverConfig := config.NewServerConfig()

	log.Printf("Server is running on port %s", serverConfig.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", serverConfig.Port), serverConfig.Router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
