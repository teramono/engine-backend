package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/teramono/utilities/pkg/broker"
	"github.com/teramono/utilities/pkg/setup"
	"rogchap.com/v8go"
)

type BackendServer struct {
	setup.CommonSetup
	rootPath string
	dbs      []broker.Address
	port     uint
}

// NewBackendServer ...
func NewBackendServer(setup setup.CommonSetup, rootPath string, port uint, dbs []broker.Address) (BackendServer, error) {
	return BackendServer{
		CommonSetup: setup,
		rootPath: rootPath,
		dbs:      dbs,
		port:     port,
	}, nil
}

// Listen ...
func (server *BackendServer) Listen() error {
	router := gin.Default()

	router.POST("/login", server.Login)
	router.Any("/run/*all", server.Run)

	return router.Run(fmt.Sprintf(":%d", server.port))
}

func (server *BackendServer) runScript(script []byte) (*v8go.Value, error) {
	// TODO: Create custom capability context for each run.
	isolate, err := v8go.NewIsolate()
	if err != nil {
		return &v8go.Value{}, err
	}

	context, err := v8go.NewContext(isolate)
	if err != nil {
		return &v8go.Value{}, err
	}

	fmt.Println(">> script =", string(script))

	return context.RunScript(string(script), "")
}
