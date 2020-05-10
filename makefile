#Go parameters

GOCMD=go
GOBUILD=$(GOCMD) build
BINARY_NAME = coolapp

SWAGGER_GEN = swagger generate server --spec ./api/spec.yml --name coolapp --target ./api/ --model-package model --api-package operation --server-package rest --exclude-main

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

run:
	go run main.go

generateapi:
	$(SWAGGER_GEN)

