package todogenerated

import (
	"context"
	"strconv"

	"m/design"

	"github.com/go-playground/validator/v10"
	"github.com/podhmo/inflexible"
	"github.com/podhmo/tenuki"
)

// -- handler TODO: generate automatically

func Hello(ctx context.Context, ev inflexible.Event) (interface{}, error) {
	var input struct {
		Name  string `json:"name" validate:"required"`
		Short bool   `json:"short"`
	}
	if err := tenuki.DecodeJSON(ev.Body, &input); err != nil {
		return nil, inflexible.NewAppError(err, 400)
	}
	if ok, err := strconv.ParseBool(ev.QueryOrHeader.Get("short")); err == nil {
		input.Short = ok
	}

	// todo: customize validation
	if err := validate.Struct(&input); err != nil {
		return nil, inflexible.NewAppError(err, 400)
	}
	return design.Hello(ctx, input.Name, &input.Short)
}

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

func init() {
	validate = validator.New()
}
