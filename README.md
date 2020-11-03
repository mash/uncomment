Uncomment
=========

Uncomment is a simple command line tool to strip the comments out of relaxed JSON and output JSON as defined in RFC 7159, or in other words a JSON5 to JSON converter.

## Usage

```
% cat infile.json
{
  // single line comments
  /*
   * or multi line comments
   */
  foo: "bar",
}

% cat infile.json | bin/uncomment
{"foo":"bar"}

% bin/uncomment infile.json
{"foo":"bar"}

% bin/uncomment -h
Usage of bin/uncomment:
  -i string
        Input file name
  -n    Do not print the trailing newline character.
  -o string
        Output file name
```

## See also

https://json5.org/
