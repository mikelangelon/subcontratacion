#Go parameters

GO_CMD=go
GO_BUILD=$(GO_CMD) build -mod=vendor
BINARY_NAME = coolapp
MAIN_PATH = cmd/supercoolservice

SWAGGER_GEN = swagger generate server --spec ./api/spec.yml --name coolapp --target ./api/ --model-package model --api-package operation --server-package rest --exclude-main

build:
	$(GO_BUILD) -o $(BINARY_NAME) $(MAIN_PATH)/main.go

run:
	go run cmd/supercoolservice/main.go

redisrun:
	docker run -d -p 6379:6379 redis

dockerhub:
	docker login --username=mikelangelon
	docker build --tag mikelangelon/test .
	docker push mikelangelon/test

kapply:
	kubectl apply -f .

kdelete:
	kubectl delete -f .

generateapi:
	$(SWAGGER_GEN)

