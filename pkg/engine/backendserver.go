// Package engine ...
package engine

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	dbEngine "github.com/teramono/engine-db/pkg/engine"
	fsEngine "github.com/teramono/engine-fs/pkg/engine"
	"github.com/teramono/utilities/pkg/file"
	"github.com/teramono/utilities/pkg/request"
	"github.com/teramono/utilities/pkg/response"
	"github.com/teramono/utilities/pkg/setup"
	"rogchap.com/v8go"
)

type BackendServer struct {
	isolates []Isolate // TODO: Details around isolate provisioning.
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

func (server *BackendServer) runScript(script []byte) (*v8go.Value, error) {
	// TODO: How to determine what isolates to use.
	mainIsolate := &server.isolates[0]

	// TODO: Create custom capability context for each script.
	context, err := v8go.NewContext(mainIsolate)
	if err != nil {
		return &v8go.Value{}, err
	}

	fmt.Println(">> script =", string(script))

	return context.RunScript(string(script), "")
}

func (server *BackendServer) fetchFile(req *request.Request, path string) ([]byte, error) {
	canonicalPath := file.Canonicalize(path, "/")
	resp, err := server.fs.Read(req, canonicalPath)
	if err != nil {
		return []byte{}, err
	}

	if resp.StatusCode != http.StatusOK {
		bodyBytes, err := response.ReadBodyBytes(resp)
		if err != nil {
			return []byte{}, err
		}

		return []byte{}, fmt.Errorf(string(bodyBytes))
	}

	return response.ReadBodyBytes(resp)
}
