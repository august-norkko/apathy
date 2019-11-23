BUILD=$(shell git rev-parse HEAD)
PROJECT=${shell basename "$(PWD)"}

GOBIN=$(shell pwd)/bin
GOVENDOR=${shell pwd}/vendor
GOFILES=$(wildcard *.go)

.PHONY: build
build:
	@echo building $(BUILD) as $(PROJECT)
	go build -o apathy $(GOFILES)

.PHONY: run
run:
	@echo running binary
	go run $(GOBIN)/$(PROJECT)

.PHONY: test
test:
	@echo running tests
	go test -v $(GOFILES)

.PHONY: clean
clean:
	@echo removing build
	rm -f $(GOBIN) $(GOVENDOR) $(GOBIN)/$(PROJECT) 
