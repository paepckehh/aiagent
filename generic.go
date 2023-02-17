package aiagent

import (
	_ "embed"
	"syscall"
)

const (
	_dot       = "."
	_linefeed  = "\n" // linefeed
	_linefeedR = '\n' // linefeed
	_msgHeader = 2    // message header line
)

//go:embed mandant.db
var db string

// getEnv ...
func getEnv(in string) (string, bool) {
	return syscall.Getenv(in)
}

// bracket
func bracket(in string) string {
	return "[" + in + "]"
}
