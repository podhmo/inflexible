package design

import (
	"context"
	"m/store"

	"golang.org/x/xerrors"
)

type Todo struct {
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func ListTodo(ctx context.Context, store store.Store, all *bool) ([]Todo, error) {
	var list []Todo
	if err := store.Load(ctx, "todo.json", &list); err != nil {
		return nil, xerrors.Errorf("load: %w", err)
	}

	if all == nil || !*all {
		r := make([]Todo, 0, len(list))
		for _, x := range list {
			if x.Done {
				r = append(r, x)
			}
		}
		list = r
	}
	return list, nil
}

func AddTodo(ctx context.Context, store store.Store, todo Todo) (*Todo, error) {
	var list []Todo
	if err := store.Load(ctx, "todo.json", &list); err != nil {
		return nil, xerrors.Errorf("load: %w", err)
	}

	list = append(list, todo)

	if err := store.Save(ctx, "todo.json", &list); err != nil {
		return nil, xerrors.Errorf("save: %w", err)
	}
	return &todo, nil
}
