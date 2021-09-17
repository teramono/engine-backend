package main

import (
	"log"

	"github.com/teramono/engine-backend/pkg/engine"
)

func main() {
	engine, err := engine.NewBackendEngine()
	if err != nil {
		log.Fatalln(err)
	}

	engine.Listen()
}
