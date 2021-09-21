package engine

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/teramono/utilities/pkg/setup"
)

// BackendEngine ...
type BackendEngine struct {
	setup            setup.Setup
	isolates         []Isolate
}

// NewBackendEngine ...
func NewBackendEngine(setup setup.Setup) (BackendEngine, error) {
	isolate, err := NewIsolate()
	if err != nil {
		return BackendEngine{}, err
	}

	return BackendEngine{
		setup:    setup,
		isolates: []Isolate{isolate},
	}, nil
}

// Listen ...
func (engine *BackendEngine) Listen() error {
	// TODO:
	//  raw - Execute functions sent in requests.
	//  fs - Execute functions linked in requests.
	//  middleware -  middleware functions are synchronous functions, their result must be gotten before proceeding.

	router := gin.Default()

	// Use GRPC instead.
	router.POST("/run", engine.Run)
	router.POST("/login", engine.LogIn)

	return router.Run()
}

func (engine *BackendEngine) fetchWorkspaceID(name string) error {
	return nil
}

func (engine *BackendEngine) getIsolate() (Isolate, error) {
	// TODO: Must be a way for creating isolates.
	return engine.isolates[0], nil
}

func (engine *BackendEngine) runScript(script string) error {
	// TODO: Run synchronous middleware first.

	isolate, err := engine.getIsolate()
	if err != nil {
		return err
	}

	// TODO:
	//  Run script in a separate goroutine. Research if multiple goroutines accessing an isolate memory is a bad thing./

	value, err := isolate.RunScript(script)
	if err != nil {
		return err
	}

	// TODO:
	//  Return result value as an object websocket message.

	fmt.Println("Value", value.String())

	return nil
}

func (engine *BackendEngine) runScriptFile(scriptFile string) error {
	// TODO:
	return nil
}
