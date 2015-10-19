TAG?=latest
APP?=interlock
DOCKERHUBUSER?=carina/$(APP)
COMMIT=`git rev-parse --short HEAD`
export GO15VENDOREXPERIMENT=1

all: build

add-deps:
	@godep save
	@rm -rf Godeps

build:
	@cd interlock && go build -a -tags 'netgo' -ldflags "-w -X github.com/rgbkrk/interlocarina/version.GitCommit=$(COMMIT) -linkmode external -extldflags -static" .

clean:
	@rm -rf interlock/interlock

image: build
	@echo Building Interlock image $(TAG)
	@docker build -t $(DOCKERHUBUSER):$(TAG) .

release: build image
	@docker push $(DOCKERHUBUSER):$(TAG)

test: clean
	@go test -v ./...

.SUFFIXES: .build
.PHONY: add-deps build clean release test
