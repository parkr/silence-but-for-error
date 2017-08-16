# silence-but-for-error

Runs a subprocess, recording any output, and only prints the output if
exiting with a non-zero exit code.

Provided as a library & a binary.

## Installation

```text
$ go get github.com/parkr/silence-but-for-error/cmd/silence-but-for-error
```

## Usage

```text
$ silence-but-for-error echo Hello, World
$ silence-but-for-error false
# /bin/false [false]
Command failed: exit status 1
```

```go
package main

import (
	"flag"

	"github.com/parkr/silence-but-for-error"
)

func main() {
	flag.Parse()

	args := flag.Args()

	if len(args) < 1 {
		silence.Log("usage: silence-but-for-error <command> [args...]")
		silence.Exit(1)
	}

	if err := silence.Run(args[0], args[1:]...); err != nil {
		silence.Log("Command failed: %+v", err)
		silence.Exit(1)
	}
}
```
