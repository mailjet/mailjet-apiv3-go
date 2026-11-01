.DEFAULT_GOAL := all

PACKAGE := github.com/mailjet/mailjet-apiv3-go
GOPATH = $(shell go env GOPATH)

GOLANGCI_LINT_PATH = $(PWD)/bin/golangci-lint
GOLANGCI_LINT_VERSION = v2.14.0

NILAWAY_PATH = $(GOPATH)/bin/nilaway
$(NILAWAY_PATH):
	go install go.uber.org/nilaway/cmd/nilaway@latest

.PHONY: all
all: test

.PHONY: test
test:
	go test . -race -count=1

.PHONY: nilaway
nilaway: $(NILAWAY_PATH)
	$(NILAWAY_PATH) -include-pkgs="$(PACKAGE)" -test=false ./...

# linter:
$(GOLANGCI_LINT_PATH):
	curl -sSfL https://golangci-lint.run/install.sh | sh -s -- -b $(dir $(GOLANGCI_LINT_PATH)) $(GOLANGCI_LINT_VERSION)

.PHONY: lint
lint: $(GOLANGCI_LINT_PATH)
	$(GOLANGCI_LINT_PATH) run
