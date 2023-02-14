# OVERVIEW
[![Go Reference](https://pkg.go.dev/badge/paepcke.de/aicancel.svg)](https://pkg.go.dev/paepcke.de/aicancel) [![Go Report Card](https://goreportcard.com/badge/paepcke.de/aicancel)](https://goreportcard.com/report/paepcke.de/aicancel) [![Go Build](https://github.com/paepckehh/aicancel/actions/workflows/golang.yml/badge.svg)](https://github.com/paepckehh/aicancel/actions/workflows/golang.yml)

[paepche.de/aicancel](https://paepcke.de/aicancel) 

Manages your subscription service in-mailbox (e.g. cancel-requests) via AI (OpenAI ChatGPT Engine).

# FEATURES

* Process requests in 84 languages (offline & online)
* Protect your OpenAI-API Key Budget [$US] with extensive local-first pre-processing and filtering 
* Protect your local infrastructure (DBs) from DoS (spam/targeted-attacks/noise)
	* Filter locally for valid correspondence email addresses (e.g. RFC conformance and validity)
	* Filter locally for supported languages (e.g. do not process emails in Hindi for a German local newspaper subscription)
* EU-GDPR compliant, does not leak personal information (e.g. email address) to cloud-based-AI-backend


# INSTALL

Its a library, you need to customize it for your individual service!

# SHOWCASE INSTALL (example app)

```
go install paepcke.de/aicancel/cmd/aicancel@latest
```

### SHOWCASE DOWNLOAD (prebuild example app)

[github.com/paepckehh/aicancel/releases](https://github.com/paepckehh/aicancel/releases)

# Requirements

* Get your free OpenAI api token here: [OpenAI API key](https://openai.com/api)

# SHOWTIME

* Input eMails: See example messages in root folder!

```Shell 

export OPENAI_API_TOKEN="<your_openai_api_key>"

cat example-email1.txt | aicancel
##############################
# AI MESSAGE ANALYSIS REPORT #
##############################
# OFFLINE PREFLIGHT ANALYSIS
## Language                  : English
## Confidence                : 100% [valid]
## Customer Email            : john.doe@gmail.com
## Customer Email RFC5322    : [valid]
## Customer Email Domain MX  : [valid]
## Customer DB entry         : [valid]
## Time needed for section   : 17.830573ms
# ONLINE ML MODULES RESOLVER 
## OpenAI GPT3 query state   : [valid]
## OpenAI GPT3 says cancel   : [yes]
## OpenAI GPT3 msg [debug]   : Yes, this email does appear to be trying to cancel a subscription service.
## Time needed for section   : 3.532461771s



cat example-email2.txt | aicancel
##############################
# AI MESSAGE ANALYSIS REPORT #
##############################
# OFFLINE PREFLIGHT ANALYSIS
## Language                  : German
## Confidence                : 100% [valid]
## Customer Email            : erika.mustermann@t-online.de
## Customer Email RFC5322    : [valid]
## Customer Email Domain MX  : [valid]
## Customer DB entry         : [valid]
## Time needed for section   : 23.310209ms
# ONLINE ML MODULES RESOLVER 
## OpenAI GPT3 query state   : [valid]
## OpenAI GPT3 says cancel   : [no]
## OpenAI GPT3 msg [debug]   : No, this email does not try to cancel a subscription service. The customer is asking for help in deciding what to do about their subscription and is not explicitly asking for it to be canceled.
## Time needed for section   : 3.943684896s



cat example-email3.txt | aicancel
##############################
# AI MESSAGE ANALYSIS REPORT #
##############################
# OFFLINE PREFLIGHT ANALYSIS
## Language                  : German
## Confidence                : 100% [valid]
## Customer Email            : erika.mustermann@t-online.de
## Customer Email RFC5322    : [valid]
## Customer Email Domain MX  : [valid]
## Customer DB entry         : [valid]
## Time needed for section   : 21.81ms
# ONLINE ML MODULES RESOLVER 
## OpenAI GPT3 query state   : [valid]
## OpenAI GPT3 says cancel   : [yes]
## OpenAI GPT3 msg [debug]   : Yes, this email does try to cancel a subscription service. The author specifically states that they want to end the subscription service due to a negative report in another publication.
## Time needed for section   : 4.099510989s



cat example-email4.txt | aicancel
##############################
# AI MESSAGE ANALYSIS REPORT #
##############################
# OFFLINE PREFLIGHT ANALYSIS
## Language                  : English
## Confidence                : 100% [valid]
## Customer Email            : john.doe@invalid.email
## Customer Email RFC5322    : [valid]
## Customer Email Domain MX  : [valid]
## Customer DB entry         : [failed] [exit]
## Time needed for section   : 15.802395ms



cat example-email5.txt | aicancel
##############################
# AI MESSAGE ANALYSIS REPORT #
##############################
# OFFLINE PREFLIGHT ANALYSIS
## Language                  : Slovene
## Confidence                : 3% [exit]
## Customer Email            : angry.customer@gmail.com
## Customer Email RFC5322    : [valid]
## Customer Email Domain MX  : [valid]
## Customer DB entry         : [failed] [exit]
## Time needed for section   : 24.805052ms


```

# TODO 

Quick Hits & Cost Saver:
- [X] Pre-process messages locally first by analyzing sender email address validity
- [X] Pre-process messages locally first by analyzing and matching message content and language
- [ ] Pre-process and normalize messages locally via spellcheck
- [ ] Pre-process and normalize messages locally via NLP/tokenizer/stemmer to reduce OpenAI token burn rate
- [ ] Add native IMAP/SMTP interfaces to allow total independent email exchange
- [ ] Add individual, language-dependent answer email templates/responses
- [ ] Add local/offline-only/trainable AI models (forward only below a certain local confidence level)
- [ ] Add new online AI APIs as they appear, to save costs and remove service dependency (e.g. Google AI)

Long-term goals (needs commercial project sponsoring):
- [ ] Allow ChatGPT to process customer data change requests (e.g. address, credit card, etc.)
- [ ] Allow ChatGPT to respond, discuss, and clarify corner cases directly with customers via email exchange
- [ ] Add SIP/Voice Interactive Gateway for doing the same via interactive communication

# DOCS

[pkg.go.dev/paepcke.de/aicancel](https://pkg.go.dev/paepcke.de/aicancel)

# TECHNICAL DETAILS

* Yes, large parts of the documentation is generated by AI.
* To be technically correct: backend is OpenAI/GPT3/text-davinici-03 based, not ChatGPT (GPT3.5) 
* This version still features a detailed GPT3 'debug' (answer) message, allowing you to (manually) verify the justification. In high-volume production environments, you are supposed to switch this off and cut the max token load by a significant number.

# CONTRIBUTION

Yes, please! PRs welcome!
