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
	<%= displayName %>=../config.test.yaml $(GOTEST) -v ./...

cover:
	<%= displayName %>=../config.test.yaml $(GOTEST) ./... -coverprofile=tmp/cover.out

html:
	$(GOCMD) tool cover -html=tmp/cover.out -o tmp/cover.html
