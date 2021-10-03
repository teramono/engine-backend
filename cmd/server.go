package main

import (
	"log"

	"github.com/teramono/engine-backend/pkg/server"
	"github.com/teramono/utilities/pkg/broker"
	"github.com/teramono/utilities/pkg/setup"
)

// TODO:
const Port = 5051
const RootPath = "../tmp/"

func main() {
	// ...
	setup, err := setup.NewCommonSetup()
	if err != nil {
		log.Fatalln(err)
	}

	// ...
	server, err := server.NewBackendServer(setup, RootPath, Port, []broker.Address{})
	if err != nil {
		log.Fatalln(err)
	}

	server.Listen()
}
