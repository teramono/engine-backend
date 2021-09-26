package engine

import (
	"bytes"
	"net/http"
	"strings"
	"time"

	"github.com/teramono/utilities/pkg/request"
	"github.com/teramono/utilities/pkg/setup"
)

// BackendInterface ...
type BackendInterface struct {
	server BackendServerConn
}

// BackendServerConn ...
type BackendServerConn struct {
	address string
	port    string
}

// NewBackendInterface ...
func NewBackendInterface(setup *setup.Setup) BackendInterface {
	return BackendInterface{}
}

// Login ...
func (backend *BackendInterface) Login(workspaceID string, req request.Request) (*http.Response, error) {
	return postRequest(&req)
}

// Run ...
func (backend *BackendInterface) Run(workspaceID string, req request.Request) (*http.Response, error) {
	return postRequest(&req)
}

func postRequest(req *request.Request) (*http.Response, error) {
	targetHost := "http://localhost:5051" // SEC: TODO: Use https. Swap gRPC in. Also should be gotten from setup

	// Http client with 10s timeout.
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	// Construct a new request.
	newReq, err := http.NewRequest(req.Method, targetHost+req.URL.URI, bytes.NewReader(req.Body))
	if err != nil {
		return nil, err
	}

	// Set headers.
	for key, value := range req.Headers {
		newReq.Header.Set(key, strings.Join(value, ","))
	}

	return client.Do(newReq)
}
