package main

import (
	"encoding/json"
	"flag"
	"log"
	"m/todogenerated"
	"os"

	"github.com/podhmo/inflexible/clilib"
)

func main() {
	name := os.Args[0]
	subCommands := []*clilib.Command{
		NewAddTodo(),
		NewListTodo(),
	}

	cmd := clilib.NewRouterCommand(name, subCommands)
	run := func() error {
		return cmd.Do([]*clilib.Command{cmd}, os.Args[1:])
	}
	if err := run(); err != nil {
		log.Fatalf("!! %+v", err)
	}
}

// TODO: generate automatically

func NewAddTodo() *clilib.Command {
	fs := flag.NewFlagSet("AddTodo", flag.ExitOnError)
	var options struct {
		Todo json.RawMessage `json:"todo"`
	}
	fs.Var(
		(*clilib.LiteralOrFileContentValue)(&options.Todo),
		"todo", "json-string or @<file-name>",
	)
	return &clilib.Command{
		FlagSet: fs,
		Options: &options,
		Do:      clilib.LiftHandler(todogenerated.AddTodo),
	}
}

func NewListTodo() *clilib.Command {
	fs := flag.NewFlagSet("ListTodo", flag.ExitOnError)
	var options struct {
		All bool `json:"all"`
	}
	fs.BoolVar(&options.All, "all", false, "")
	return &clilib.Command{
		FlagSet: fs,
		Options: &options,
		Do:      clilib.LiftHandler(todogenerated.ListTodo),
	}
}
