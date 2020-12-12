package inflexible

import (
	"context"
	"io"
)

type HandlerFunc func(context.Context, Event) (interface{}, error)
func (f HandlerFunc) Handle(ctx context.Context, ev Event) (interface{}, error){
	return f(ctx, ev)
}
type Handler interface{
	Handle(context.Context, Event) (interface{}, error)
}
type getter interface {
	Get(string) string
}

type Event struct {
	Name    string
	QueryOrHeader getter // TODO: refine
	Body    io.ReadCloser
}

// -- context

type key string

const (
	keyEvent key = ":event"
)

func WithEvent(ctx context.Context, ev Event) context.Context {
	return context.WithValue(ctx, keyEvent, ev)
}
