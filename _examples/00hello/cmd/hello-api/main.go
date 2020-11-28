package main

import (
	"log"
	"net/http"
	"os"

	"m/todogenerated"

	"github.com/podhmo/inflexible/weblib"
	"github.com/podhmo/tenuki"
)

func main() {
	addr := ":33333"
	if v := os.Getenv("ADDR"); v != "" {
		addr = v
	}

	run := func() error {
		mux := SetupHTTPHandler()
		return Run(mux, addr)
	}
	if err := run(); err != nil {
		log.Fatalf("!! %+v", err)
	}
}

func Run(mux http.Handler, addr string) error {
	// todo: graceful stop
	log.Printf("listen %s...", addr)
	return http.ListenAndServe(addr, mux)
}

// TODO: generate automatically

func SetupHTTPHandler() http.Handler {
	mux := http.DefaultServeMux // todo: fix

	// TODO: generate the endpoint returns openapi doc
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tenuki.Render(w, r).JSON(200, map[string]interface{}{
			"methods": []string{"Hello"},
		})
	})

	mux.HandleFunc("/Hello", weblib.LiftHandler(todogenerated.Hello))
	return mux
}
