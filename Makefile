.PHONY: build
build:
	go build ./cmd/apiserver/

.PHONY: test
test:
	go test -v -race ./...


.DEFAULT_GOAL := build