package engine

import "rogchap.com/v8go"

type Isolate struct {
	*v8go.Isolate
}

func NewIsolate() (Isolate, error) {
	isolate, err := v8go.NewIsolate()
	if err != nil {
		return Isolate{}, err
	}

	return Isolate{
		Isolate: isolate,
	}, nil
}

func (isolate *Isolate) RunScript(script string) (*v8go.Value, error) {
	context, err := v8go.NewContext()
	if err != nil {
		return nil, err
	}

	return context.RunScript(script, "")
}
