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
