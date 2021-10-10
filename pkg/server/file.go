package server

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/teramono/engine-backend/pkg/crud"
	"github.com/teramono/utilities/pkg/configs"
)


// CanonicalWorkspacePath is the workspace canonical URL on the machine.
//
// Making it a type to prevent mistake.
type CanonicalWorkspacePath string


const (
	authScriptPath = "/system/auth.js"
	surlManifestFilename = "surl.yaml"
	IndexScriptFilename = "index.js"
)

func (server *BackendServer) getSurlPath(pathFromURLSuffix string, filename string) string {
	return fmt.Sprintf("surl/%s/%s", pathFromURLSuffix, filename)
}

func (server *BackendServer) getWorkspaceCanonicalPath(workspaceID string) (CanonicalWorkspacePath, error) {
	// Resolve path for canonical workspace path.
	path := filepath.Clean(filepath.Join(server.Config.Engines.Backend.RootPath, "workspaces", workspaceID))
	path, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}

	return CanonicalWorkspacePath(path), nil
}

// ReadRelativeFile reads a file in a workspace folder.
// path has to be relative to `/absolute-path-to-rootPath/workspaces/:id`
func (server *BackendServer) readRelativeFile(canonicalWorkspacePath CanonicalWorkspacePath, relativePath string) ([]byte, error) {
	// Append relative path to workspace path and canonicalize.
	path := filepath.Clean(filepath.Join(string(canonicalWorkspacePath), relativePath))
	path, err := filepath.Abs(path)
	if err != nil {
		return []byte{}, err
	}

	// Read file.
	var bytes []byte
	bytes, err = ioutil.ReadFile(path)
	if err != nil {
		return []byte{}, err
	}

	return bytes, nil
}

func (server *BackendServer) fetchSurlManifest(canonicalWorkspacePath CanonicalWorkspacePath, pathFromURLSuffix string) (configs.SurlManifest, error) {
	surlManifestPath := server.getSurlPath(pathFromURLSuffix, surlManifestFilename)

	manifestBytes, err := server.readRelativeFile(canonicalWorkspacePath, surlManifestPath)
	if err != nil {
		return configs.SurlManifest{}, err
	}

	return configs.NewSurlManifest(manifestBytes, configs.YAML)
}

func (server *BackendServer) fetchMiddlewareScripts(canonicalWorkspacePath CanonicalWorkspacePath, surlManifest *configs.SurlManifest) ([]crud.Script, error) {
	scripts := []crud.Script{}

	for _, path := range surlManifest.MiddlewareScripts {
		scriptBytes, err := server.readRelativeFile(canonicalWorkspacePath, path)
		if err != nil {
			return scripts, err
		}

		scripts = append(scripts, crud.Script{
			Filename: path,
			Content:  scriptBytes,
		})
	}

	return scripts, nil
}
