package tests

import (
	"testing"

	"github.com/teramono/engine-backend/pkg/engine"
)

func TestBackend(t *testing.T) {
	backend := engine.NewBackend()
	backend.RunScriptURL("/users?name=John", engine.NewRequest())
}
