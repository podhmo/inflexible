package todogenerated

import (
	"context"
	"strconv"

	"m/design"
	"m/store"

	"github.com/go-playground/validator/v10"
	"github.com/podhmo/inflexible"
	"github.com/podhmo/tenuki"
)

// -- handler TODO: generate automatically
type addTodoFunc func(ctx context.Context, store store.Store, todo design.Todo) (*design.Todo, error)

func (f addTodoFunc) Handle(ctx context.Context, ev inflexible.Event) (interface{}, error) {
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
	return f(ctx, registry.Store, input.Todo)
}

type listTodoFunc func(ctx context.Context, store store.Store, all *bool) ([]design.Todo, error)

func (f listTodoFunc) Handle(ctx context.Context, ev inflexible.Event) (interface{}, error) {
	var input struct {
		All bool `json:"all"`
	}
	if ok, err := strconv.ParseBool(ev.QueryOrHeader.Get("all")); err == nil {
		input.All = ok
	}

	// todo: customize validation
	if err := validate.Struct(&input); err != nil {
		return nil, inflexible.NewAppError(err, 400)
	}

	registry, err := design.GetRegistry(ctx)
	if err != nil {
		return nil, err
	}
	return f(ctx, registry.Store, &input.All)
}

// bind

var (
	AddTodo  = addTodoFunc(design.AddTodo)
	ListTodo = listTodoFunc(design.ListTodo)
)

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

func init() {
	validate = validator.New()
}
