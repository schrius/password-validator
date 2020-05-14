GOCMD=go
GOBUILD=$(GOCMD) build -o ./bin
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOVET=$(GOCMD) vet

.PHONY: deps clean build

all: deps test clean build

# test all packages
test:
	$(GOTEST) ./...

# get all packages
deps:
	$(GOGET) -u ./...

# static analysis
static:
	$(GOVET) ./...

clean: 
	rm -rf ./bin
	
build:
	$(GOBUILD) ./password_validator.go
