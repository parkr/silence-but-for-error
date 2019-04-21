package silence

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

// CommandStatus represents the success of the command.
type CommandStatus int

const (
	// Unknown is a fallback CommandStatus.
	Unknown CommandStatus = iota
	// Failure is a failing command.
	Failure
	Success
)

// Runner is a container for running a command.
type Runner struct {
	cmd     *exec.Cmd
	command []string
	output  *bytes.Buffer
}

// NewRunner creates a new runner.
func NewRunner() *Runner {
	return &Runner{
		output: &bytes.Buffer{},
	}
}

// Run executes a command.
func (r *Runner) Run(name string, arg ...string) error {
	r.cmd = exec.Command(name, arg...)
	r.cmd.Stdin = os.Stdin
	r.cmd.Stdout = r.output
	r.cmd.Stderr = r.output

	r.Log("# %s %+v", r.cmd.Path, r.cmd.Args)
	r.command = []string{r.cmd.Path}
	r.command = append(r.command, r.cmd.Args...)

	return r.cmd.Run()
}

// Log writes some formatted text to the Runner's output.
func (r *Runner) Log(format string, arg ...interface{}) {
	fmt.Fprintf(r.output, format+"\n", arg...)
}

// Output returns the Runner's output.
func (r *Runner) Output() string {
	return r.output.String()
}

// Exit terminates the command. If the CommandStatus parameter is "Failure",
// then the Runner also prints its output to stdout.
func (r *Runner) Exit(c CommandStatus) {
	if c == Failure {
		// If the exit status is non-zero, then I want to see the output.
		fmt.Print(r.Output())
		os.Exit(1)
	}
	os.Exit(0)
}
