GOCMD=go
GOBUILD=$(GOCMD) build -o ./
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test -v. 
GOGET=$(GOCMD) get
GOVET=$(GOCMD) vet

.PHONY: deps clean build

all: deps test build

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
	rm ./password_validator
	
build:
	$(GOBUILD)password_validator ./password_validator.go
