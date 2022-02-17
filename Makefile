include Makefile.env

GOARCH = amd64
OS = linux

.DEFAULT_GOAL := all

all: source-build

.PHONY: source-build
source-build:
	CGO_ENABLED=0 GOOS="$(OS)" GOARCH="$(GOARCH)" go build -o write-module main.go

.PHONY: docker-build
docker-build: source-build
	docker build . -t ${IMG} -f Dockerfile
	rm write-module


.PHONY: clean
clean:
	rm -f ./write-module

.PHONY: test
test:
	go test -v ./...

include hack/make-rules/verify.mk
include hack/make-rules/tools.mk
include hack/make-rules/helm.mk
include hack/make-rules/docker.mk
