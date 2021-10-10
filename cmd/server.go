package main

import (
	"log"

	"github.com/teramono/engine-backend/pkg/server"
	"github.com/teramono/utilities/pkg/setup"
)

func main() {
	// Establish common setup.
	setup, err := setup.NewCommonSetup()
	if err != nil {
		log.Panic(err)
	}

	// Create server.
	server, err := server.NewBackendServer(setup)
	if err != nil {
		log.Panic(err)
	}

	defer server.BrokerClient.Close()

	// Activate subscriptions.
	err = server.ActivateSubscriptions()
	if err != nil {
		log.Panic(err)
	}
}
