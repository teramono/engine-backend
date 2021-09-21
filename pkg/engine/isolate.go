package engine

import "rogchap.com/v8go"

// Isolate ...
type Isolate struct {
	*v8go.Isolate
}

// NewIsolate ...
func NewIsolate() (Isolate, error) {
	isolate, err := v8go.NewIsolate()
	if err != nil {
		return Isolate{}, err
	}

	return Isolate{
		Isolate: isolate,
	}, nil
}

// RunScript ...
func (isolate *Isolate) RunScript(script string) (*v8go.Value, error) {
	context, err := v8go.NewContext()
	if err != nil {
		return nil, err
	}

	return context.RunScript(script, "")
}
