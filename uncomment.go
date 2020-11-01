package uncomment

import (
	"encoding/json"
	"io"

	"github.com/flynn/json5"
)

func Uncomment(r io.Reader, w io.Writer) error {
	dec := json5.NewDecoder(r)
	obj := map[string]interface{}{}
	if err := dec.Decode(&obj); err != nil {
		return err
	}
	enc := json.NewEncoder(w)
	if err := enc.Encode(obj); err != nil {
		return err
	}
	return nil
}
