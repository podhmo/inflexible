package main

import (
	"flag"
	"log"
	"m/todogenerated"
	"os"

	"github.com/podhmo/inflexible/clilib"
)

func main() {
	name := os.Args[0]
	subCommands := []*clilib.Command{
		NewHello(),
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

func NewHello() *clilib.Command {
	fs := flag.NewFlagSet("Hello", flag.ExitOnError)
	var options struct {
		Name string `json:"name"`
	}
	fs.StringVar(&options.Name, "name", "", "")
	return &clilib.Command{
		FlagSet: fs,
		Options: &options,
		Do:      clilib.LiftHandler(todogenerated.Hello),
	}
}
