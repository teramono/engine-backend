// Package engine ...
package engine

import (
	"github.com/gin-gonic/gin"
	dbEngine "github.com/teramono/engine-db/pkg/engine"
	fsEngine "github.com/teramono/engine-fs/pkg/engine"
	"github.com/teramono/utilities/pkg/setup"
)

type BackendServer struct {
	isolates []Isolate
	setup    setup.Setup
	fs       fsEngine.FSInterface
	db       dbEngine.DBInterface
}

// NewBackendServer ...
func NewBackendServer(setup setup.Setup) (BackendServer, error) {
	isolate, err := NewIsolate()
	fs := fsEngine.NewFSInterface(&setup)
	db := dbEngine.NewDBInterface(&setup)

	if err != nil {
		return BackendServer{}, err
	}

	return BackendServer{
		isolates: []Isolate{isolate},
		setup:    setup,
		fs:       fs,
		db:       db,
	}, nil
}

// Listen ...
func (server *BackendServer) Listen() error {
	router := gin.Default()

	router.POST("/login", server.Login)
	router.Any("/run/*all", server.Run)

	return router.Run(":5051") // TODO: Get from setup.Config
}
