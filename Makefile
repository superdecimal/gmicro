
BIN=bin
GOBIN:=$(CURDIR)/$(BIN)
GB=$(BIN)/gobin
MINI=$(BIN)/minikube

$(shell mkdir -p $(BIN))

export GOBIN=$(CURDIR)/$(BIN)
export PATH:=$(GOBIN):$(PATH)
export GO111MODULE=on

# We do not support windows. Sorry :)
UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Linux)
	OS = linux
endif
ifeq ($(UNAME_S),Darwin)
	OS = darwin
endif

.PHONY: deps
deps:
	GO111MODULE=off go get -u github.com/myitcv/gobin
	$(GB) github.com/golangci/golangci-lint/cmd/golangci-lint@v1.23.7

.PHONY: lint
lint: deps
	golangci-lint run

.PHONY: test
test:
	 go test -cover -race ./...

.PHONY: clean
clean:
	rm -rf $(GOBIN)

minikube:
	curl -Lo $(MINI) https://storage.googleapis.com/minikube/releases/latest/minikube-$(OS)-amd64 \
	&& chmod +x $(MINI)
	$(MINI) start