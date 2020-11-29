package todogenerated

import (
	"context"

	"m/design"

	"github.com/go-playground/validator/v10"
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

	// todo: customize validation
	if err := validate.Struct(&input); err != nil {
		return nil, inflexible.NewAppError(err, 400)
	}

	registry, err := design.GetRegistry(ctx)
	if err != nil {
		return nil, err
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

	// todo: customize validation
	if err := validate.Struct(&input); err != nil {
		return nil, inflexible.NewAppError(err, 400)
	}

	registry, err := design.GetRegistry(ctx)
	if err != nil {
		return nil, err
	}
	return design.ListTodo(ctx, registry.Store, &input.All)
}

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

func init() {
	validate = validator.New()
}
