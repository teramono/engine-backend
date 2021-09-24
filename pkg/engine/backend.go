package engine

// Backend ...
type Backend struct {
	server BackendServerConn
}

// BackendServerConn ...
type BackendServerConn struct{}

// Request ...
type Request struct{}

// NewBackend ...
func NewBackend() Backend {
	return Backend{}
}

// RunScriptURL ...
func (backend *Backend) RunScriptURL(url string, req Request) {}

// NewRequest ...
func NewRequest() Request {
	return Request{}
}
