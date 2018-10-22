CMD="github.com/nick96/enmass/cmd/enmass"
PKG="github.com/nick96/enmass/pkg"
COVER="cover.out"

.PHONY: build test generate coverage \
	test-cmd test-messaging test-peopleapi \
	generate-cmd generate-messaging generate-peopleapi

build:
	go build ${CMD}

test: test-cmd test-messaging test-peopleapi

test-cmd: generate-cmd
	@go test ${CMD} -cover

test-messaging: generate-messaging
	@go test ${PKG}/messaging -cover

test-peopleapi: generate-peopleapi
	@go test ${PKG}/peopleapi -cover

generate: generate-cmd generate-messaging generate-peopleapi
	@go generate ${CMD} ${PKG}/messaging ${PKG}/peopleapi

generate-cmd:
	@go generate ${CMD}

generate-messaging:
	@go generate ${PKG}/messaging

generate-peopleapi:
	@go generate ${PKG}/peopleapi

coverage: 
	@go test ${CMD} ${PKG}/messaging ${PKG}/peopleapi -coverprofile=${COVER}
	@go tool cover -html=${COVER}
