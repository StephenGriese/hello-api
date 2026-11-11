NAME := hello-api
BINARY := api

# The sentinel lives in the *common* git dir's hooks directory, resolved
# through git: in a linked worktree `.git` is a file, so a literal
# `.git/hooks` path would not exist there. The common dir (and the
# core.hooksPath config `make init` writes) is shared by every worktree, so
# one init initializes them all. Outside a repository we fall back to the
# plain path, where the guard below fires with its own message instead.
GIT_COMMON_DIR := $(shell git rev-parse --git-common-dir 2>/dev/null)
HOOKS_SENTINEL := $(if $(GIT_COMMON_DIR),$(GIT_COMMON_DIR)/hooks,.git/hooks)/.githooks-installed

.PHONY: all init link checks staticcheck lint test build coverage report check-format modules clean

# Everything runs serially, even under `make -j`: golangci-lint --fix can
# rewrite files while the other checks read them.
.NOTPARALLEL:

# Plain `make` builds: checks, then a clean compile.
all: build

# Guard target — aborts with a helpful message if init was never run.
$(HOOKS_SENTINEL):
	@echo ""
	@echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
	@echo "⚠️  Run 'make init' first to configure git hooks."
	@echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
	@echo ""
	@exit 1

init:
	git config core.hooksPath .githooks
	@mkdir -p $(if $(GIT_COMMON_DIR),$(GIT_COMMON_DIR)/hooks,.git/hooks)
	@touch $(HOOKS_SENTINEL)
	@echo "✓ Git hooks configured. You're ready to develop."

# Create or refresh the ~/bin symlink to the built binary (the convention
# here: ~/bin holds one symlink per project, pointing at that project's
# built binary).
link: build
	@mkdir -p $(HOME)/bin
	@ln -sf $(CURDIR)/$(BINARY) $(HOME)/bin/$(NAME)
	@echo "✓ Linked ~/bin/$(NAME) -> $(CURDIR)/$(BINARY)"

# checks gate the build; the recipe then removes the old binary and
# compiles from scratch, so every build is clean. The removal and the
# compile are recipe lines rather than prerequisites so their order can
# never be scrambled — recipe lines always run in order, unlike
# prerequisites under `make -j`.
build: checks
	rm -f $(BINARY)
	go build -o $(BINARY) cmd/main.go

checks: $(HOOKS_SENTINEL) staticcheck lint test

staticcheck:
	staticcheck ./...

lint:
	golangci-lint -v run --fix ./...

test:
	go test ./... -coverprofile=coverage.out

coverage:
	go tool cover -func coverage.out | grep "total:" | \
	awk '{print ((int($$3) > 80) != 1) }'

report:
	go tool cover -html=coverage.out -o cover.html

check-format:
	test -z $$(go fmt ./...)

modules:
	go mod tidy

clean:
	rm -f $(BINARY) coverage.out cover.html
