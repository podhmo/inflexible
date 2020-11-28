package clilib

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

// implement flag.Value
var emptyJSON = json.RawMessage(`{}`)

func isEmptyJSON(msg json.RawMessage) bool {
	if len(msg) == 0 {
		return true
	}

	n := len(emptyJSON)
	if n != len(msg) {
		return false
	}
	for i := 0; i < n; i++ {
		if msg[i] != emptyJSON[i] {
			return false
		}
	}
	return true
}

type LiteralOrFileContentValue json.RawMessage

func (v *LiteralOrFileContentValue) String() string {
	return string(*v) // ???
}
func (v *LiteralOrFileContentValue) Set(filenameOrContent string) error {
	if filenameOrContent == "" {
		*v = LiteralOrFileContentValue(emptyJSON)
		return nil
	}

	if !strings.HasPrefix(filenameOrContent, "@") {
		*v = LiteralOrFileContentValue(filenameOrContent)
		return nil
	}

	b, err := ioutil.ReadFile(strings.TrimPrefix(filenameOrContent, "@"))
	if err != nil {
		return err
	}
	*v = LiteralOrFileContentValue(b)
	return nil
}
