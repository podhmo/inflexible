package weblib

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/podhmo/inflexible"
	"github.com/podhmo/tenuki"
)

type QueryOrHeader struct {
	Query  url.Values
	Header http.Header
}

func (v *QueryOrHeader) Get(key string) string {
	val := v.Query.Get(key)
	if val != "" {
		return val
	}
	return v.Header.Get(key)
}

// TODO: rename
func LiftHandler(h inflexible.Handler) http.HandlerFunc {
	name := fmt.Sprintf("%v", h) // todo: fix
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			t := recover()
			if t != nil {
				log.Printf("hmm %+v", t)
				tenuki.Render(w, r).JSON(500, map[string]string{"message": fmt.Sprintf("%s", t)})
			}
		}()

		ev := inflexible.Event{
			Name: name,
			Body: r.Body,
			QueryOrHeader: &QueryOrHeader{
				Query:  r.URL.Query(),
				Header: r.Header,
			},
		}

		ctx := inflexible.WithEvent(r.Context(), ev)
		result, err := h.Handle(ctx, ev)
		if err != nil {
			code := 500
			if x, ok := err.(inflexible.HasCode); ok {
				code = x.Code()
			}
			tenuki.Render(w, r).JSON(code, map[string]string{"message": err.Error()})
			return
		}
		tenuki.Render(w, r).JSON(200, result)
	}
}
