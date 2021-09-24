// Package engine ...
package engine

import (
	"github.com/gin-gonic/gin"
	"github.com/teramono/utilities/pkg/setup"
)

type BackendServer struct {
	setup    setup.Setup
	isolates []Isolate
}

// NewBackendServer ...
func NewBackendServer(setup setup.Setup) (BackendServer, error) {
	isolate, err := NewIsolate()
	if err != nil {
		return BackendServer{}, err
	}

	return BackendServer{
		setup:    setup,
		isolates: []Isolate{isolate},
	}, nil
}

// Listen ...
func (server *BackendServer) Listen() error {
	router := gin.Default()

	// Use GRPC instead.
	router.POST("/", server.RunScriptURL)

	return router.Run()
}
