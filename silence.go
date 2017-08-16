package silence

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

var Logger = &bytes.Buffer{}

func Run(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = Logger
	cmd.Stderr = Logger
	Log("# %s %+v", cmd.Path, cmd.Args)
	return cmd.Run()
}

func Log(message string, arg ...interface{}) {
	fmt.Fprintf(Logger, message+"\n", arg...)
}

func Exit(code int) {
	if code != 0 {
		// If the exit status is non-zero, then I want to see the output.
		fmt.Print(Logger.String())
	}
	os.Exit(code)
}
