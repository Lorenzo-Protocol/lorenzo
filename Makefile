#!/usr/bin/make -f

COMMIT := $(shell git log -1 --format='%H')
VERSION ?= $(shell git describe --tags --always)

PACKAGES_SIMTEST=$(shell go list ./... | grep '/simulation')
LEDGER_ENABLED ?= true
SDK_PACK := $(shell go list -m github.com/cosmos/cosmos-sdk | sed  's/ /\@/g')
TM_VERSION := $(shell go list -m github.com/cometbft/cometbft | sed 's:.* ::')
DOCKER := $(shell which docker)
BUILDDIR ?= $(CURDIR)/build

export GO111MODULE = on

# process build tags

build_tags = netgo
ifeq ($(LEDGER_ENABLED),true)
  ifeq ($(OS),Windows_NT)
    GCCEXE = $(shell where gcc.exe 2> NUL)
    ifeq ($(GCCEXE),)
      $(error gcc.exe not installed for ledger support, please install or set LEDGER_ENABLED=false)
    else
      build_tags += ledger
    endif
  else
    UNAME_S = $(shell uname -s)
    ifeq ($(UNAME_S),OpenBSD)
      $(warning OpenBSD detected, disabling ledger support (https://github.com/cosmos/cosmos-sdk/issues/1988))
    else
      GCC = $(shell command -v gcc 2> /dev/null)
      ifeq ($(GCC),)
        $(error gcc not installed for ledger support, please install or set LEDGER_ENABLED=false)
      else
        build_tags += ledger
      endif
    endif
  endif
endif

ifeq (cleveldb,$(findstring cleveldb,$(LORENZO_STAKING_BUILD_OPTIONS)))
  build_tags += gcc cleveldb
endif
build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))

whitespace :=
whitespace := $(whitespace) $(whitespace)
comma := ,
build_tags_comma_sep := $(subst $(whitespace),$(comma),$(build_tags))

# process linker flags

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=lorenzo \
		  -X github.com/cosmos/cosmos-sdk/version.AppName=lorenzod \
		  -X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
		  -X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
		  -X "github.com/cosmos/cosmos-sdk/version.BuildTags=$(build_tags_comma_sep)" \
	      -X github.com/cometbft/cometbft/version.TMCoreSemVer=$(TM_VERSION)

ifeq (cleveldb,$(findstring cleveldb,$(LORENZO_STAKING_BUILD_OPTIONS)))
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=cleveldb
endif
ifeq ($(LINK_STATICALLY),true)
  ldflags += -linkmode=external -extldflags "-Wl,-z,muldefs -static"
endif
ifeq (,$(findstring nostrip,$(LORENZO_STAKING_BUILD_OPTIONS)))
  ldflags += -w -s
endif
ldflags += $(LDFLAGS)
ldflags := $(strip $(ldflags))

BUILD_FLAGS := -tags "$(build_tags)" -ldflags '$(ldflags)'
# check for nostrip option
ifeq (,$(findstring nostrip,$(LORENZO_STAKING_BUILD_OPTIONS)))
  BUILD_FLAGS += -trimpath
endif

all: tools install

# The below include contains the tools and runsim targets.
include contrib/devtools/Makefile

.PHONY: install
install: go.sum
	go install -mod=readonly $(BUILD_FLAGS) ./cmd/lorenzod

.PHONY: build
build: go.sum
	go build $(BUILD_FLAGS) -o build/lorenzod ./cmd/lorenzod

###############################################################################
###                                E2E tests                                ###
###############################################################################

# Executes IBC tests via rollup-e2e-testing
e2e-test-ibc:
	cd e2e && go test -timeout=25m -race -v -run TestIBCTransfer .

# Executes IBC tests via rollup-e2e-testing
e2e-test-ibc-timeout:
	cd e2e && go test -timeout=25m -race -v -run TestIBCTransferTimeout .

# Executes all tests via rollup-e2e-testing
e2e-test-all: e2e-test-ibc e2e-test-ibc-timeout

.PHONY: e2e-test-ibc e2e-test-ibc-timeout e2e-test-all

###############################################################################
###                                Protobuf                                 ###
###############################################################################

protoVer=0.11.6
protoImageName=ghcr.io/cosmos/proto-builder:$(protoVer)
protoImage=$(DOCKER) run --network=host --rm -v $(CURDIR):/workspace --workdir /workspace --user 0 $(protoImageName)

# ------
# NOTE: If you are experiencing problems running these commands, try deleting
#       the docker images and execute the desired command again.
#
proto-all: proto-format proto-lint proto-gen

proto-gen:
	@echo "Generating Protobuf files"
	$(protoImage) sh ./scripts/protocgen.sh

proto-swagger-gen:
	@echo "Downloading Protobuf dependencies"
	@make proto-download-deps
	@echo "Generating Protobuf Swagger"
	$(protoImage) sh ./scripts/protoc-swagger-gen.sh

proto-format:
	@echo "Formatting Protobuf files"
	$(protoImage) find ./ -name *.proto -exec clang-format -i {} \;

proto-lint:
	@echo "Linting Protobuf files"
	@$(protoImage) buf lint --error-format=json

proto-check-breaking:
	@echo "Checking Protobuf files for breaking changes"
	$(protoImage) buf breaking --against $(HTTPS_GIT)#branch=main

###############################################################################
###                                Releasing                                ###
###############################################################################

PACKAGE_NAME:=github.com/Lorenzo-Protocol/lorenzo
GOLANG_CROSS_VERSION  = v1.20
GOPATH ?= '$(HOME)/go'
release-dry-run:
	docker run \
		--rm \
		--privileged \
		-e CGO_ENABLED=1 \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-v `pwd`:/go/src/$(PACKAGE_NAME) \
		-v ${GOPATH}/pkg:/go/pkg \
		-w /go/src/$(PACKAGE_NAME) \
		ghcr.io/goreleaser/goreleaser-cross:${GOLANG_CROSS_VERSION} \
		--clean --skip-validate --skip-publish --snapshot

release:
	@if [ ! -f ".release-env" ]; then \
		echo "\033[91m.release-env is required for release\033[0m";\
		exit 1;\
	fi
	docker run \
		--rm \
		--privileged \
		-e CGO_ENABLED=1 \
		--env-file .release-env \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-v `pwd`:/go/src/$(PACKAGE_NAME) \
		-w /go/src/$(PACKAGE_NAME) \
		ghcr.io/goreleaser/goreleaser-cross:${GOLANG_CROSS_VERSION} \
		release --clean --skip-validate

.PHONY: release-dry-run release


###############################################################################
###                                Linting                                  ###
###############################################################################
golangci_lint_cmd=golangci-lint
golangci_version=v1.53.3

lint:
	@echo "--> Running linter"
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(golangci_version)
	@$(golangci_lint_cmd) run --timeout=10m

lint-fix:
	@echo "--> Running linter"
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(golangci_version)
	@$(golangci_lint_cmd) run --fix --out-format=tab --issues-exit-code=0

format:
	@go install mvdan.cc/gofumpt@latest
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(golangci_version)
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "./client/docs/statik/statik.go" -not -path "./tests/mocks/*" -not -name "*.pb.go" -not -name "*.pb.gw.go" -not -name "*.pulsar.go" -not -path "./crypto/keys/secp256k1/*" | xargs gofumpt -w -l
	$(golangci_lint_cmd) run --fix

.PHONY: lint lint-fix format
