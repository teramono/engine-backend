package crud

import (
	"rogchap.com/v8go"
)

type Script struct {
	Filename string
	Content  []byte
}

func RunScript(script Script) (*v8go.Value, error) {
	// TODO: Create custom capability context for each run.
	isolate, err := v8go.NewIsolate()
	if err != nil {
		return &v8go.Value{}, err
	}

	context, err := v8go.NewContext(isolate)
	if err != nil {
		return &v8go.Value{}, err
	}

	return context.RunScript(string(script.Content), script.Filename)
}
