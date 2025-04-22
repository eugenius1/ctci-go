SHELL := bash
.SHELLFLAGS := -euo pipefail -c

GOLANGCI_LINT_VERSION := v2.1

default: build lint test

build:
	go build -v ./...

test:
	go test ./... -race -coverprofile=coverage.out | grep -v "no test files"

# Preserving caches between consecutive runs.
# See https://golangci-lint.run/welcome/install/#docker
GOLANGCI_LINT_DOCKER_RUN := docker run --rm -t \
	-v $(shell pwd):/app -w /app \
	--user $(shell id -u):$(shell id -g) \
	-v $(shell go env GOCACHE):/.cache/go-build -e GOCACHE=/.cache/go-build \
	-v $(shell go env GOMODCACHE):/.cache/mod -e GOMODCACHE=/.cache/mod \
	-v ~/.cache/golangci-lint:/.cache/golangci-lint -e GOLANGCI_LINT_CACHE=/.cache/golangci-lint \
	golangci/golangci-lint:$(GOLANGCI_LINT_VERSION) golangci-lint

lint:
	$(GOLANGCI_LINT_DOCKER_RUN) run

lint-fix:
	$(GOLANGCI_LINT_DOCKER_RUN) run --fix
