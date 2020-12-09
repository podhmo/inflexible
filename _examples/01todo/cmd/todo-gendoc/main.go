package main

import (
	"m/design"
	"reflect"
	"strings"

	reflectopenapi "github.com/podhmo/reflect-openapi"
)

// TODO: merge

func main() {
	c := reflectopenapi.Config{
		SkipValidation: true,
		Selector: &struct {
			reflectopenapi.MergeParamsInputSelector
			reflectopenapi.FirstParamOutputSelector
		}{},
		IsRequiredCheckFunction: func(tag reflect.StructTag) bool {
			return strings.Contains(tag.Get("validate"), "required")
		},
	}
	c.EmitDoc(func(m *reflectopenapi.Manager) {
		{
			op := m.Visitor.VisitFunc(design.ListTodo)
			m.Doc.AddOperation("/ListTodo", "POST", op)
		}
		{
			op := m.Visitor.VisitFunc(design.AddTodo)
			m.Doc.AddOperation("/AddTodo", "POST", op)
		}
	})
}
