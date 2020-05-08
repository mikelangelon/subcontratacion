#Go parameters

GOCMD=go
GOBUILD=$(GOCMD) build
BINARY_NAME = testapp

build:
	$(GOBUILD) -o $(BINARY_NAME) -v
