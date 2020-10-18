#
# aws-whats-new-bot
# Copyright (c) 2020 - Puru Tuladhar (ptuladhar3@gmail.com)
# See LICENSE file.
# 
go-install:
	cd src && go mod download

go-build:
	cd src && \
	GOOS=linux GOARCH=amd64 go build -o ../bin/bot .

go-run:
	cd src && go run .