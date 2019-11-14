GOBIN=$(shell pwd)/bin
GOPATH=$(shell pwd)/vendor:$(shell pwd)
GOFILES=$(wildcard cmd/*.go)


build:
	@echo "Building"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build -o bin/apathy $(GOFILES)

run:
	@echo "Running $(GOFILES)"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go run $(GOFILES)

stop:
	@echo "kill pid"
	# ...

clean:
	@echo "Cleaning bin"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) rm -rf $(GOBIN)

.PHONY: build run start clean