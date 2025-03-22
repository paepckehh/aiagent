package aiagent

import (
	_ "embed"
	"strconv"
	"strings"
	"syscall"
)

const (
	_dot           = "."
	_linefeed      = "\n" // linefeed
	_linefeedR     = '\n' // linefeed
	_msgHeader     = 2    // message header line
	_err           = " [error]"
	_exit          = " [exit]"
	_dlinefeed     = "\n\n"
	_space         = " "
	_sep           = ';'
	_sep1          = ","
	_sep2          = " / "
	_unknownBody   = "unknown -> invalid message body"
	_unknownHeader = "unknown -> invalid message header"
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
func ctoa(m *EMail) string {
	var s strings.Builder
	s.WriteString(strconv.Itoa(int(m.Local.Lang.Confidence*100)) + "% ")
	if m.Local.Lang.Confidence < m.Local.TargetLangConfidence {
		s.WriteString("[exit]")
	} else {
		s.WriteString("[valid]")
	}
	return s.String()
}

// itoa int to ascii
func itoa(in int) string {
	return strconv.Itoa(in)
}

// price float to US$
func price(in float64) string {
	return strconv.FormatFloat(in, 'f', 5, 64) + " US$ "
}

// isArray
func isArray(in []string) string {
	if len(in) < 1 {
		return "[none]"
	}
	return strings.Join(in, _sep1)
}
