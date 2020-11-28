package design

import (
	"context"
	"m/store"
)

type Registry struct {
	Store store.Store
}

func GetRegistry(ctx context.Context) (*Registry, error) {
	return &Registry{
		Store: &store.FileStore{},
	}, nil
}
