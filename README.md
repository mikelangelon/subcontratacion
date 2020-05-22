# Project Title

This is a golang project just to play with kubernetes. Personal development & stuff.

## Getting Started

Run go mod init && go mod vendor
Run `make generateapi` to generate the swagger api files

### Prerequisites

Go 1.14: https://golang.org/
kubectl:  https://kubernetes.io/docs/tasks/tools/install-kubectl/

Something to use kubernetes locally/remote. As for example okteto.com account.
Account in okteto.com or some alternative to use kubernetes

Using okteto.com:
- Create account in okteto.com 
- Link kubernetes credentials to it: https://okteto.com/docs/cloud/credentials
    - get okteto CLI: curl https://get.okteto.com -sSfL | sh
    - okteto login

### Getting it working

Either:
Run `make run`

or build and run the docker image
`docker build --tag mikelangelon/test .`
`docker run -p 8080:8080 mikelangelon/test`

## Deployment

First, add the docker image in docker hub to allow kubernetes to use it:

`docker login --username=mikelangelon --password=$MYPASSWORD`
`docker build --tag mikelangelon/test .`
`docker push mikelangelon/test`

Now we can apply k8s

`docker apply -f .`