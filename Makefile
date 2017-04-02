GO  := GO15VENDOREXPERIMENT=1 go
pkgs = $(shell $(GO) list ./... | grep -v /vendor/)


all: check vet

check:
	@echo ">> checking formatting"
	@! gofmt -d $(shell find . -path ./vendor -prune -o -name '*.go' -print) | grep '^'

vet:
	@echo ">> vetting code"
	@$(GO) vet $(pkgs)

.PHONY: all check vet
