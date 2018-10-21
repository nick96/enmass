.PHONY: build test generate generate-peopleapi generate-messaging test-cli test-messaging test-peopleapi

build:
	go build

test: generate test-cli test-messaging test-peopleapi

generate: generate-messaging generate-peopleapi

generate-peopleapi:
	@cd peopleapi && go generate

generate-messaging:
	@cd messaging && go generate

test-cli: generate
	@go test

test-messaging: generate-messaging
	@cd messaging && go test

test-peopleapi: generate-peopleapi
	@cd peopleapi && go test
