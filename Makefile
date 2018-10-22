PROJ="github.com/nick96/enmass"
COVER="cover.out"

.PHONY: build test generate coverage

build:
	go build

test: generate
	@go test ${PROJ} ${PROJ}/messaging ${PROJ}/peopleapi

generate: 
	@go generate ${PROJ} ${PROJ}/messaging ${PROJ}/peopleapi

coverage: 
	@go test ${PROJ} ${PROJ}/messaging ${PROJ}/peopleapi -coverprofile=${COVER}
	@go tool cover -html=${COVER}

