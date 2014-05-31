GO ?= go
export GOPATH:=$(shell pwd)
export BUILDTAGS=debug


default: all

all: client daemon fmt

fmt:
	go fmt ./...

client:
	go install -tags '$(BUILDTAGS)' client

daemon:
	go install -tags '$(BUILDTAGS)' daemon

test:
	$(GO) test ./...

testrace:
	$(GO) test -race ./...

clean:
	$(GO) clean
