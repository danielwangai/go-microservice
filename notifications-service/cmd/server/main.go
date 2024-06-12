package main

import (
	"log"

	"github.com/danielwangai/twiga-foods/notifications-service/internal/transport"
)

func main() {
	err := transport.RunServer()
	if err != nil {
		log.Fatalf("Could not initialize server: %v", err)
	}
}
