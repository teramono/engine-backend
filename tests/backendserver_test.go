package tests

import (
	"testing"

	"github.com/teramono/engine-backend/pkg/engine"
	"github.com/teramono/utilities/pkg/setup"
)

func TestBackendServer(t *testing.T) {
	setup, err := setup.NewSetup(false)
	if err != nil {
		// ...
	}

	server, err := engine.NewBackendServer(setup)
	if err != nil {
		// ...
	}

	server.Listen()
}
