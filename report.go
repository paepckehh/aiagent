package aiagent

import (
	"strings"

	lang "github.com/abadojack/whatlanggo"
)

const (
	_err           = " [error]"
	_exit          = " [exit]"
	_dlinefeed     = "\n\n"
	_space         = " "
	_sep           = ';'
	_sep2          = " / "
	_unknownBody   = "unknown -> invalid message body"
	_unknownHeader = "unknown -> invalid message header"
)

// buildReport
func (m *EMail) buildReport() string {
	var s strings.Builder
	s.WriteString(_linefeed)
	s.WriteString("##############################" + _linefeed)
	s.WriteString("# AI MESSAGE ANALYSIS REPORT #" + _linefeed)
	s.WriteString("##############################" + _linefeed)
	s.WriteString("# INBOUND MESSAGE [debug]    : " + _dlinefeed + m.Raw + _dlinefeed)
	s.WriteString("# OFFLINE PREFLIGHT ANALYSIS   " + _linefeed)
	s.WriteString("## Language                  : " + lang.Langs[m.Local.Lang.Lang] + _linefeed)
	s.WriteString("## Confidence                : " + ctoa(m.Local.Lang.Confidence))
	if m.Local.Lang.Confidence < m.Local.TargetLangConfidence {
		s.WriteString(" [exit]")
	} else {
		s.WriteString(" [valid]")
	}
	s.WriteString(_linefeed)
	s.WriteString("## SpellFixes                : " + m.SpellSummary() + _linefeed)
	s.WriteString("## Customer Email            : " + m.Local.Addr.String() + _linefeed)
	s.WriteString("## Customer Email RFC5322    : " + validExit(m.Local.AddrRFC) + _linefeed)
	s.WriteString("## Customer Email Domain MX  : " + validExit(m.Local.AddrMX) + _linefeed)
	s.WriteString("## Customer DB entry         : " + validExit(m.Local.AddrDB) + _linefeed)
	s.WriteString("## Raw / Filtered Characters : " + itoa(m.OpenAI.Raw.Chars) + _sep2 + itoa(m.OpenAI.Msg.Chars) + _linefeed)
	s.WriteString("## Raw / Filtered Words      : " + itoa(m.OpenAI.Raw.Words) + _sep2 + itoa(m.OpenAI.Msg.Words) + _linefeed)
	s.WriteString("## Raw / Filtered GPT3 Token : " + itoa(m.OpenAI.Raw.Token) + _sep2 + itoa(m.OpenAI.Msg.Token) + _linefeed)
	s.WriteString("## Raw / Filtered GPT3 Price : " + price(m.OpenAI.Raw.Cost) + _sep2 + price(m.OpenAI.Msg.Cost) + _linefeed)
	s.WriteString("## Time needed for section   : " + m.Local.ProcessedTime.String() + _linefeed)
	if m.OpenAI.Processed {
		s.WriteString("# ONLINE ML MODULES RESOLVER " + _linefeed)
		s.WriteString("## OpenAI query state        : " + valid(m.OpenAI.State) + _linefeed)
		s.WriteString("## OpenAI says cancel        : " + yesno(m.OpenAI.IsCancel) + _linefeed)
		s.WriteString("## OpenAI analysis [debug]   : " + _dlinefeed + m.OpenAI.Answer + _dlinefeed)
		if len(m.OpenAI.Response) > 10 {
			s.WriteString("## OpenAI Auto Response email: " + _dlinefeed + m.OpenAI.Response + _dlinefeed)
		}
		s.WriteString("## Time needed for section   : " + m.OpenAI.ProcessedTime.String() + _linefeed)
	}
	s.WriteString(_linefeed + _linefeed)
	return s.String()
}
