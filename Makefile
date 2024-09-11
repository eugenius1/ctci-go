SHELL := bash
.SHELLFLAGS := -euo pipefail -c

test:
	go test ./... -race | grep -v "no test files"

lint:
	golangci-lint run