package main

import (
	"flag"
	"fmt"
	"os/exec"
	"strconv"
	"syscall"

	"github.com/parkr/silence-but-for-error"
)

type intSlice struct {
	codes []int
}

func (i *intSlice) String() string {
	return fmt.Sprintf("%+v", i.codes)
}

func (i *intSlice) Set(exitCodeStr string) error {
	exitCode, err := strconv.Atoi(exitCodeStr)
	if err != nil {
		return err
	}

	i.codes = append(i.codes, exitCode)
	return nil
}

func main() {
	var ignoredExitCodes intSlice
	flag.Var(&ignoredExitCodes, "ignore-exit-code", "Ignore a non-successful exit code (you may specify multiple times).")
	flag.Parse()

	fmt.Printf("ignored exit codes: %+v\n", ignoredExitCodes)

	args := flag.Args()

	if len(args) < 1 {
		silence.Log("usage: silence-but-for-error <command> [args...]")
		silence.Exit(1)
	}

	if err := silence.Run(args[0], args[1:]...); err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			// The program has exited with an exit code != 0

			// This works on both Unix and Windows. Although package syscall
			// is generally platform dependent, WaitStatus is defined for
			// both Unix and Windows and in both cases has an ExitStatus()
			// method with the same signature.
			if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				exitStatus := status.ExitStatus()
				for _, ignored := range ignoredExitCodes.codes {
					if ignored == exitStatus {
						silence.Exit(0)
					}
				}
			}
		} else {
			silence.Log("Command failed: %+v", err)
			silence.Exit(1)
		}
	}
}
