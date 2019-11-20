BUILD=$(shell git rev-parse HEAD)
PROJECT=${shell basename "$(PWD)"}

GOBIN=$(shell pwd)/bin
GOVENDOR=${shell pwd}/vendor
GOPATH=$(shell pwd)/vendor:$(shell pwd)
GOFILES=$(wildcard cmd/*.go)

build:
	@echo building $(BUILD) as $(PROJECT)
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build -o ${GOBIN}/${PROJECT} $(GOFILES)

run:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go run $(GOBIN)/$(PROJECT)

test:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go test -v $(GOFILES)

remove:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) rm -f $(GOBIN)/$(PROJECT)

clean:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) rm -rf $(GOBIN) ${GOVENDOR}

.PHONY: test remove clean