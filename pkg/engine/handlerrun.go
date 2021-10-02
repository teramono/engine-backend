package engine

import (
	"fmt"
	"strings"

	"github.com/teramono/utilities/pkg/configs"
	"github.com/teramono/utilities/pkg/file"
	"github.com/teramono/utilities/pkg/request"
	"rogchap.com/v8go"
)

func (server *BackendServer) fetchAndRunAuthScript(req *request.Request) error {
	fmt.Println(">>>", 1)
	// Fetch auth script.
	authScript, err := server.fetchFile(req, "/system/auth.js")
	if err != nil {
		return err
	}

	fmt.Println(">>>", 2)

	result, err := server.runScript(authScript)
	if err != nil {
		return err
	}

	fmt.Println(">> authscript result =", result)

	if !result.Boolean() {
		return fmt.Errorf("authentication not successful") // TODO: More context might help.
	}

	fmt.Println(">>>", 3)
	return nil
}

func (server *BackendServer) fetchAndRunBundledScript(req *request.Request) (*v8go.Value, error) {
	// Path
	path := trimPrefixFromPath(req)

	// Fetch manifest
	manifest, err := server.fetchFile(req, "/surl/"+path+"/manifest.json")
	if err != nil {
		return &v8go.Value{}, err
	}

	// Fetch middleware scripts from manifest.
	middlewareScripts, err := server.fetchMiddlewareScripts(req, manifest)
	if err != nil {
		return &v8go.Value{}, err
	}

	// Fetch index script.
	indexScript, err := server.fetchFile(req, "/surl/"+path+"/index.js")
	if err != nil {
		return &v8go.Value{}, err
	}

	// Bundle scripts together.
	bundledScript, err := bundleScripts(append(middlewareScripts, indexScript))
	if err != nil {
		return &v8go.Value{}, err
	}

	// Run bundled scripts.
	return server.runScript(bundledScript)
}

func (server *BackendServer) fetchMiddlewareScripts(req *request.Request, manifest []byte) ([][]byte, error) {
	// Fetch middleware scripts in manifest.
	scriptManifest, err := configs.NewScriptManifest(string(manifest), configs.JSON)
	if err != nil {
		return [][]byte{}, nil
	}

	// TODO: Need an fs.ReadMultiple implementation.
	var scripts = [][]byte{}
	for _, middlewarePath := range scriptManifest.MiddlewarePaths {
		// TODO: Implement canonicalize
		absolutePath := file.Canonicalize(middlewarePath, "/")
		bytes, err := server.fetchFile(req, absolutePath)
		if err != nil {
			return [][]byte{}, err
		}

		scripts = append(scripts, bytes)
	}

	return scripts, nil
}

// SEC: TODO: What is the right strategy for bundling.
func bundleScripts(scripts [][]byte) ([]byte, error) {
	// TODO: Bundle the scripts.
	return []byte{}, nil
}

func trimPrefixFromPath(req *request.Request) string {
	return strings.TrimPrefix("/run", req.URL.Path)
}
