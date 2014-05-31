GO ?= go

GOPATH  := $(CURDIR)/_vendor:$(GOPATH)

all: build test


build:
	$(GO) build

test:
	$(GO) test ./...

testrace:
	$(GO) test -race ./...

clean:
	$(GO) clean
