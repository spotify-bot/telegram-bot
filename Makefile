#!make

SHELL=/bin/sh

CONFIF_ENV_FILE ?= app.env

lint:
	golangci-lint run --disable errcheck

deps:
	go mod download

build: deps
	env GOOS=linux GOARCH=amd64 go build -o build cmd/main.go

docker.build:
	docker compose build --pull


docker.run:
	docker compose \
	--env-file ${CONFIF_ENV_FILE} \
	up -d
