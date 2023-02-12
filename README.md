# OVERVIEW
[![Go Reference](https://pkg.go.dev/badge/paepcke.de/aicancel.svg)](https://pkg.go.dev/paepcke.de/aicancel) [![Go Report Card](https://goreportcard.com/badge/paepcke.de/aicancel)](https://goreportcard.com/report/paepcke.de/aicancel) [![Go Build](https://github.com/paepckehh/aicancel/actions/workflows/golang.yml/badge.svg)](https://github.com/paepckehh/aicancel/actions/workflows/golang.yml)

[paepche.de/aicancel](https://paepcke.de/aicancel) 

Manages your subscription service in-mailbox (eg. cancel-requests) via AI (OpenAI ChatGPT Engine).

# Features

* Process automaticly requests ( [84](https://github.com/abadojack/whatlanggo/blob/master/SUPPORTED_LANGUAGES.md#supported-languages) languages )
* Protect your OpenAI-API Key Budget [$US] by extensive local-first pre-processing & filtering 
* Protect your local Infrastruture (Mandant Database Requests) from DoS (spam/targeted-attacs/noise)
	* Filter locally for valid correspondence email addresses (eg. RFC conformance and validity)
	* Filter locally for supported languages (eg. do not process emails in hindi for a german local newspaper subscription)
	
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
## Confidence                : 100%
## Customer Email            : john.doe@gmail.com
## Customer Email RFC5322    : [valid]
## Customer Email Domain MX  : [valid]
## Customer DB entry         : [valid]
## Time needed for section   : 19.251771ms
# ONLINE ML MODULES RESOLVER 
## OpenAI GPT3 query state   : [valid]
## OpenAI GPT3 says cancel   : [yes]
## OpenAI GPT3 msg [debug]   : Yes, this email does appear to be trying to cancel a subscription service.
## Time needed for section   : 3.423134375s

cat example-email2.txt | aicancel
##############################
# AI MESSAGE ANALYSIS REPORT #
##############################
# OFFLINE PREFLIGHT ANALYSIS
## Language                  : German
## Confidence                : 100%
## Customer Email            : erika.mustermann@t-online.de
## Customer Email RFC5322    : [valid]
## Customer Email Domain MX  : [valid]
## Customer DB entry         : [valid]
## Time needed for section   : 24.521666ms
# ONLINE ML MODULES RESOLVER 
## OpenAI GPT3 query state   : [valid]
## OpenAI GPT3 says cancel   : [no]
## OpenAI GPT3 msg [debug]   : No, this email does not really try to cancel a subscription service. It is a question from a customer asking what to do about their subscription.
## Time needed for section   : 4.450808021s

cat example-email3.txt | aicancel
##############################
# AI MESSAGE ANALYSIS REPORT #
##############################
# OFFLINE PREFLIGHT ANALYSIS
## Language                  : German
## Confidence                : 100%
## Customer Email            : erika.mustermann@t-online.de
## Customer Email RFC5322    : [valid]
## Customer Email Domain MX  : [valid]
## Customer DB entry         : [valid]
## Time needed for section   : 22.706354ms
# ONLINE ML MODULES RESOLVER 
## OpenAI GPT3 query state   : [valid]
## OpenAI GPT3 says cancel   : [yes]
## OpenAI GPT3 msg [debug]   : Yes, this email is trying to cancel a subscription service.
## Time needed for section   : 3.285378802s

cat example-email4.txt | aicancel
##############################
# AI MESSAGE ANALYSIS REPORT #
##############################
# OFFLINE PREFLIGHT ANALYSIS
## Language                  : English
## Confidence                : 100%
## Customer Email            : john.doe@invalid.email
## Customer Email RFC5322    : [valid]
## Customer Email Domain MX  : [valid]
## Customer DB entry         : [failed] [exit]
## Time needed for section   : 17.033072ms

cat example-email5.txt | aicancel
##############################
# AI MESSAGE ANALYSIS REPORT #
##############################
# OFFLINE PREFLIGHT ANALYSIS
## Language                  : Slovene
## Confidence                : 3%
## Customer Email            : angry.customer@gmail.com
## Customer Email RFC5322    : [valid]
## Customer Email Domain MX  : [valid]
## Customer DB entry         : [failed] [exit]
## Time needed for section   : 25.026875ms


```

# TODO 

Archive further cost savings - via:
* [] pre-process messages locally via NLP/tokenizer to reduce OpenAI token burn rate
* [] add new Online-AI APIs as they appear, to save costs and remove service dependecy (eg. google-ai)
* [] add local/offline-only/train-able AI-Models (forward only below a certain local confidence level
* [] add interfaces to let discuss and clarify ChatGPT corner cases with the customer 
* [] add individual, defined answer email templates
* [] add native IMAP/SMTP Interfaces for mail eXchange
* [] add SIP/Voice Interactive Gateway for Interactive Communication with the Customer

# DOCS

[pkg.go.dev/paepcke.de/aicancel](https://pkg.go.dev/paepcke.de/aicancel)

# TECHNICAL DETAILS

* To be technically correct: backend is OpenAI/GPT3/text-davinici-03 based not, ChatGPT(GPT3.5) 

# CONTRIBUTION

Yes, Please! PRs Welcome! 



