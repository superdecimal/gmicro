
BIN=bin
GOBIN:=$(CURDIR)/$(BIN)
GB=$(BIN)/gobin

$(shell mkdir -p $(BIN))

export GOBIN=$(CURDIR)/$(BIN)
export PATH:=$(GOBIN):$(PATH)
export GO111MODULE=on

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