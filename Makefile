# If the "VERSION" environment variable is not set, then use this value instead
VERSION?=1.0.0
TIME=$(shell date +%FT%T%z)
GOVERSION=$(shell go version | awk '{print $$3}' | sed s/go//)

LDFLAGS=-ldflags "\
	-X github.com/softwarespot/public-holidays/internal/version.Version=${VERSION} \
	-X github.com/softwarespot/public-holidays/internal/version.Time=${TIME} \
	-X github.com/softwarespot/public-holidays/internal/version.User=${USER} \
	-X github.com/softwarespot/public-holidays/internal/version.GoVersion=${GOVERSION} \
	-s \
	-w \
"

build:
	@echo building to public-holidays-cli
	@go build $(LDFLAGS) -o public-holidays-cli

test:
	go test -cover -v ./...

.PHONY: build test
