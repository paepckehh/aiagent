package aiagent

import (
	"strings"

	lang "github.com/abadojack/whatlanggo"
)

// buildReport
func (m *EMail) buildReport() string {
	var s strings.Builder
	s.WriteString("####################################################" + _linefeed)
	s.WriteString("####################################################" + _linefeed)
	s.WriteString("#          -= AI MESSAGE ANALYSIS REPORT =-        #" + _linefeed)
	s.WriteString("####################################################" + _linefeed)
	s.WriteString("####################################################" + _linefeed)
	s.WriteString("# INBOUND MESSAGE [debug]    : " + _dlinefeed + m.Raw + _dlinefeed)
	s.WriteString("# OFFLINE PREFLIGHT ANALYSIS   " + _linefeed)
	s.WriteString("## Language                  : " + lang.Langs[m.Local.Lang.Lang] + _linefeed)
	s.WriteString("## Confidence                : " + ctoa(m) + _linefeed)
	s.WriteString("## SpellFixes                : " + m.SpellSummary() + _linefeed)
	s.WriteString("## Customer Email            : " + m.Local.Addr.String() + _linefeed)
	s.WriteString("## Customer Email RFC5322    : " + validExit(m.Local.AddrRFC) + _linefeed)
	s.WriteString("## Customer Email Domain MX  : " + validExit(m.Local.AddrMX) + _linefeed)
	s.WriteString("## Customer DB entry         : " + validExit(m.Local.AddrDB) + _linefeed)
	s.WriteString("## Anonymized EMailAddresses : " + isArray(m.Privacy.EMails) + _linefeed)
	s.WriteString("## Anonymized URLs           : " + isArray(m.Privacy.URLs) + _linefeed)
	s.WriteString("## Anonymized PhoneNumbers   : " + isArray(m.Privacy.Phones) + _linefeed)
	s.WriteString("## Raw / Filtered Characters : " + itoa(m.AiModel.Raw.Chars) + _sep2 + itoa(m.AiModel.Msg.Chars) + _linefeed)
	s.WriteString("## Raw / Filtered Words      : " + itoa(m.AiModel.Raw.Words) + _sep2 + itoa(m.AiModel.Msg.Words) + _linefeed)
	s.WriteString("## Raw / Filtered GPT3 Token : " + itoa(m.AiModel.Raw.Token) + _sep2 + itoa(m.AiModel.Msg.Token) + _linefeed)
	s.WriteString("## Raw / Filtered GPT3 Price : " + price(m.AiModel.Raw.Cost) + _sep2 + price(m.AiModel.Msg.Cost) + _linefeed)
	if m.AiModel.Processed {
		s.WriteString("## Time needed for section   : " + m.Local.ProcessedTime.String() + _dlinefeed)
		s.WriteString("# ONLINE ML MODULES RESOLVER " + _linefeed)
		s.WriteString("## AiModel query state        : " + valid(m.AiModel.State) + _linefeed)
		s.WriteString("## AiModel says cancel        : " + yesno(m.AiModel.IsCancel) + _linefeed)
		s.WriteString("## AiModel analysis [debug]   : " + _dlinefeed + m.AiModel.Answer + _dlinefeed)
		if len(m.AiModel.Response) > 10 {
			s.WriteString("## AiModel Auto Response email: " + _dlinefeed + m.AiModel.Response + _dlinefeed)
		}
		s.WriteString("## Time needed for section   : " + m.AiModel.ProcessedTime.String() + _linefeed)
	} else {
		s.WriteString("## No AiModel query performed, inbound data quality failed." + _linefeed)
		s.WriteString("## Time needed for section   : " + m.Local.ProcessedTime.String() + _linefeed)
	}
	s.WriteString(_dlinefeed)
	return s.String()
}
