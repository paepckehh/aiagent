#!/bin/sh
export HTTPS_PROXY="192.168.10.80:9090"
export SSL_CERT_FILE="/etc/ssl/rootCA.pem"
export OPENAI_API_TOKEN="$( cat /usr/store/.keys/openai/token )"
rm ./aicancel > /dev/null 2>&1
go build -v
cat ../../example-email1.txt | ./aicancel
cat ../../example-email2.txt | ./aicancel
cat ../../example-email3.txt | ./aicancel
cat ../../example-email4.txt | ./aicancel
cat ../../example-email5.txt | ./aicancel
