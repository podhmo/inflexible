package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"m/todogenerated"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/podhmo/inflexible/weblib"
	reflectopenapi "github.com/podhmo/reflect-openapi"
	"github.com/podhmo/reflect-openapi/handler"
)

func main() {
	addr := ":33333"
	if v := os.Getenv("ADDR"); v != "" {
		addr = v
	}

	run := func() error {
		mux := SetupHTTPHandler(addr)
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

func SetupHTTPHandler(addr string) http.Handler {
	mux := &http.ServeMux{}

	c := &reflectopenapi.Config{
		Selector: &struct {
			reflectopenapi.MergeParamsInputSelector
			reflectopenapi.FirstParamOutputSelector
		}{},
	}
	doc, err := c.BuildDoc(context.Background(), func(m *reflectopenapi.Manager) {
		{
			path := "/ListTodo"
			action := todogenerated.ListTodo
			mux.HandleFunc(path, weblib.LiftHandler(action))
			op := m.Visitor.VisitFunc(action)
			m.Doc.AddOperation(path, "POST", op)
		}
		{
			path := "/AddTodo"
			action := todogenerated.AddTodo
			mux.HandleFunc(path, weblib.LiftHandler(action))
			op := m.Visitor.VisitFunc(action)
			m.Doc.AddOperation(path, "POST", op)
		}
	})
	if err != nil {
		panic(err) // xxx
	}

	// swagger-ui
	doc.Servers = append([]*openapi3.Server{{
		URL:         fmt.Sprintf("http://localhost%s", addr),
		Description: "local development server",
	}}, doc.Servers...)
	mux.Handle("/openapi/", handler.NewHandler(doc, "/openapi/"))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/openapi", 302)
	})

	return mux
}
