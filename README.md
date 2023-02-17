# OVERVIEW
[![Go Reference](https://pkg.go.dev/badge/paepcke.de/aiagent.svg)](https://pkg.go.dev/paepcke.de/aiagent) [![Go Report Card](https://goreportcard.com/badge/paepcke.de/aiagent)](https://goreportcard.com/report/paepcke.de/aiagent) [![Go Build](https://github.com/paepckehh/aiagent/actions/workflows/golang.yml/badge.svg)](https://github.com/paepckehh/aiagent/actions/workflows/golang.yml)

[paepche.de/aiagent](https://paepcke.de/aiagent) 

Manages your subscription service in-mailbox (e.g. cancel-requests) via AI (OpenAI ChatGPT Engine).

# FEATURES

* Process requests in [84](https://github.com/abadojack/whatlanggo/blob/master/SUPPORTED_LANGUAGES.md#supported-languages) languages (offline & online)
* Protect your OpenAI-API Key Budget [$US] with extensive local-first pre-processing and filtering 
* Protect your local infrastructure (DBs) from DoS (spam/targeted-attacks/noise)
	* Filter locally for valid correspondence email addresses (e.g. RFC conformance and validity)
	* Filter locally for supported languages (e.g. do not process emails in Hindi for a German local newspaper subscription)
* EU-GDPR compliant, does not leak personal information (e.g. email address) to cloud-based-AI-backend


# INSTALL

Its a library, you need to customize it for your individual service!

# SHOWCASE INSTALL (example app)

```
go install paepcke.de/aiagent/cmd/aiagent@latest
```

### SHOWCASE DOWNLOAD (prebuild example app)

[github.com/paepckehh/aiagent/releases](https://github.com/paepckehh/aiagent/releases)

# Requirements

* Get your free OpenAI api token here: [OpenAI API key](https://openai.com/api)

# SHOWTIME

* Input eMails: See example messages in root folder!

```Shell 

export OPENAI_API_TOKEN="<your_openai_api_key>"

> cat example-email1.txt | aiagent
##############################
# AI MESSAGE ANALYSIS REPORT #
##############################
# INBOUND MESSAGE [debug]    : 

From:    john.doe@gmail.com
To:      contract@service.com
Subject: I`m done! 

Hello MyService Corp,

I changed my mind, my life has changed. I dont want any kind of
subscription based service, not from you or anyone else! 

Please cancel all services and refund. 

Greetings,
John


# OFFLINE PREFLIGHT ANALYSIS   
## Language                  : English
## Confidence                : 100% [valid]
## SpellFixes                : [none]
## Customer Email            : john.doe@gmail.com
## Customer Email RFC5322    : [valid]
## Customer Email Domain MX  : [valid]
## Customer DB entry         : [valid]
## Raw / Filtered Characters : 283 / 225
## Raw / Filtered Words      : 41 / 37
## Raw / Filtered GPT3 Token : 84 / 59
## Raw / Filtered GPT3 Price : 0.00168 US$  / 0.00118 US$ 
## Time needed for section   : 3.901066719s
# ONLINE ML MODULES RESOLVER 
## OpenAI query state        : [valid]
## OpenAI says cancel        : [yes]
## OpenAI analysis [debug]   : 

Yes, this email appears to be attempting to cancel a subscription service.

## OpenAI Auto Response email: 

Dear John,

We are sorry to hear that you no longer wish to use our subscription service. 
We understand that life changes and will accept your request to cancel. 

We are sad to see you go, and would like to ask if there is anything we can 
do to keep you as a subscriber? If there is, please let us know, and we will 
do our best to accommodate your request. 

If you have any further questions or concerns, please do not hesitate to contact us.

Thank you for being part of MyService Corp.

Sincerely,
Michael

## Time needed for section   : 13.471735885s



> cat example-email2.txt | aiagent
##############################
# AI MESSAGE ANALYSIS REPORT #
##############################
# INBOUND MESSAGE [debug]    : 

From:    erika.mustermann@t-online.de
To:      vertrag@service.com
Subject: Mein Gala Abo

Hallo Kundenservice, 
ich war letzte Woche bei Tante Erna in Wuppertal zum 93. Geburtstag eingeladen.
Und da meinte mein Enkel nach der 2. Tasse Kaffee ich sollte endlich mein Gala
Abo abbestellen, da ich sowieso die kleinen Buchstaben nicht mehr lesen kann.

Ich will diese Zeitschrift sowieso nicht mehr seit diesem negativen Bericht in
der Apotheken-Rundschau. Aber manchmal haben die einfach gute Berichte. Was soll
ich jetzt tun?

Erika


# OFFLINE PREFLIGHT ANALYSIS   
## Language                  : German
## Confidence                : 100% [valid]
## SpellFixes                : [none]
## Customer Email            : erika.mustermann@t-online.de
## Customer Email RFC5322    : [valid]
## Customer Email Domain MX  : [valid]
## Customer DB entry         : [valid]
## Raw / Filtered Characters : 534 / 467
## Raw / Filtered Words      : 77 / 73
## Raw / Filtered GPT3 Token : 202 / 172
## Raw / Filtered GPT3 Price : 0.00404 US$  / 0.00344 US$ 
## Time needed for section   : 3.697031042s
# ONLINE ML MODULES RESOLVER 
## OpenAI query state        : [valid]
## OpenAI says cancel        : [no]
## OpenAI analysis [debug]   : 

No, this email does not try to cancel a subscription service. The customer is asking 
what they should do, but they have not made a decision either way.

## Time needed for section   : 5.218959115s



> cat example-email3.txt | aiagent
##############################
# AI MESSAGE ANALYSIS REPORT #
##############################
# INBOUND MESSAGE [debug]    : 

From:    erika.mustermann@t-online.de
To:      vertrag@service.com
Subject: Gala Abo kuendigung

Hallo Kundenservice, 
ich war letzte Woche bei Tante Erna in Wuppertal zum 93. Geburtstag eingeladen.
Und da meinte mein Enkel nach der 2. Tasse Kaffee ich sollte endlich mein Gala
Abo abbestellen, da ich sowieso die kleinen Buchstaben nicht mehr lesen kann.

Ich will diese Zeitschrift sowieso nicht mehr seit diesem negativen Bericht in
der Apotheken-Rundschau. Bitte beenden sie sofort (!) mein Gala Abo. Sofort!

Erika


# OFFLINE PREFLIGHT ANALYSIS   
## Language                  : German
## Confidence                : 100% [valid]
## SpellFixes                : [none]
## Customer Email            : erika.mustermann@t-online.de
## Customer Email RFC5322    : [valid]
## Customer Email Domain MX  : [valid]
## Customer DB entry         : [valid]
## Raw / Filtered Characters : 521 / 454
## Raw / Filtered Words      : 74 / 70
## Raw / Filtered GPT3 Token : 201 / 171
## Raw / Filtered GPT3 Price : 0.00402 US$  / 0.00342 US$ 
## Time needed for section   : 3.712945625s
# ONLINE ML MODULES RESOLVER 
## OpenAI query state        : [valid]
## OpenAI says cancel        : [yes]
## OpenAI analysis [debug]   : 

Yes, this email tries to cancel a subscription service. The author mentions 
they want to cancel their Gala Abo and makes a request to have it done immediately.

## OpenAI Auto Response email: 

Liebe Frau Erika,

vielen Dank für Ihre Email. Wir bedauern sehr, dass Sie Ihr Gala Abo kündigen möchten. 
Wir haben Ihre Kündigung erhalten und werden diese umgehend bearbeiten. 

Wir verstehen, dass Sie nicht mehr zufrieden sind und würden gerne wissen, ob es etwas 
gibt, das wir für Sie tun können, um Ihnen zu helfen, Ihr Abo zu behalten. 

Bitte lassen Sie es uns wissen, wenn Sie Fragen oder Anregungen zu Ihrem Abonnement haben. 
Wir werden unser Bestes tun, um Ihnen zu helfen.

Mit freundlichen Grüßen, 

Michael

## Time needed for section   : 21.037154219s



> cat example-email4.txt | aiagent
##############################
# AI MESSAGE ANALYSIS REPORT #
##############################
# INBOUND MESSAGE [debug]    : 

From:    john.doe@invalid.email
To:      subscription@service.com
Subject: I`m done! 

Hello MyService Corp,
Please cancel all services and refund. 

Greetings,
John


# OFFLINE PREFLIGHT ANALYSIS   
## Language                  : English
## Confidence                : 100% [valid]
## SpellFixes                : [none]
## Customer Email            : john.doe@invalid.email
## Customer Email RFC5322    : [valid]
## Customer Email Domain MX  : [valid]
## Customer DB entry         : [failed] [exit]
## Raw / Filtered Characters : 167 / 101
## Raw / Filtered Words      : 18 / 14
## Raw / Filtered GPT3 Token : 54 / 28
## Raw / Filtered GPT3 Price : 0.00108 US$  / 0.00056 US$ 
## Time needed for section   : 3.924910468s



> cat example-email5.txt | aiagent
##############################
# AI MESSAGE ANALYSIS REPORT #
##############################
# INBOUND MESSAGE [debug]    : 

From:    angry.customer@gmail.com
To:      contract@service.com
Subject: vmflmvopfmeopvvfkmvkfmvkefnkvneVEF5-5j4f

WUpkvlUdtUla3QlVUs86/8tQ+1GlImhlcvKH6tCNK+UQQ+tGiP1hPOlFFe6N5Yjf5L/o9OLenFo6sP
ZamtupDgJGEhWtdQPA+LP+f1LFAynpQbgmWBKmgDUCY8IOS+mxEeWYKnKsvvkALgo4kvzhmlox2YSC
gDdFEPil1IRIRv72anrvtTBmGP+VBldGK8/+1ZTcz/vIJ5G9gJOILzm2yGiXqy06lcn/Lch9i1o+nS
1BvjpiE6Ij0DNoKaKXZx9z3W3ZNVf8CU35avvKR6YES5OIlNxz/XwejZuP1dDQD2eDJrRwpUhMoC/A
xeqc0SLfLT1RXCLWv9Se5DWrswvo4Y9bkS24xQYOlpgz1We2KDyH/tiyDwtimoePmaI4A7x2mtazMo


# OFFLINE PREFLIGHT ANALYSIS   
## Language                  : Italian
## Confidence                : 10% [exit]
## SpellFixes                : [none]
## Customer Email            : angry.customer@gmail.com
## Customer Email RFC5322    : [valid]
## Customer Email Domain MX  : [valid]
## Customer DB entry         : [failed] [exit]
## Raw / Filtered Characters : 511 / 447
## Raw / Filtered Words      : 11 / 7
## Raw / Filtered GPT3 Token : 353 / 328
## Raw / Filtered GPT3 Price : 0.00706 US$  / 0.00656 US$ 
## Time needed for section   : 3.693537708s



> cat example-email6.txt | aiagent
##############################
# AI MESSAGE ANALYSIS REPORT #
##############################
# INBOUND MESSAGE [debug]    : 

From:    angry.customer@gmail.com
To:      contract@service.com
Subject: I'm very angry!

I`m so angry that I`m very bad at spell the words, 
so you will have a hard time to read my email.

Please do me a favour and honour my explict wish to 
cancel my subscriptons. NOW! Or I post on faceboook 
about you hypocrite!

Angry Customer 


# OFFLINE PREFLIGHT ANALYSIS   
## Language                  : English
## Confidence                : 100% [valid]
## SpellFixes                : explicit,subscriptions,facebook
## Customer Email            : angry.customer@gmail.com
## Customer Email RFC5322    : [valid]
## Customer Email Domain MX  : [valid]
## Customer DB entry         : [failed] [exit]
## Raw / Filtered Characters : 335 / 272
## Raw / Filtered Words      : 55 / 51
## Raw / Filtered GPT3 Token : 103 / 74
## Raw / Filtered GPT3 Price : 0.00206 US$  / 0.00148 US$ 
## Time needed for section   : 3.86417427s

```

# TODO

Pre-process messages offline 
- [X] analyzing sender email address validity
- [X] analyzing and match message content and language
- [X] fix spelling errors
- [X] detect and isolatlate locally, to save token payload, protect privacy
	- [ ] detect additional customer provided email addresses in message body and remove 
	- [ ] detect any customer provided urls in message body and remove 
	- [ ] detect any format of custom phone number format and remove 
- [ ] Preprocess via NLP/tokenizer/stemmer to reduce OpenAI token burn rate
- [ ] Add local/offline-only/trainable AI models (forward only below a certain local confidence level)

Interfaces 
- [ ] Add native IMAP/SMTP interfaces to allow total independent email exchange
- [ ] Add individual, language-dependent answer email templates/responses
- [ ] Add new online AI APIs as they appear, to save costs and remove service dependency (e.g. Google AI)

Long-term goals (needs commercial project sponsoring):
- [ ] Allow ChatGPT to process customer data change requests (e.g. address, credit card, etc.)
- [ ] Allow ChatGPT to respond, discuss, and clarify corner cases directly with customers via email exchange
- [ ] Add SIP/Voice Interactive Gateway for doing the same via interactive communication

# DOCS

[pkg.go.dev/paepcke.de/aiagent](https://pkg.go.dev/paepcke.de/aiagent)

# CONTRIBUTION

Yes, please! PRs welcome!





