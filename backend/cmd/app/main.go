package main

import (
	"github.com/TSI-Projects/group-project/internal/db"
	"github.com/TSI-Projects/group-project/internal/server"
	log "github.com/sirupsen/logrus"
)

func main() {
	dbClient, err := db.NewDBClient()
	if err != nil {
		log.Fatalf("Failed to init database: %v", err)
	}
	defer dbClient.Close()

	server := server.NewServer(dbClient)
	if err := server.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
