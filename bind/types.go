package bind

import (
	"context"
	"io"
)

type Symbol struct {
	Name    string
	Package *Package
}

type Package struct {
	Name string
	Path string
}

type EmitFunc func(ctx context.Context, w io.Writer, filename string) error
type Code struct {
	*Symbol
	Emit    EmitFunc
	Deps    []*Code
	Imports []string
}
