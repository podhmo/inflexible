package design

import (
	"context"
	"fmt"
)

func Hello(ctx context.Context, name string, short *bool) (string, error) {
	if short != nil && *short {
		return fmt.Sprintf("Hi %s", name), nil
	}
	return fmt.Sprintf("Hello %s", name), nil
}
