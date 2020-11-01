package uncomment

import (
	"errors"
	"flag"
	"io"
	"os"

	"github.com/mattn/go-isatty"
)

type Flags struct {
	input, output string
	Reader        io.ReadCloser
	Writer        io.WriteCloser
}

func NewFlags() (Flags, error) {
	f := parseFlags()
	if isatty.IsTerminal(os.Stdin.Fd()) {
		if f.input == "" {
			return f, errors.New("input missing")
		}
		r, err := os.Open(f.input)
		if err != nil {
			return f, err
		}
		f.Reader = r
	} else {
		if f.input != "" {
			return f, errors.New("input can be either STDIN or a file, not both")
		}
		f.Reader = os.Stdin
	}

	if f.output != "" {
		w, err := os.OpenFile(f.output, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return f, err
		}
		f.Writer = w
	} else {
		f.Writer = os.Stdout
	}
	return f, nil
}

func (f Flags) Close() error {
	err := f.Reader.Close()
	err2 := f.Writer.Close()
	if err != nil {
		return err
	}
	if err2 != nil {
		return err2
	}
	return nil
}

func parseFlags() Flags {
	f := Flags{}
	flag.StringVar(&f.input, "i", "", "Input file name")
	flag.StringVar(&f.output, "o", "", "Output file name")
	flag.Parse()
	return f
}
