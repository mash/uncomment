package uncomment

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/flynn/json5"
)

type Options struct {
	NoTrailingNewline   bool
	PrintOutputFilename string
}

func Uncomment(r io.Reader, w io.Writer, options Options) error {
	dec := json5.NewDecoder(r)
	var obj interface{}
	if err := dec.Decode(&obj); err != nil {
		return err
	}

	if options.NoTrailingNewline {
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

	if options.PrintOutputFilename != "" {
		fmt.Printf("%s", options.PrintOutputFilename)
	}

	return nil
}
