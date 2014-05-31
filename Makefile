GO ?= go
export GOPATH:=$(shell pwd)
export BUILDTAGS=debug


default: all

all: client daemon fmt

deps:
	go get  -d github.com/zdannar/flogger

fmt:
	go fmt ./...

client: deps
	go install -tags '$(BUILDTAGS)' client

daemon: deps
	go install -tags '$(BUILDTAGS)' daemon

test:
	$(GO) test ./...

testrace:
	$(GO) test -race ./...

clean:
	$(GO) clean
