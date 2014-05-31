GO ?= go

GOPATH  := $(CURDIR)/_vendor:$(GOPATH)

all: build test


build:
	$(GO) build client
	$(GO) build daemon

test:
	$(GO) test ./...

testrace:
	$(GO) test -race ./...

clean:
	$(GO) clean
