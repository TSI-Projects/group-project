package main

import (
	"fmt"
	"net/http"

	"github.com/TSI-Projects/group-project/internal/config"
	log "github.com/sirupsen/logrus"
)

func main() {
	config := config.NewAppConfig()

	if config.Port == "" {
		log.Errorln("Server port is not defined")
		return
	}

	log.Printf("Server is running on port %s", config.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", config.Port), config.Router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
