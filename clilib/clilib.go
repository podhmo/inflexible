package clilib

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"strconv"

	"github.com/podhmo/inflexible"
)

type DoFunc func(path []*Command, args []string) error

func LiftHandler(h inflexible.HandlerFunc, callbacks ...func() error) DoFunc {
	return func(path []*Command, args []string) error {
		cmd := path[len(path)-1]
		if err := cmd.Parse(args); err != nil {
			return err
		}
		for _, cb := range callbacks {
			if err := cb(); err != nil {
				cmd.Usage()
				return err
			}
		}
		name := cmd.Name()

		var b bytes.Buffer
		enc := json.NewEncoder(&b)
		if err := enc.Encode(cmd.Options); err != nil {
			cmd.Usage()
			return err
		}

		// TODO: fix
		if ok, _ := strconv.ParseBool(os.Getenv("DEBUG")); ok {
			fmt.Fprintln(os.Stderr, "----------------------------------------")
			fmt.Fprint(os.Stderr, "data: ", b.String())
			fmt.Fprintln(os.Stderr, "----------------------------------------")
		}

		var dummy url.Values
		ev := inflexible.Event{
			Name:    name,
			Body:    ioutil.NopCloser(&b),
			Headers: dummy,
		}

		ctx := context.Background()
		result, err := h(ctx, ev)
		if err != nil {
			return err
		}

		// TODO: result handling customization
		{
			enc := json.NewEncoder(os.Stdout)
			enc.SetIndent("", "  ")
			return enc.Encode(result)
		}
	}
}
