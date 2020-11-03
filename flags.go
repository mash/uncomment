package uncomment

import (
	"errors"
	"flag"
	"io"
	"os"

	"github.com/mattn/go-isatty"
)

type Flags struct {
	input, output     string
	noTrailingNewline bool
}

func (f Flags) Options() Options {
	opts := Options{}
	if f.noTrailingNewline {
		opts.NoTrailingNewline = true
	}
	if f.output != "" {
		opts.PrintOutputFilename = f.output
	}
	return opts
}

func ParseFlags() Flags {
	f := Flags{}
	flag.StringVar(&f.input, "i", "", "Input file name.")
	flag.StringVar(&f.output, "o", "", "Output file name.")
	flag.BoolVar(&f.noTrailingNewline, "n", false, "Do not print the trailing newline character.")
	flag.Parse()

	// -i can be omited
	if in := flag.Arg(0); in != "" && f.input == "" {
		f.input = in
	}
	return f
}

func Session(f Flags) (io.ReadCloser, io.WriteCloser, error) {
	var r io.ReadCloser
	var w io.WriteCloser
	var err error
	if isatty.IsTerminal(os.Stdin.Fd()) {
		if f.input == "" {
			return nil, nil, errors.New("input missing")
		}
		r, err = os.Open(f.input)
		if err != nil {
			return nil, nil, err
		}
	} else {
		// piped
		if f.input != "" {
			return nil, nil, errors.New("input can be either STDIN or a file, not both")
		}
		r = os.Stdin
	}

	if f.output != "" {
		w, err = os.OpenFile(f.output, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return nil, nil, err
		}
	} else {
		w = os.Stdout
	}
	return r, w, nil
}
