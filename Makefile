# Allow users to define GOFLAGS, but override it with flags mandatory for
# compilation.
GOFLAGS ?=
ifndef .GOFLAGS_GUARD
.GOFLAGS_GUARD := 1
override GOFLAGS := -mod=vendor $(GOFLAGS)
endif
export GOFLAGS
ifdef DEBUG
$(info GOFLAGS = $(GOFLAGS))
endif

SOURCES := $(shell \
	find . -name '*.go' | \
	grep -Ev './(vendor)/' | \
	xargs)
ifdef DEBUG
$(info SOURCES = $(SOURCES))
endif

GO ?= go
GOGENERATE ?= $(GO) generate
GOINSTALL ?= $(GO) install
GOBUILD ?= $(GO) build
GOMOD ?= $(GO) mod
GORUN ?= $(GO) run
GOTEST ?= $(GO) test
GOTOOL ?= $(GO) tool

GOIMPORTS ?= $(GORUN) golang.org/x/tools/cmd/goimports
GOLANGCI_LINT := $(GORUN) github.com/golangci/golangci-lint/cmd/golangci-lint --timeout 10m
GOLINT := $(GORUN) golang.org/x/lint/golint

################################################################################
## Go tests 
################################################################################

.PHONY: build
build:
	$(GOBUILD) ./...

.PHONY: test
test:
	$(GOTEST) -failfast -coverprofile=coverage.out ./... $(SILENT_CMD_SUFFIX)


.PHONY: test/race
test/race:
	$(GOTEST) -race ./... $(SILENT_CMD_SUFFIX)

.PHONY: cover
cover: cover/text

.PHONY: cover/html
cover/html:
	$(GOTOOL) cover -html=coverage.out

.PHONY: cover/text
cover/text:
	$(GOTOOL) cover -func=coverage.out

################################################################################
## Linters and formatters
################################################################################

.PHONY: goimports
goimports:
	@$(GOIMPORTS) -w $(SOURCES)

.PHONY: lint
lint:
	$(GOLANGCI_LINT) run

.PHONY: lint-comments
lint-comments:
	$(GO) list ./... | grep -v /vendor/ | xargs -L1 $(GOLINT) -set_exit_status

.PHONY: vendors
vendors:
	$(GOMOD) vendor
	$(GOMOD) tidy

.PHONY: git/diff
git/diff:
	@if ! $(GITDIFF) --quiet; then \
		printf 'Found changes on local workspace. Please run this target and commit the changes\n' ; \
		exit 1; \
	fi


