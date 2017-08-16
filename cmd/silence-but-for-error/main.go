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
