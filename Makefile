
BIN:=bin
GOBIN:=$(CURDIR)/$(BIN)
GB=$(GOBIN)/gobin

$(shell mkdir -p $(BIN))

export PATH:=$(GOBIN):$(PATH)

.PHONY: deps
deps:
	GO111MODULE=off go get -u github.com/myitcv/gobin
	$(GB) github.com/golangci/golangci-lint/cmd/golangci-lint@v1.23.7

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	 go test -cover -race ./...

.PHONY: clean
clean:
	rm -rf $(GOBIN)