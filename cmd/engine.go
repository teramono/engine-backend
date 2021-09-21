package main

import (
	"log"

	"github.com/teramono/engine-backend/pkg/engine"
	"github.com/teramono/utilities/pkg/setup"
)

func main() {
	// ...
	setup, err := setup.NewSetup()
	if err != nil {
		log.Fatalln(err)
	}

	// ...
	engine, err := engine.NewBackendEngine(setup)
	if err != nil {
		log.Fatalln(err)
	}

	engine.Listen()
}
