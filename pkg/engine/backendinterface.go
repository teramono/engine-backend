package engine

import (
	"net/http"

	"github.com/teramono/utilities/pkg/request"
	"github.com/teramono/utilities/pkg/setup"
)

const Host = "http://localhost:5051" // TODO: SEC: Use Https, gRPC, setup config

// BackendInterface ...
type BackendInterface struct {
	setup *setup.Setup
}

// NewBackendInterface ...
func NewBackendInterface(setup *setup.Setup) BackendInterface {
	return BackendInterface{setup}
}

// Login ...
func (backend *BackendInterface) Login(req request.Request) (*http.Response, error) {
	req.URL.Host = Host
	return req.Send()
}

// Run ...
func (backend *BackendInterface) Run(req request.Request) (*http.Response, error) {
	req.URL.Host = Host
	return req.Send()
}
