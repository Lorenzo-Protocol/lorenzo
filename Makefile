#!/usr/bin/make -f

COMMIT := $(shell git log -1 --format='%H')
VERSION ?= $(shell git describe --tags --always)
LEDGER_ENABLED ?= true
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

ifeq (cleveldb,$(findstring cleveldb,$(LORENZO_BUILD_OPTIONS)))
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

ifeq (cleveldb,$(findstring cleveldb,$(LORENZO_BUILD_OPTIONS)))
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=cleveldb
endif
ifeq ($(LINK_STATICALLY),true)
  ldflags += -linkmode=external -extldflags "-Wl,-z,muldefs -static"
endif
ifeq (,$(findstring nostrip,$(LORENZO_BUILD_OPTIONS)))
  ldflags += -w -s
endif
ldflags += $(LDFLAGS)
ldflags := $(strip $(ldflags))

BUILD_FLAGS := -tags "$(build_tags)" -ldflags '$(ldflags)'

# check for nostrip option
ifeq (,$(findstring nostrip,$(LORENZO_BUILD_OPTIONS)))
  BUILD_FLAGS += -trimpath
endif

# check for debug option
ifeq (debug,$(findstring debug,$(LORENZO_BUILD_OPTIONS)))
  BUILD_FLAGS += -gcflags "all=-N -l"
endif

all: tools install

# The below include contains the tools and runsim targets.
include contrib/devtools/Makefile

###############################################################################
###                                 Build                                   ###
###############################################################################

.PHONY: install
install: go.sum
	go install -mod=readonly $(BUILD_FLAGS) ./cmd/lorenzod

.PHONY: build
build: go.sum
	go build $(BUILD_FLAGS) -o build/lorenzod ./cmd/lorenzod

.PHONY: clean
clean:
	rm -rf \
	$(BUILDDIR)/ \
	.testnets

###############################################################################
###                           Tests & Simulation                            ###
###############################################################################

test: test-unit
test-all: test-unit test-race
# For unit tests we don't want to execute the upgrade tests in tests/e2e but
# we want to include all unit tests in the subfolders (tests/e2e/*)
PACKAGES_UNIT=$(shell go list ./... | grep -v '/tests/e2e$$')
TEST_PACKAGES=./...
TEST_TARGETS := test-unit test-unit-cover test-race

# Test runs-specific rules. To add a new test target, just add
# a new rule, customise ARGS or TEST_PACKAGES ad libitum, and
# append the new rule to the TEST_TARGETS list.
test-unit: ARGS=-timeout=15m -race
test-unit: TEST_PACKAGES=$(PACKAGES_UNIT)

test-race: ARGS=-race
test-race: TEST_PACKAGES=$(PACKAGES_NOSIMULATION)
$(TEST_TARGETS): run-tests

test-unit-cover: ARGS=-timeout=15m -race -coverprofile=coverage.txt -covermode=atomic
test-unit-cover: TEST_PACKAGES=$(PACKAGES_UNIT)

test-e2e:
	@if [ -z "$(TARGET_VERSION)" ]; then \
		echo "Building docker image from local codebase"; \
		make build-docker; \
	fi
	@mkdir -p ./build
	@rm -rf build/.tabid
	@INITIAL_VERSION=$(INITIAL_VERSION) TARGET_VERSION=$(TARGET_VERSION) \
	E2E_SKIP_CLEANUP=$(E2E_SKIP_CLEANUP) MOUNT_PATH=$(MOUNT_PATH) CHAIN_ID=$(CHAIN_ID) \
	go test -v ./tests/e2e -run ^TestIntegrationTestSuite$

run-tests:
ifneq (,$(shell which tparse 2>/dev/null))
	go test -mod=readonly -json $(ARGS) $(EXTRA_ARGS) $(TEST_PACKAGES) | tparse
else
	go test -mod=readonly $(ARGS)  $(EXTRA_ARGS) $(TEST_PACKAGES)
endif

test-import:
	@go test ./tests/importer -v --vet=off --run=TestImportBlocks --datadir tmp \
	--blockchain blockchain
	rm -rf tests/importer/tmp

test-rpc:
	./scripts/integration-test-all.sh -t "rpc" -q 1 -z 1 -s 2 -m "rpc" -r "true"

test-rpc-pending:
	./scripts/integration-test-all.sh -t "pending" -q 1 -z 1 -s 2 -m "pending" -r "true"

.PHONY: run-tests test test-all test-import test-rpc $(TEST_TARGETS)

benchmark:
	@go test -mod=readonly -bench=. $(PACKAGES_NOSIMULATION)
.PHONY: benchmark

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

GOLANGCI_LINT_CMD=golangci-lint
GOLANGCI_VERSION=v1.53.3

lint:
	@echo "--> Running linter"
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLANGCI_VERSION)
	@$(GOLANGCI_LINT_CMD) run --timeout=10m

lint-fix:
	@echo "--> Running linter"
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLANGCI_VERSION)
	@$(GOLANGCI_LINT_CMD) run --fix --out-format=tab --issues-exit-code=0

format:
	@go install mvdan.cc/gofumpt@latest
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLANGCI_VERSION)
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "./client/docs/statik/statik.go" -not -path "./tests/mocks/*" -not -name "*.pb.go" -not -name "*.pb.gw.go" -not -name "*.pulsar.go" -not -path "./crypto/keys/secp256k1/*" | xargs gofumpt -w -l
	$(GOLANGCI_LINT_CMD) run --fix

.PHONY: lint lint-fix format

###############################################################################
###                                Localnet                                 ###
###############################################################################

localnet-build-env:
	$(MAKE) -C contrib/images lorenzod-env

localnet-build-dlv:
	$(MAKE) -C contrib/images lorenzod-dlv

localnet-init:
	$(DOCKER) run --rm -v $(CURDIR)/.testnets:/data lorenzo/lorenzod \
			  testnet init-files --v 4 -o /data --starting-ip-address 192.168.10.2 --keyring-backend=test

localnet-start:
	docker-compose up -d

localnet-stop:
	docker-compose down

###############################################################################
###                        Compile Solidity Contracts                       ###
###############################################################################

# Currently only support for compiling ERC20 contracts.
CONTRACTS_PRJ := erc20
CONTRACTS_DIR := contracts/$(CONTRACTS_PRJ)
EVM_VERSION := paris
COMPILED_DIR := $(CONTRACTS_DIR)/compiled_contracts
TMP_DIR := $(CONTRACTS_DIR)/tmp
TMP_COMPILED := $(TMP_DIR)/compiled.json
TMP_JSON := $(TMP_DIR)/tmp.json

contracts-compile: contracts-check-project-name contracts-install-dependencies contracts-create-json contracts-clean-up

contracts-check-project-name:
	@if [ -z "$(CONTRACTS_PRJ)" ]; then \
		echo "CONTRACTS_PRJ is not set. Please set it to the project name of the contracts you want to compile."; \
		exit 1; \
	fi

# Install open-zeppelin solidity contracts
contracts-install-dependencies:
	@echo "Importing open-zeppelin contracts..."
	@cd $(CONTRACTS_DIR) && npm install
	@mv $(CONTRACTS_DIR)/node_modules/@openzeppelin $(CONTRACTS_DIR)
	@rm -rf $(CONTRACTS_DIR)/node_modules $(CONTRACTS_DIR)/package-lock.json

# Compile, filter out and format contracts into the following format.
contracts-create-json:
	@for c in $(shell ls $(CONTRACTS_DIR) | grep '\.sol' | sed 's/.sol//g'); do \
  		mkdir -p $(TMP_DIR) ;\
  		mkdir -p $(COMPILED_DIR) ;\
		echo "\nCompiling solidity contract $${c}..." ;\
		solc --evm-version $(EVM_VERSION) --combined-json abi,bin $(CONTRACTS_DIR)/$${c}.sol > $(TMP_COMPILED);\
		echo "Formatting JSON..." ;\
		get_contract=$$(jq '.contracts["$(CONTRACTS_DIR)/'$$c'.sol:'$$c'"]' $(TMP_COMPILED)) ;\
		add_contract_name=$$(echo $$get_contract | jq '. + { "contractName": "'$$c'" }') ;\
		echo $$add_contract_name | jq > $(TMP_JSON) ;\
		abi_string=$$(echo $$add_contract_name | jq -cr '.abi') ;\
		echo $$add_contract_name | jq --arg newval "$$abi_string" '.abi = $$newval' > $(TMP_JSON) ;\
        mv $(TMP_JSON) $(COMPILED_DIR)/$${c}.json ;\
        rm -rf $(TMP_DIR) ;\
	done

# Clean tmp files
contracts-clean-up:
	@echo "Cleaning up..."
	@rm -rf $(CONTRACTS_DIR)/@openzeppelin
	@echo "Finished cleaning up."