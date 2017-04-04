GO  := GO15VENDOREXPERIMENT=1 go

all: check

check:
	@echo ">> checking formatting"
	@! gofmt -d $(shell find . -path ./vendor -prune -o -name '*.go' -print) | grep '^'

.PHONY: all check
