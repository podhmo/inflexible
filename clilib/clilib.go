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
	var POSTDATA json.RawMessage // support same-expression as web API

	return func(path []*Command, args []string) error {
		cmd := path[len(path)-1]

		cmd.Var(
			(*LiteralOrFileContentValue)(&POSTDATA),
			"POSTDATA", "json-string or @<file-name>, support same-expression as web API",
		)

		if err := cmd.Parse(args); err != nil {
			return err
		}
		for _, cb := range callbacks {
			if err := cb(); err != nil {
				cmd.Usage()
				return err
			}
		}

		var b []byte
		if !isEmptyJSON(POSTDATA) {
			b = POSTDATA
		} else {
			var err error
			b, err = json.MarshalIndent(cmd.Options, "", "  ")
			if err != nil {
				cmd.Usage()
				return err
			}
		}
		if ok, _ := strconv.ParseBool(os.Getenv("DEBUG")); ok {
			fmt.Fprintln(os.Stderr, "----------------------------------------")
			fmt.Fprintln(os.Stderr, "data: ", string(b))
			fmt.Fprintln(os.Stderr, "----------------------------------------")
		}

		var dummy url.Values
		ev := inflexible.Event{
			Name:    cmd.Name(),
			Body:    ioutil.NopCloser(bytes.NewBuffer(b)),
			Headers: dummy,
		}

		ctx := context.Background()
		result, err := h(ctx, ev)
		if err != nil {
			if x, ok := err.(inflexible.HasCode); ok && x.Code() == 400 {
				cmd.Usage()
			}
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
