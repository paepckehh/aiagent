// aiagent provides an ai based service to process subscriton (cancel) based emails
package aiagent

import (
	"context"
	"errors"
	"strings"

	lang "github.com/abadojack/whatlanggo"
	spell "github.com/golangci/misspell"
	addr "github.com/mcnijman/go-emailaddress"
	gpt3 "github.com/sashabaranov/go-gpt3"

	"paepcke.de/aiagent/gpt3encoder"
	"paepcke.de/dnscache"
)

// GPT3 AI Model defaults
const (
	GPT3AIModel         string  = gpt3.GPT3TextDavinci003
	GPT3AICosts         float64 = 0.0200 / 1000
	GPT3AIPromtCancel   string  = "Does the following email try to cancel a subscription service?"
	GPT3AIPromtResponse string  = "I'm Callcenter Agent Michael. The following customer email tries to canel a subscription service and maybe contains some reasons why. Please answer this email. Tell this customer that you accept the canel request, but you are sad to see him leave. Ask the customer if there is anything you can do, to keep this subscription. If the customer provides reasons for the cancel request, try to argue politly. Write this answer email in language "
)

// Payload holds the metrics for a payload
type Payload struct {
	Chars int     // number of characters
	Words int     // number of words
	Token int     // number of tokens
	Cost  float64 // cost in $US for one token
}

// OpenAI holds the OpenAI.org interface
type OpenAI struct {
	State    bool
	IsCancel bool
	Answer   string
	Response string
	Raw      Payload
	Msg      Payload
}

// EMail holds the Raw message and all parsable attributes
type EMail struct {
	Raw        string            // unprocessed raw input text
	Message    string            // message
	Lang       lang.Info         // message body language and confidence level
	Addr       addr.EmailAddress // verified customer email
	AddrRFC    bool              // Addr is RFC5322 conform and has ICANN TLD
	AddrMX     bool              // Addr Domain has valid MX record
	AddrDB     bool              // Addr match found in mandant DB
	OpenAI     OpenAI            // OpenAI response
	Supported  []string          // Supported
	SpellFixed []spell.Diff      // SpellFixed words
}

// SetOpenAI parses the message body via OpenAI/GPT3
func (m *EMail) SetOpenAI() error {
	if m.Message == "" || len(m.Message) < 10 {
		return errors.New("Message is empty or too small. Unable to process.")
	}
	if m.Lang.Lang < 1 {
		return errors.New("Message language is unknown. Unable to process.")
	}
	token, ok := getEnv("OPENAI_API_TOKEN")
	if !ok {
		m.OpenAI.Answer = bracket("unable to read API token from env")
		return errors.New(m.OpenAI.Answer)
	}
	c := gpt3.NewClient(token)
	ctx := context.Background()
	req := gpt3.CompletionRequest{
		Model:       GPT3AIModel,
		MaxTokens:   64,
		Prompt:      GPT3AIPromtCancel + _linefeed + m.Message,
		Temperature: 0,
	}
	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		m.OpenAI.Answer = bracket(err.Error())
		return err
	}
	m.OpenAI.State = true
	m.OpenAI.Answer = resp.Choices[0].Text[1:]
	if len(m.OpenAI.Answer) > 2 {
		if m.OpenAI.Answer[:3] == "Yes" {
			m.OpenAI.IsCancel = true
			c := gpt3.NewClient(token)
			ctx := context.Background()
			req := gpt3.CompletionRequest{
				Model:       GPT3AIModel,
				MaxTokens:   300,
				Prompt:      GPT3AIPromtResponse + lang.Langs[m.Lang.Lang] + _dot + _linefeed + m.Message,
				Temperature: 0.6,
			}
			resp, err := c.CreateCompletion(ctx, req)
			if err != nil {
				m.OpenAI.Response = bracket(err.Error())
				return err
			}
			m.OpenAI.Response = resp.Choices[0].Text[1:]
		}
	}
	return nil
}

// Tokenize parses (offline) the message body and replaces words via stemmer token
func (m *EMail) Tokenize() error {
	return nil
}

// CountToken counts the real GT3 compatible Token for Raw and Message, calculate the costs in US$
func (m *EMail) CountToken() error {
	_ = m.OpenAI.Raw.Calc(m.Raw)
	_ = m.OpenAI.Msg.Calc(m.Message)
	return nil
}

// Calc Payload
func (p *Payload) Calc(msg string) error {
	p.Chars = strings.Count(msg, "")
	p.Words = len(strings.Fields(msg))
	t, err := gpt3encoder.NewEncoder()
	if err != nil {
		return errors.New("[gpt3encoder] unrecoverable error")
	}
	enc, err := t.Encode(msg)
	if err != nil {
		return errors.New("[gpt3encoder] unrecoverable error")
	}
	p.Token = len(enc)
	p.Cost = float64(p.Token) * GPT3AICosts
	return nil
}

// SpellFix parses (offline) the message body and replaces words via stemmer token
func (m *EMail) SpellFix() error {
	switch lang.Langs[m.Lang.Lang] {
	case "English":
		r := spell.New()
		m.Message, m.SpellFixed = r.Replace(m.Message)
	}
	return nil
}

// SpellSummary provides a list of spell-fixed words as string from from SpellCheck diff
func (m *EMail) SpellSummary() string {
	l := len(m.SpellFixed)
	if l < 1 {
		return "[none]"
	}
	result := make([]string, l)
	for v, diff := range m.SpellFixed {
		result[v] = diff.Corrected
	}
	return strings.Join(result, ",")
}

// SetAddr parses and validates sender address
func (m *EMail) SetAddr() error {
	l, idx := len(m.Raw), 0
	if l > 256 { // do no process first line with more than first 256 characters
		l = 256
	}
	for {
		if idx > l {
			return errors.New("No Sender Address: Input Message to Short")
		}
		if m.Raw[idx] == _linefeedR {
			break
		}
		idx++
	}
	line := m.Raw[:idx]
	emails := addr.FindWithIcannSuffix([]byte(line), false)
	if len(emails) > 0 {
		m.AddrRFC, m.AddrMX = true, true
		m.Addr = *emails[0]
		if _, err := dnscache.LookupMX(m.Addr.Domain); err != nil {
			m.AddrMX = false
		}
	}
	if strings.Contains(db, m.Addr.String()) {
		m.AddrDB = true
	}
	return nil
}

// SetLang validates senders message language
func (m *EMail) SetLang() error {
	m.Lang = lang.Detect(m.Message)
	return nil
}

// SetMessage parses message body
func (m *EMail) SetMessage() error {
	l, idx, header := len(m.Raw), 0, 0
	for header < _msgHeader {
		if idx > l {
			return errors.New("Input Message to Short")
		}
		if m.Raw[idx] == _linefeedR {
			header++
		}
		idx++
	}
	m.Message = m.Raw[idx:]
	return nil
}
