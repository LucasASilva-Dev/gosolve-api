include ./config/server.mak

GIT_VERSION := $(shell git describe --tags --long --dirty --always)
COMMIT := $(shell git rev-parse HEAD)
PACKAGES := "$(shell go list ./... | grep -v tests)"
BUILD_ARGS := -a -installsuffix cgo -ldflags="-w -s -X gosolve-api/internal/version.IDENTIFIER=$(GIT_VERSION) -X gosolve-api/internal/version.COMMIT=$(COMMIT)"
SHELL := /bin/bash
APP := gosolve
DOCKER_CMD := "docker"
DOCKER_REGISTRY := "docker-hub.local"
DOCKER_IMAGE := "$(APP)-alpine:$(COMMIT)-local"
VERSION:=$(shell cat version)

install: ## Install dependencies
	@echo "..."
	@go mod tidy
	@go mod vendor
	@echo "✅"

test: ## Runs test
	@echo "..."
	@ENV=test go test ./... $(shell echo $(PACKAGES))
	@echo "✅"

coverage: ## Runs test with coverage
	@ENV=test go test ./... --coverprofile cov.txt
	@go tool cover --func=cov.txt

run: ## Run using go run
	@ENV=dev go run cmd/$(APP)/main.go -u ${HOST} -p ${PORT} -l ${LOG_LEVEL} server 

run-bin: ## Run using binary 
	./build/$(APP) -u ${HOST} -p ${PORT} -l ${LOG_LEVEL} server

build-image:
	@$(DOCKER_CMD) build --network host -t $(DOCKER_IMAGE) \
		--build-arg APP="$(APP)" \
		--build-arg VERSION="$(VERSION)" \
		-f dockerfiles/Dockerfile .
	@$(DOCKER_CMD) tag $(APP)-alpine:$(VERSION)-local $(DOCKER_REGISTRY)/${APP}/$(DOCKER_IMAGE)
	@$(DOCKER_CMD) tag $(APP)-alpine:$(VERSION)-local gosolve


run-image: build-image ## Run docker image
	@$(DOCKER_CMD) run --rm -it --network host $(DOCKER_REGISTRY)/${APP}/$(DOCKER_IMAGE) $${ARGS}

build-bin: ## Builds a GO binary
	@GOOS=$(GOOS) GOARCH=$(GOARCH) go build $(BUILD_ARGS) -o "build/$(APP)" ./cmd/$(APP)

version: ## Get current version
	@echo $(GIT_VERSION)

help: ## This help message
	@echo -e "$$(grep -hE '^\S+:.*##' $(MAKEFILE_LIST) | sed -e 's/:.*##\s*/:/' -e 's/^\(.\+\):\(.*\)/\\x1b[36m\1\\x1b[m:\2/' | column -c2 -t -s :)"
