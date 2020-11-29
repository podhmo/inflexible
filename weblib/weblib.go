package weblib

import (
	"fmt"
	"log"
	"net/http"

	"github.com/podhmo/inflexible"
	"github.com/podhmo/tenuki"
)

// TODO: rename
func LiftHandler(h inflexible.HandlerFunc) http.HandlerFunc {
	name := fmt.Sprintf("%v", h) // todo: fix
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			t := recover()
			if t != nil {
				log.Printf("hmm %+v", t)
				tenuki.Render(w, r).JSON(500, map[string]string{"message": fmt.Sprintf("%s", t)})
			}
		}()

		// TODO:headers and queries
		ev := inflexible.Event{
			Name:    name,
			Body:    r.Body,
			Headers: r.URL.Query(), // + headers
		}

		ctx := inflexible.WithEvent(r.Context(), ev)
		result, err := h(ctx, ev)
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
