#!/bin/sh
export HTTPS_PROXY="192.168.10.80:9090"
export SSL_CERT_FILE="/etc/ssl/rootCA.pem"
export OPENAI_API_TOKEN="$(cat /usr/store/.keys/openai/token)"
rm ./aiagent > /dev/null 2>&1
go build -v
#go build -v -mod=readonly
cat ../../example-email1.txt | ./aiagent
cat ../../example-email2.txt | ./aiagent
cat ../../example-email3.txt | ./aiagent
cat ../../example-email4.txt | ./aiagent
cat ../../example-email5.txt | ./aiagent
cat ../../example-email6.txt | ./aiagent
rm ./aiagent > /dev/null 2>&1
