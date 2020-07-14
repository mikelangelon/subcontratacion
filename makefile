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

cronjobdh:
	cd cronjob; \
    docker build --tag mikelangelon/cronjob .
	docker push mikelangelon/cronjob
docker:
	docker build --tag mikelangelon/test .
	docker container run -p 8080:8080 -d --name subcon mikelangelon/test
dockerrm:
	docker container stop subcon
	docker container rm subcon

kapply:
	kubectl apply -f .

kdelete:
	kubectl delete -f .

generateapi:
	rm -rf /api/model
	rm -rf /api/rest
	$(SWAGGER_GEN)

dev.start:
	docker-compose up --detach --no-recreate

dev.rm:
	docker-compose down




