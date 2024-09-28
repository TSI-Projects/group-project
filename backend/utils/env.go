package utils

import (
	"log"
	"os"
)

func LookupEnv(envVar string) string {
	value, ok := os.LookupEnv(envVar)
	if !ok {
		log.Fatalf("Environment variable '%s' is missing. Please ensure it is set before starting server", envVar)
	}
	return value
}
