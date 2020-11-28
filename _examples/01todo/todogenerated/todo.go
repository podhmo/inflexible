package todogenerated

import (
	"context"

	"m/design"

	"github.com/podhmo/inflexible"
	"github.com/podhmo/tenuki"
)

// -- handler TODO: generate automatically

func AddTodo(ctx context.Context, ev inflexible.Event) (interface{}, error) {
	var input struct {
		Todo design.Todo `json:"todo" validate:"required"`
	}
	if err := tenuki.DecodeJSON(ev.Body, &input); err != nil {
		return nil, inflexible.NewAppError(err, 400)
	}

	registry, err := design.GetRegistry(ctx)
	if err != nil {
		return nil, inflexible.NewAppError(err, 500)
	}
	return design.AddTodo(ctx, registry.Store, input.Todo)
}

func ListTodo(ctx context.Context, ev inflexible.Event) (interface{}, error) {
	var input struct {
		All bool `json:"all"`
	}
	if err := tenuki.DecodeJSON(ev.Body, &input); err != nil {
		return nil, inflexible.NewAppError(err, 400)
	}

	registry, err := design.GetRegistry(ctx)
	if err != nil {
		return nil, inflexible.NewAppError(err, 500)
	}
	return design.ListTodo(ctx, registry.Store, &input.All)
}
