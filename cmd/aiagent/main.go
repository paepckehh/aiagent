package main

import (
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	lang "github.com/abadojack/whatlanggo"
	"paepcke.de/aiagent"
)

const (
	_app               = "[aiagent]"
	_err               = " [error]"
	_exit              = " [exit]"
	_linefeed          = "\n"
	_dlinefeed         = "\n\n"
	_space             = " "
	_sep               = ';'
	_sep2              = " / "
	_unknownBody       = "unknown -> invalid message body"
	_unknownHeader     = "unknown -> invalid message header"
	_langConfidenceMin = 0.70 // lang confidence level required
)

// ...
func main() {
	if !isPipe() {
		errExit("Error:Unable to read from pipe.\nExample: cat message.txt | aiagent ")
	}
	t0 := time.Now()
	m := aiagent.EMail{
		Raw: getPipe(),
	}
	_ = m.SetMessage()
	_ = m.SetLang()
	_ = m.SpellFix()
	_ = m.SetAddr()
	_ = m.Tokenize()
	_ = m.CountToken()
	var s strings.Builder
	s.WriteString(_linefeed)
	s.WriteString("##############################" + _linefeed)
	s.WriteString("# AI MESSAGE ANALYSIS REPORT #" + _linefeed)
	s.WriteString("##############################" + _linefeed)
	s.WriteString("# INBOUND MESSAGE [debug]    : " + _dlinefeed + m.Raw + _dlinefeed)
	s.WriteString("# OFFLINE PREFLIGHT ANALYSIS   " + _linefeed)
	s.WriteString("## Language                  : " + lang.Langs[m.Lang.Lang] + _linefeed)
	s.WriteString("## Confidence                : " + ctoa(m.Lang.Confidence))
	if m.Lang.Confidence < 0.84 {
		s.WriteString(" [exit]")
	} else {
		s.WriteString(" [valid]")
	}
	s.WriteString(_linefeed)
	s.WriteString("## SpellFixes                : " + m.SpellSummary() + _linefeed)
	s.WriteString("## Customer Email            : " + m.Addr.String() + _linefeed)
	s.WriteString("## Customer Email RFC5322    : " + validExit(m.AddrRFC) + _linefeed)
	s.WriteString("## Customer Email Domain MX  : " + validExit(m.AddrMX) + _linefeed)
	s.WriteString("## Customer DB entry         : " + validExit(m.AddrDB) + _linefeed)
	s.WriteString("## Raw / Filtered Characters : " + itoa(m.OpenAI.Raw.Chars) + _sep2 + itoa(m.OpenAI.Msg.Chars) + _linefeed)
	s.WriteString("## Raw / Filtered Words      : " + itoa(m.OpenAI.Raw.Words) + _sep2 + itoa(m.OpenAI.Msg.Words) + _linefeed)
	s.WriteString("## Raw / Filtered GPT3 Token : " + itoa(m.OpenAI.Raw.Token) + _sep2 + itoa(m.OpenAI.Msg.Token) + _linefeed)
	s.WriteString("## Raw / Filtered GPT3 Price : " + price(m.OpenAI.Raw.Cost) + _sep2 + price(m.OpenAI.Msg.Cost) + _linefeed)
	s.WriteString("## Time needed for section   : " + time.Since(t0).String() + _linefeed)
	if m.AddrRFC && m.AddrMX && m.AddrDB && m.Lang.Confidence > 0.75 {
		t0 = time.Now()
		_ = m.SetOpenAI()
		s.WriteString("# ONLINE ML MODULES RESOLVER " + _linefeed)
		s.WriteString("## OpenAI query state        : " + valid(m.OpenAI.State) + _linefeed)
		s.WriteString("## OpenAI says cancel        : " + yesno(m.OpenAI.IsCancel) + _linefeed)
		s.WriteString("## OpenAI analysis [debug]   : " + _dlinefeed + m.OpenAI.Answer + _dlinefeed)
		if len(m.OpenAI.Response) > 10 {
			s.WriteString("## OpenAI Auto Response email: " + _dlinefeed + m.OpenAI.Response + _dlinefeed)
		}
		s.WriteString("## Time needed for section   : " + time.Since(t0).String() + _linefeed)
	}
	s.WriteString(_linefeed + _linefeed)
	out(s.String())
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
