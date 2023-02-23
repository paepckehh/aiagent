# OVERVIEW
[![Go Reference](https://pkg.go.dev/badge/paepcke.de/aiagent.svg)](https://pkg.go.dev/paepcke.de/aiagent) [![Go Report Card](https://goreportcard.com/badge/paepcke.de/aiagent)](https://goreportcard.com/report/paepcke.de/aiagent) [![Go Build](https://github.com/paepckehh/aiagent/actions/workflows/golang.yml/badge.svg)](https://github.com/paepckehh/aiagent/actions/workflows/golang.yml)


![ai_generated_ai_callcenter_agent](https://github.com/paepckehh/paepckehh/raw/main/logos/aiagent.png)

[paepcke.de/aiagent](https://paepcke.de/aiagent) 

Manages your subscription service in-mailbox (e.g. cancel-requests) via AI (OpenAI ChatGPT Engine).

# FEATURES

* Process requests in [84](https://github.com/abadojack/whatlanggo/blob/master/SUPPORTED_LANGUAGES.md#supported-languages) languages (offline & online)
* Protect your OpenAI-API Key Budget [$US] with extensive local-first pre-processing and filtering 
* Protect your local infrastructure (DBs) from DoS (spam/targeted-attacks/noise)
	* Filter locally for valid correspondence email addresses (e.g. RFC conformance and validity)
	* Filter locally for supported languages (e.g. do not process emails in Hindi for a German local newspaper subscription)
* EU-GDPR compliant (WIP, see todo)


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

```
export OPENAI_API_TOKEN="<your_openai_api_key>"

cat ../../example-email1.txt | ./aiagent
####################################################
####################################################
#          -= AI MESSAGE ANALYSIS REPORT =-        #
####################################################
####################################################
# INBOUND MESSAGE [debug]    : 

From:    john.doe@gmail.com
To:      contract@service.com
Subject: I`m done! 

Hello MyService Corp,

I changed my mind, my life has changed. 

I dont want any kind of subscription 
service, not from you or anyone else! 

Please cancel all subscriptions.

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
## Anonymized EMailAddresses : [none]
## Anonymized URLs           : [none]
## Anonymized PhoneNumbers   : [none]
## Raw / Filtered Characters : 273 / 215
## Raw / Filtered Words      : 38 / 34
## Raw / Filtered GPT3 Token : 82 / 57
## Raw / Filtered GPT3 Price : 0.00164 US$  / 0.00114 US$ 
## Time needed for section   : 4.277123229s

# ONLINE ML MODULES RESOLVER 
## OpenAI query state        : [valid]
## OpenAI says cancel        : [yes]
## OpenAI analysis [debug]   : 

Yes, this email written attempt to cancel a subscription service.

## OpenAI Auto Response email: 

Dear John,

We are sorry to hear that you have decided to cancel your subscription 
with us. We understand that life can change and we accept your request 
to cancel your subscription. We will cancel it as soon as possible. 

We are sad to see you go and would like to ask if there is anything we 
can do to keep your subscription? We would love to hear your feedback 
and see if there is anything we can do to improve our service. 

Thank you for your time and we hope to hear from you soon. 

Sincerely,
MyService Corp

## Time needed for section   : 13.079184167s


cat ../../example-email2.txt | ./aiagent
####################################################
####################################################
#          -= AI MESSAGE ANALYSIS REPORT =-        #
####################################################
####################################################
# INBOUND MESSAGE [debug]    : 

From:    erika.mustermann@t-online.de
To:      vertrag@service.com
Subject: Gala Abo kuendigung

Hallo Kundenservice, 
ich war letzte Woche bei Tante Erna in Wuppertal
zum 93. Geburtstag eingeladen. Und da meinte mein
Enkel bereits nach der ersten Tasse Kaffee ich
sollte endlich mein Gala Abo abbestellen, da ich
die kleinen Buchstaben nicht mehr lesen kann.

Ich will diese Zeitschrift sowieso nicht mehr seit
diesem negativen Bericht in der Apotheken-Rundschau.

Bitte beenden sie sofort (!) mein Gala Abo. SOFORT!

Erika


# OFFLINE PREFLIGHT ANALYSIS   
## Language                  : German
## Confidence                : 100% [valid]
## SpellFixes                : [none]
## Customer Email            : erika.mustermann@t-online.de
## Customer Email RFC5322    : [valid]
## Customer Email Domain MX  : [valid]
## Customer DB entry         : [valid]
## Anonymized EMailAddresses : [none]
## Anonymized URLs           : [none]
## Anonymized PhoneNumbers   : [none]
## Raw / Filtered Characters : 526 / 459
## Raw / Filtered Words      : 74 / 70
## Raw / Filtered GPT3 Token : 202 / 172
## Raw / Filtered GPT3 Price : 0.00404 US$  / 0.00344 US$ 
## Time needed for section   : 4.013430521s

# ONLINE ML MODULES RESOLVER 
## OpenAI query state        : [valid]
## OpenAI says cancel        : [yes]
## OpenAI analysis [debug]   : 

Yes, this email written attempts to cancel a subscription service.

## OpenAI Auto Response email: 

Sehr geehrte Frau Erika,

vielen Dank für Ihre E-Mail. Wir bedauern sehr, dass Sie Ihr Gala Abo 
kündigen möchten. Wir akzeptieren Ihre Kündigung und werden es zum 
nächstmöglichen Zeitpunkt beenden. Wir sind traurig, Sie als Kundin 
zu verlieren. Gibt es irgendetwas, das wir tun können, um Sie davon 
abzuhalten, Ihr Abo zu kündigen? Wir würden uns freuen, wenn wir 
Ihnen weiterhelfen könnten.

Mit freundlichen Grüßen,
[Name]

## Time needed for section   : 16.498451198s


cat ../../example-email3.txt | ./aiagent
####################################################
####################################################
#          -= AI MESSAGE ANALYSIS REPORT =-        #
####################################################
####################################################
# INBOUND MESSAGE [debug]    : 

From:    erika.mustermann@t-online.de
To:      vertrag@service.com
Subject: Mein Gala Abo

Hallo Kundenservice, 
ich war letzte Woche bei Tante Erna in Wuppertal
zum 93. Geburtstag eingeladen. Und da meinte mein
Enkel nach der ersten Tasse Kaffee gleich ich sollte
endlich mein Gala Abo abbestellen, da ich sowieso die
kleinen Buchstaben nicht mehr lesen kann.

Ich will diese Zeitschrift sowieso nicht mehr seit
diesem negativen Bericht in der Apotheken-Rundschau. 

Aber manchmal haben die einfach gute Berichte. 

Was soll ich jetzt tun?

Ich bin auch gerne erreichbar unter (040) 555 555 123 
und erna@badmail.de 
und erna@myprivateEmail.de
und https://facebook.com/erna-koeln-moers
und https://twitter.com/@erna-koeln-moers

Erika, 


# OFFLINE PREFLIGHT ANALYSIS   
## Language                  : German
## Confidence                : 100% [valid]
## SpellFixes                : [none]
## Customer Email            : erika.mustermann@t-online.de
## Customer Email RFC5322    : [valid]
## Customer Email Domain MX  : [valid]
## Customer DB entry         : [valid]
## Anonymized EMailAddresses : erna@badmail.de,erna@myprivateEmail.de
## Anonymized URLs           : https://facebook.com/erna-koeln-moers,https://twitter.com/@erna-koeln-moers
## Anonymized PhoneNumbers   : [none]
## Raw / Filtered Characters : 739 / 561
## Raw / Filtered Words      : 96 / 88
## Raw / Filtered GPT3 Token : 284 / 212
## Raw / Filtered GPT3 Price : 0.00568 US$  / 0.00424 US$ 
## Time needed for section   : 4.04336177s

# ONLINE ML MODULES RESOLVER 
## OpenAI query state        : [valid]
## OpenAI says cancel        : [no]
## OpenAI analysis [debug]   : 

No, this email does not appear to be an attempt to cancel a 
subscription service. The reader instead appears to be asking 
for advice on whether or not to cancel the subscription.

## Time needed for section   : 5.274627552s


cat ../../example-email6.txt | ./aiagent
####################################################
####################################################
#          -= AI MESSAGE ANALYSIS REPORT =-        #
####################################################
####################################################
# INBOUND MESSAGE [debug]    : 

From:    angry.customer@totalfakeemail
To:      contract@service.com
Subject: I'm very angry!

I`m so angry that I`m very bad at spell the words, 
so you will have a hard time to read my email with
your computer.

Please do me a favour and honour my explict wish to 
cancel my subscriptons. NOW! Or I post on faceboook 
about you hypocrite!

Angry Customer 


# OFFLINE PREFLIGHT ANALYSIS   
## Language                  : English
## Confidence                : 100% [valid]
## SpellFixes                : explicit,subscriptions,facebook
## Customer Email            : 
## Customer Email RFC5322    : [failed] [exit]
## Customer Email Domain MX  : [failed] [exit]
## Customer DB entry         : [valid]
## Anonymized EMailAddresses : [none]
## Anonymized URLs           : [none]
## Anonymized PhoneNumbers   : [none]
## Raw / Filtered Characters : 359 / 291
## Raw / Filtered Words      : 58 / 54
## Raw / Filtered GPT3 Token : 108 / 77
## Raw / Filtered GPT3 Price : 0.00216 US$  / 0.00154 US$ 
## No OpenAI query performed, inbound data quality failed.
## Time needed for section   : 4.289477605s


cat ../../example-email7.txt | ./aiagent
####################################################
####################################################
#          -= AI MESSAGE ANALYSIS REPORT =-        #
####################################################
####################################################
# INBOUND MESSAGE [debug]    : 

From:    marie@gmail.com
To:      contract@service.com
Subject: Hola!

Por favor, cancele el servicio de suscripción
lo antes posible.

Atentamente,
Marie


# OFFLINE PREFLIGHT ANALYSIS   
## Language                  : Spanish
## Confidence                : 100% [valid]
## SpellFixes                : [none]
## Customer Email            : marie@gmail.com
## Customer Email RFC5322    : [valid]
## Customer Email Domain MX  : [valid]
## Customer DB entry         : [failed] [exit]
## Anonymized EMailAddresses : [none]
## Anonymized URLs           : [none]
## Anonymized PhoneNumbers   : [none]
## Raw / Filtered Characters : 156 / 101
## Raw / Filtered Words      : 18 / 14
## Raw / Filtered GPT3 Token : 59 / 36
## Raw / Filtered GPT3 Price : 0.00118 US$  / 0.00072 US$ 
## Time needed for section   : 3.963220104s

# ONLINE ML MODULES RESOLVER 
## OpenAI query state        : [valid]
## OpenAI says cancel        : [yes]
## OpenAI analysis [debug]   : 

Yes, this email written attempt is attempting to cancel a subscription service.

## OpenAI Auto Response email: 

Hola Marie,

Entendemos que desea cancelar su servicio de suscripción lo antes posible. 
Estamos tristes de verlo irse, pero aceptamos su solicitud de cancelación. 
¿Hay algo que podamos hacer para que reconsidere su decisión y mantenga su 
suscripción? Por favor, háganos saber si hay algo que podamos hacer para ayudarlo.

Atentamente,
[Tu nombre]

## Time needed for section   : 14.111641041s


cat ../../example-email8.txt | ./aiagent
####################################################
####################################################
#          -= AI MESSAGE ANALYSIS REPORT =-        #
####################################################
####################################################
# INBOUND MESSAGE [debug]    : 

From:    john.doe@invalidemail
To:      subscription@service.com
Subject: I`m done! 

Hello MyService Corp,
I do not want to pay anymore for your
service. Cancel all subscriptions. Now!

Greetings,
John


# OFFLINE PREFLIGHT ANALYSIS   
## Language                  : English
## Confidence                : 100% [valid]
## SpellFixes                : [none]
## Customer Email            : 
## Customer Email RFC5322    : [failed] [exit]
## Customer Email Domain MX  : [failed] [exit]
## Customer DB entry         : [valid]
## Anonymized EMailAddresses : [none]
## Anonymized URLs           : [none]
## Anonymized PhoneNumbers   : [none]
## Raw / Filtered Characters : 204 / 139
## Raw / Filtered Words      : 26 / 22
## Raw / Filtered GPT3 Token : 62 / 37
## Raw / Filtered GPT3 Price : 0.00124 US$  / 0.00074 US$ 
## No OpenAI query performed, inbound data quality failed.
## Time needed for section   : 4.339552864s

```

# TODO

Pre-process messages offline 
- [X] analyzing sender email address validity
- [X] analyzing and match message content and language
- [X] fix spelling errors
- [X] detect and isolatlate locally, to save token payload, protect privacy, EU-GDPR compliance
	- [X] detect and remove unwanted additional customer provided email addresses in message body
	- [X] detect and remove unwanted customer provided urls in message body 
	- [ ] detect and remove unwanted customers phone numbers
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

# ARTWORK

Generated by OpenAI.

# CONTRIBUTION

Yes, please! PRs welcome!








