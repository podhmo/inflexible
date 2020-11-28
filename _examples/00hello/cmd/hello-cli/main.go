package main

import (
	"flag"
	"log"
	"m/todogenerated"
	"os"

	"github.com/podhmo/inflexible/clilib"
)

func main() {
	cmd := clilib.NewRouterCommand(
		os.Args[0], []*clilib.Command{
			NewHello(),
		},
	)

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