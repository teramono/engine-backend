package engine

import "github.com/teramono/utilities/pkg/setup"

// Backend ...
type Backend struct {
	server BackendServerConn
}

// BackendServerConn ...
type BackendServerConn struct {
	address string
	port    string
}

// Request ...
type Request struct {
	URL string
}

// NewBackend ...
func NewBackend(setup *setup.Setup) Backend {
	return Backend{}
}

// NewRequest ...
func NewRequest() Request {
	return Request{}
}

// Login ...
func (backend *Backend) Login(workspaceID string) {
	// Send to route.
}

// Run ...
func (backend *Backend) Run(workspaceID string, req Request) {
	// Send to route.
}
