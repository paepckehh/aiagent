package aiagent

import (
	_ "embed"
	"strconv"
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

// yesno  ...
func yesno(state bool) string {
	if state {
		return "[yes]"
	}
	return "[no]"
}

// valid  ...
func valid(state bool) string {
	if state {
		return "[valid]"
	}
	return "[failed]"
}

// validExit  ...
func validExit(state bool) string {
	if state {
		return "[valid]"
	}
	return "[failed] [exit]"
}

// ctoa confidence float64 to ascii percentage
func ctoa(in float64) string {
	return strconv.Itoa(int(in*100)) + "%"
}

// itoa int to ascii
func itoa(in int) string {
	return strconv.Itoa(in)
}

// price float to US$
func price(in float64) string {
	return strconv.FormatFloat(in, 'f', 5, 64) + " US$ "
}
