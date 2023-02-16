// aicancel provides an ai based service to process subscriton (cancel) based emails
package aicancel

import (
	"context"
	_ "embed"
	"errors"
	"strings"
	"syscall"

	lang "github.com/abadojack/whatlanggo"
	spell "github.com/golangci/misspell"
	addr "github.com/mcnijman/go-emailaddress"
	gpt3 "github.com/sashabaranov/go-gpt3"
	"paepcke.de/dnscache"
)

const (
	_linefeedR = '\n' // linefeed
	_msgHeader = 2    // message header line
)

//go:embed mandant.db
var db string

// OpenAI
type OpenAI struct {
	State    bool
	IsCancel bool
	Message  string
}

// EMail holds the Raw message and all parsable attributes
type EMail struct {
	Raw        string            // unprocessed raw input text
	Body       string            // message body
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
	if m.Body == "" {
		if err := m.SetBody(); err != nil {
			return errors.New("Unable to set Message Body: " + err.Error())
		}
	}
	token, ok := getEnv("OPENAI_API_TOKEN")
	if !ok {
		m.OpenAI.Message = bracket("unable to read API token from env")
		return errors.New(m.OpenAI.Message)
	}
	c := gpt3.NewClient(token)
	ctx := context.Background()
	req := gpt3.CompletionRequest{
		Model:       gpt3.GPT3TextDavinci003,
		MaxTokens:   64,
		Prompt:      "Does the following email really try to cancel a subscription service?\n" + m.Body,
		Temperature: 0,
	}
	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		m.OpenAI.Message = bracket(err.Error())
		return err
	}
	m.OpenAI.State = true
	m.OpenAI.Message = resp.Choices[0].Text[1:]
	if len(m.OpenAI.Message) > 2 {
		if m.OpenAI.Message[:3] == "Yes" {
			m.OpenAI.IsCancel = true
		}
	}
	return nil
}

// Tokenize parses (offline) the message body and replaces words via stemmer token
func (m *EMail) Tokenize() error {
	if m.Body == "" {
		if err := m.SetBody(); err != nil {
			return errors.New("Unable to set Message Body: " + err.Error())
		}
	}
	return nil
}

// SpellFix parses (offline) the message body and replaces words via stemmer token
func (m *EMail) SpellFix() error {
	if m.Body == "" {
		if err := m.SetBody(); err != nil {
			return errors.New("Unable to set Message Body: " + err.Error())
		}
	}
	switch lang.Langs[m.Lang.Lang] {
	case "English":
		r := spell.New()
		m.Body, m.SpellFixed = r.Replace(m.Body)
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

// SetLang validates sender email address
func (m *EMail) SetLang() error {
	if m.Body == "" {
		if err := m.SetBody(); err != nil {
			return errors.New("Unable to set Message Body: " + err.Error())
		}
	}
	m.Lang = lang.Detect(m.Body)
	return nil
}

// SetBody parses message bodu
func (m *EMail) SetBody() error {
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
	m.Body = m.Raw[idx:]
	return nil
}

// getEnv ...
func getEnv(in string) (string, bool) {
	return syscall.Getenv(in)
}

// bracket
func bracket(in string) string {
	return "[" + in + "]"
}
