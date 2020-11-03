package uncomment

import (
	"encoding/json"
	"io"

	"github.com/flynn/json5"
)

type Option int

const (
	NoTrailingNewline Option = iota
)

func Uncomment(r io.Reader, w io.Writer, options ...Option) error {
	noTrailingNewline := false
	for _, o := range options {
		if o == NoTrailingNewline {
			noTrailingNewline = true
		}
	}

	dec := json5.NewDecoder(r)
	obj := map[string]interface{}{}
	if err := dec.Decode(&obj); err != nil {
		return err
	}

	if noTrailingNewline {
		b, err := json.Marshal(obj)
		if err != nil {
			return err
		}
		if _, err := w.Write(b); err != nil {
			return err
		}
	} else {
		enc := json.NewEncoder(w)
		if err := enc.Encode(obj); err != nil {
			return err
		}
	}
	return nil
}
