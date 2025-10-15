.ONESHELL:
.DEFAULT_GOAL := help

# Set the version of the tools
GOLANG_CI_LINT_VERSION=1.62.0 #https://github.com/golangci/golangci-lint
REVIVE_VERSION=1.5.1 #https://github.com/mgechev/revive
MOCKERY_VERSION=2.47.0 #https://github.com/vektra/mockery

# Set the environment variables
GOBIN=$(shell pwd)/bin
GOPROXY=https://goproxy.cn,https://goproxy.io,direct
export GOBIN
export GOPROXY

MODULE_PACKAGE=github.com/versegeek/verse-go
VERSION_PACKAGE=$(MODULE_PACKAGE)/pkg/version
GIT_COMMIT=$(shell git rev-parse HEAD)
ifeq ($(origin GIT_VERSION), undefined)
GIT_VERSION = $(shell git describe --tags --always --match='v*')
endif
GO_LDFLAGS += \
  -X $(VERSION_PACKAGE).GitVersion=$(GIT_VERSION) \
  -X $(VERSION_PACKAGE).GitCommit=$(GIT_COMMIT) \
  -X $(VERSION_PACKAGE).BuildDate=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')


.PHONY: help
help:
	

.PHONY: all
all: tidy build

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: tools
tools:
	cd tools && go mod tidy

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o main ./cmd/main.go

.PHONY: run
run: tidy
	go run cmd/main.go

.PHONY: lint
# lint: goimports golangci-lint
lint: golangci-lint


# .PHONY: goimports
# goimports: bin/goimports
# 	bin/goimports -w .

.PHONY: golangci-lint
golangci-lint: bin/golangci-lint
	$(GOBIN)/golangci-lint run ./...

.PHONY: generate
generate:
	go generate ./...

.PHONY: test
test:
	go test -cover ./...


.PHONY: docker_build
docker_build:
	docker build --force-rm -t main -f Dockerfile .

.PHONY: docker_run
docker_run:
	

.PHONY: bin/goimports
bin/goimports:
	go install golang.org/x/tools/cmd/goimports@latest

.PHONY: bin/golangci-lint
bin/golangci-lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v$(GOLANG_CI_LINT_VERSION)


  #
  #//go:generate go run entgo.io/ent/cmd/ent generate --target ./internal/ent ./ent/schema