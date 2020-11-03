package uncomment

import (
	"bytes"
	"io"
	"testing"
)

func TestUncomment(t *testing.T) {
	tests := []struct {
		in            io.Reader
		expected      string
		expectedError string
		options       []Option
	}{
		{
			in: bytes.NewBufferString(`{k:/* commented */ "v",}`),
			expected: `{"k":"v"}
`,
		},
		{
			in:       bytes.NewBufferString(`{k:/* commented */ "v",}`),
			expected: `{"k":"v"}`,
			options:  []Option{NoTrailingNewline},
		},
		{
			in:            bytes.NewBufferString(`{k:// "v",}`),
			expected:      `{"k":"v"}`,
			expectedError: "unexpected EOF",
		},
		{
			in: bytes.NewBufferString(`[1]`),
			expected: `[1]
`,
		},
		{
			in: bytes.NewBufferString(`"foo"`),
			expected: `"foo"
`,
		},
	}
	for _, test := range tests {
		out := &bytes.Buffer{}

		err := Uncomment(test.in, out, test.options...)
		if err != nil || test.expectedError != "" {
			if e, g := test.expectedError, err; e != g.Error() {
				t.Errorf("expected error |%s| but got |%s|\n", e, g)
			}
		} else if e, g := test.expected, out.String(); e != g {
			t.Errorf("expected |%s| but got |%s|\n", e, g)
		}
	}
}
