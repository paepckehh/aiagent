#!/bin/sh
doit() {
	echo "cat ../../$FILE | ./aiagent"
	cat ../../$FILE | ./aiagent
}
export HTTPS_PROXY="192.168.10.80:9090"
export SSL_CERT_FILE="/etc/ssl/rootCA.pem"
export OPENAI_API_TOKEN="$(cat /usr/store/.keys/openai/token)"
rm ./aiagent > /dev/null 2>&1
# go build -v
go build -v -mod=readonly
FILE="example-email1.txt" && doit
FILE="example-email2.txt" && doit
FILE="example-email3.txt" && doit
# FILE="example-email4.txt" && doit
# FILE="example-email5.txt" && doit
FILE="example-email6.txt" && doit
FILE="example-email7.txt" && doit
FILE="example-email8.txt" && doit
rm ./aiagent > /dev/null 2>&1
