// +build gen

package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/podhmo/inflexible/bind"
)

type Todo struct {
	Title string
	Done  bool
}

func UpdateTodo(ctx context.Context, todo Todo) error {
	return nil
}

func main() {
	var out string

	flag.StringVar(&out, "out", "", "out file (default stdout)")
	flag.Parse()

	var o io.Writer = os.Stdout
	if out != "" {
		f, err := os.Create(out)
		if err != nil {
			log.Fatalf("!! %+v", err)
		}
		defer f.Close()
		o = f
	}

	b := bind.NewBinder(&bind.Package{Name: "foo", Path: "github.com/podhmo/inflexible/bind/testdata/foo"})
	var script []*bind.Code
	script = append(script, b.MustBindAction("/todo/{todo_id}", UpdateTodo))

	io.WriteString(o, "// Code generated by github.com/podhmo/inflexible/bind/gen.go; DO NOT EDIT.\n")
	io.WriteString(o, "\n")
	io.WriteString(o, "package foo\n")
	io.WriteString(o, "\n")
	io.WriteString(o, "import(\n")
	seen := map[string]bool{"main": true}
	for _, code := range script {
		for _, path := range code.Imports {
			if _, ok := seen[path]; ok {
				continue
			}
			seen[path] = true
			fmt.Fprintf(o, "	%q\n", path)
		}
	}
	io.WriteString(o, ")\n")
	io.WriteString(o, "\n")
	for _, code := range script {
		bind.Emit(o, code)
	}
}
