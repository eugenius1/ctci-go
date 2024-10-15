SHELL := bash
.SHELLFLAGS := -euo pipefail -c

all:
	go build -v ./...

test:
	go test ./... -race -coverprofile=coverage.out | grep -v "no test files"

lint:
	golangci-lint run

lint-fix:
	golangci-lint run --fix
