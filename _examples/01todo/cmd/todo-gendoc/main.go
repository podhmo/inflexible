package main

import (
	"m/design"

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
