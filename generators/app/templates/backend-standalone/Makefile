# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
SWAGCMD=~/go/bin/swag

all: test

doc:
	$(SWAGCMD) init
run: doc
	$(GOCMD) run main.go
test:
	OPS_CONFIG=../config.test.yaml $(GOTEST) -v ./...

cover:
	OPS_CONFIG=../config.test.yaml $(GOTEST) ./... -coverprofile=tmp/cover.out

html:
	$(GOCMD) tool cover -html=tmp/cover.out -o tmp/cover.html
