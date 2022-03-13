GOBIN ?= $(shell go env GOPATH)/bin

.PHONY: all
all: test

.PHONY: test
test:
	go test -v ./... # do not add -race for test with GOARCH=386

.PHONY: lint
lint: $(GOBIN)/staticcheck
	go vet ./...
	staticcheck -checks all,-ST1000 ./...

$(GOBIN)/staticcheck:
	cd && go get honnef.co/go/tools/cmd/staticcheck

.PHONY: clean
clean:
	go clean
