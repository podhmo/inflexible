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

	if err := Run(addr); err != nil {
		log.Fatalf("!! %+v", err)
	}
}

func Run(addr string) error {
	mux := http.DefaultServeMux
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tenuki.Render(w, r).JSON(200, map[string]interface{}{
			"methods": []string{"Hello"},
		})
	})
	mux.HandleFunc("/Hello", weblib.LiftHandler(todogenerated.Hello))

	log.Printf("listen %s...", addr)
	return http.ListenAndServe(addr, mux)
}
