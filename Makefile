BIN=bin
GOBIN:=$(CURDIR)/$(BIN)
DOCKER_ACCOUNT = superdecimal
SERVICES_DOCKERFILES  = $(wildcard services/*/Dockerfile)
CHARTS = $(wildcard deploy/*/Chart.yaml)
BRANCH_VERSION=$(shell git rev-parse --short HEAD)

# Tools
GB=$(BIN)/gobin
PC=protoc
LINT=golangci-lint

$(shell mkdir -p $(BIN))

export GOBIN=$(CURDIR)/$(BIN)
export PATH:=$(PATH):$(GOBIN)
export GO111MODULE=on

# We do not support windows. Sorry :)
UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Linux)
	OS = linux
endif
ifeq ($(UNAME_S),Darwin)
	OS = darwin
endif

.PHONY: base-tools
base-tools:
	GO111MODULE=off go get -u github.com/myitcv/gobin

.PHONY: tools
ci-tools: base-tools
	curl https://raw.githubusercontent.com/helm/helm/master/scripts/get-helm-3 | bash
	$(GB) github.com/golangci/golangci-lint/cmd/golangci-lint@v1.23.7

.PHONY: lint
lint: ci-tools
	$(LINT) run

.PHONY: test
test:
	 go test -cover -race ./...

.PHONY: clean
clean:
	rm -rf $(GOBIN)

.PHONY: generate-proto
generate-proto:
	@echo GENERATING PROTO...
	@protoc  --go_out=plugins=grpc:pkg -I=$(PWD) proto/*.proto
	@echo DONE

.PHONY: generate-mocks
generate-mocks:
	@echo GENERATING MOCKS...
	@mkdir -p pkg/proto/mock/
	@go generate ./...
	@echo DONE

build-all: $(SERVICES_DOCKERFILES) 

services/%/Dockerfile:
	docker build -f $@ -t $(DOCKER_ACCOUNT)/$*:$(BRANCH_VERSION) .
	docker tag $(DOCKER_ACCOUNT)/$*:$(BRANCH_VERSION) $(DOCKER_ACCOUNT)/$*:latest

deploy/%/Chart.yaml:
	helm lint deploy/$*

lint-all-charts: $(CHARTS)

.PHONY: dev-env
dev-env: base-tools
ifeq ($(OS),darwin)
	brew install protobuf
	brew install helm
endif
ifeq ($(OS),linux)
	apt update
	apt install protobuf-compiler
	curl https://raw.githubusercontent.com/helm/helm/master/scripts/get-helm-3 | bash

endif
	$(GB) github.com/golang/protobuf/protoc-gen-go@v1.3.4
	$(GB) github.com/golang/mock/mockgen@v1.4.1
