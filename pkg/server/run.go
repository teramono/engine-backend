package server

import (
	"fmt"

	"github.com/teramono/tera/pkg/runtime"
	"github.com/teramono/utilities/pkg/configs"
	"github.com/teramono/utilities/pkg/messages"
)

func (server *BackendServer) runSurlIndexScript(canonicalWorkspacePath CanonicalWorkspacePath, pathFromURLSuffix string) ([]byte, error) {
	indexScriptPath := server.getSurlPath(pathFromURLSuffix, IndexScriptFilename)
	fileBytes, err := server.readRelativeFile(canonicalWorkspacePath, indexScriptPath)
	if err != nil {
		return []byte{}, err
	}

	// Create temp runtime.
	rt, err := runtime.NewRuntime()
	if err != nil {
		return []byte{}, err
	}

	// Run script.
	result, err := rt.RunScript(runtime.Script{
		Filename: indexScriptPath,
		Content:  fileBytes,
	})
	if err != nil {
		return []byte{}, err
	}

	// Expected an object.
	obj, err := result.AsObject()
	if err != nil {
		return []byte{}, fmt.Errorf("%s", messages.WrongReturnType("object", indexScriptPath))
	}

	return obj.MarshalJSON()
}

func (server *BackendServer) runAuthScript(canonicalWorkspacePath CanonicalWorkspacePath) (bool, error) {
	fileBytes, err := server.readRelativeFile(canonicalWorkspacePath, authScriptPath)
	if err != nil {
		return false, err
	}

	// Create temp runtime.
	rt, err := runtime.NewRuntime()
	if err != nil {
		return false, err
	}

	// Run script.
	result, err := rt.RunScript(runtime.Script{
		Filename: authScriptPath,
		Content:  fileBytes,
	})
	if err != nil {
		return false, err
	}

	// Expected an boolean.
	if !result.IsBoolean() {
		return false, fmt.Errorf("%s", messages.WrongReturnType("boolean", authScriptPath))
	}

	return result.Boolean(), nil
}

func (server *BackendServer) runMiddlewareScripts(canonicalWorkspacePath CanonicalWorkspacePath, surlManifest *configs.SurlManifest) (bool, error) {
	scripts, err := server.fetchMiddlewareScripts(canonicalWorkspacePath, surlManifest)
	if err != nil {
		return false, err
	}

	for _, script := range scripts {
		// Create temp runtime each time.
		rt, err := runtime.NewRuntime()
		if err != nil {
			return false, err
		}

		// Run script.
		result, err := rt.RunScript(script)
		if err != nil {
			return false, err
		}

		// Expected an boolean.
		if !result.IsBoolean() {
			return false, fmt.Errorf("%s", messages.WrongReturnType("boolean", script.Filename))
		}

		// Check the result is not false.
		if !result.Boolean() {
			return false, nil
		}
	}

	return true, nil
}
