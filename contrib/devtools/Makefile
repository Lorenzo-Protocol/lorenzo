###
# Find OS and Go environment
# GO contains the Go binary
# FS contains the OS file separator
###
OS := $(shell uname)

ifeq ($(OS),Windows_NT)
  GO := $(shell where go.exe 2> NUL)
  FS := "\\"
else
  GO := $(shell command -v go 2> /dev/null)
  FS := "/"
endif

ifeq ($(GO),)
  $(error could not find go. Is it in PATH? $(GO))
endif

###############################################################################
###                                Functions                                ###
###############################################################################

go_get = $(if $(findstring Windows_NT,$(OS)),\
IF NOT EXIST $(GITHUBDIR)$(FS)$(1)$(FS) ( mkdir $(GITHUBDIR)$(FS)$(1) ) else (cd .) &\
IF NOT EXIST $(GITHUBDIR)$(FS)$(1)$(FS)$(2)$(FS) ( cd $(GITHUBDIR)$(FS)$(1) && git clone https://github.com/$(1)/$(2) ) else (cd .) &\
,\
mkdir -p $(GITHUBDIR)$(FS)$(1) &&\
(test ! -d $(GITHUBDIR)$(FS)$(1)$(FS)$(2) && cd $(GITHUBDIR)$(FS)$(1) && git clone https://github.com/$(1)/$(2)) || true &&\
)\
cd $(GITHUBDIR)$(FS)$(1)$(FS)$(2) && git fetch origin && git checkout -q $(3)

mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
mkfile_dir := $(shell cd $(shell dirname $(mkfile_path)); pwd)


###############################################################################
###                                 Tools                                   ###
###############################################################################

PREFIX ?= /usr/local
BIN ?= $(PREFIX)/bin
UNAME_S ?= $(shell uname -s)
UNAME_M ?= $(shell uname -m)

GOPATH ?= $(shell $(GO) env GOPATH)
GITHUBDIR := $(GOPATH)$(FS)src$(FS)github.com

TOOLS_DESTDIR  ?= $(GOPATH)/bin
STATIK         = $(TOOLS_DESTDIR)/statik
RUNSIM         = $(TOOLS_DESTDIR)/runsim

tools: pre-commit

# Install pre-commit
pre-commit:
	@echo "Installing pre-commit..."
	@if [ "$(OS)" = "Darwin" ]; then \
		echo "Detected macOS. Installing pre-commit using brew..."; \
		brew install pre-commit; \
	else \
		if command -v pip > /dev/null 2>&1; then \
			echo "Detected pip. Installing pre-commit..."; \
			pip install pre-commit; \
		elif command -v pip3 > /dev/null 2>&1; then \
			echo "Detected pip3. Installing pre-commit..."; \
			pip3 install pre-commit; \
		else \
			echo "Neither pip nor pip3 found. Please install pip or pip3 to proceed."; \
			exit 1; \
		fi \
	fi


.PHONY: tools pre-commit