.PHONY: build run test

PROJECT?=jumia-go

default: build

build: test build-local

build-local:
	go build -o ./app ./src/cmd

run: build
	./app

test: regenerate
	go fmt ./...
	go test -vet all ./...

test-race:
	go test -v -race ./...

test-integration:
	go test -v -vet all -tags=integration ./... -coverprofile=integration.out

test-all: test test-race test-integration

cover-ci:
	go tool cover -func=integration.out

cover:
	echo unit tests only
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

generate: get-tools
	go generate ./...

clean-mock:
	find src/internal -iname '*_mock.go' -exec rm {} \;

regenerate: clean-mock generate

get-tools:
	go install github.com/golang/mock/mockgen@v1.6.0
