package main

import (
	"log"

	"github.com/teramono/engine-backend/pkg/engine"
	"github.com/teramono/utilities/pkg/setup"
)

func main() {
	// ...
	setup, err := setup.NewSetup(true)
	if err != nil {
		log.Fatalln(err)
	}

	// ...
	server, err := engine.NewBackendServer(setup)
	if err != nil {
		log.Fatalln(err)
	}

	server.Listen()
}
