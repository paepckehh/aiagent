package main

import (
	"io"
	"os"

	"paepcke.de/aiagent"
)

const (
	_app      = "[aiagent]"
	_err      = " [error]"
	_exit     = " [exit]"
	_linefeed = "\n"
	_space    = " "
)

func main() {
	if !isPipe() {
		errExit("Error:Unable to read from pipe.\nExample: cat message.txt | aiagent ")
	}
	m := aiagent.EMail{
		Raw: getPipe(),
	}
	m.Local.TargetLangConfidence = 0.84
	_ = m.ProcessLocal()
	if m.Local.AddrRFC && m.Local.AddrMX && m.Local.Lang.Confidence > m.Local.TargetLangConfidence {
		//		_ = m.ProcessOpenAI()
		_ = m.ProcessOllama()
	}
	out(m.Report())
}

// out ...
func out(msg string) {
	os.Stdout.Write([]byte(msg))
}

// errExit ...
func errExit(msg string) {
	out(_app + _err + _exit + _space + msg + _linefeed)
	os.Exit(1)
}

// isPipe ...
func isPipe() bool {
	out, _ := os.Stdin.Stat()
	return out.Mode()&os.ModeCharDevice == 0
}

// getPipe ...
func getPipe() string {
	pipe, err := io.ReadAll(os.Stdin)
	if err != nil {
		errExit("reading data from pipe")
	}
	return string(pipe)
}
