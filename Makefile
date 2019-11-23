# Makefile

BUILD=$(shell git rev-parse HEAD)
PROJECT=$(shell basename "${PWD}")
GOVENDOR=$(shell pwd)/vendor
GOFILES=$(wildcard *.go)
TESTFILES=$(wildcard **/*_test.go)

.PHONY: build
build:
	@echo building $(BUILD) as $(PROJECT)
	GOOS=linux go build -o $(PROJECT) $(GOFILES)

.PHONY: run
run:
	@echo running binary
	go run $(PROJECT)

.PHONY: test
test:
	@echo running tests
	go test -v $(TESTFILES)

.PHONY: clean
clean:
	@echo removing build
	rm -rf -f $(GOVENDOR) $(PROJECT)
