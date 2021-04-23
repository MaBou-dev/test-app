PROJECT_NAME := $(notdir $(CURDIR))
VERSION := $(shell cat VERSION)

.DEFAULT_GOAL: help

.PHONY: help
#help: @ help menu
help:
	@grep --extended-regexp '[a-zA-Z0-9\.\-]+:.*?@ .*$$' Makefile | tr --delete '#' | sort | awk 'BEGIN {FS = ":.*?@ "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

CONTAINER_NAME := $(PROJECT_NAME)_container

#build.api: @ compile code and build container
build.api:
	docker build --target bin --tag $(CONTAINER_NAME):$(VERSION) api
	docker save --output $(CONTAINER_NAME).tar $(CONTAINER_NAME)

#run.api: @ run container
run.api:
	docker run --dettach $(CONTAINER_NAME):$(VERSION)




