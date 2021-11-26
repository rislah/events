.PHONY: lint fmt list-all-linters

DOCKER_WORKDIR=$(shell pwd)
DOCKER_USER=$(shell id -u)
DOCKER_GROUP=$(shell id -g)

DOCKER_LINTER_CMD=docker run -v $(DOCKER_WORKDIR):$(DOCKER_WORKDIR) -w $(DOCKER_WORKDIR) uber/prototool:latest

fmt:
	$(DOCKER_LINTER_CMD) prototool format schema --overwrite

list-all-linters:
	$(DOCKER_LINTER_CMD) prototool lint --list-all-linters

generate:
	$(DOCKER_LINTER_CMD) prototool generate schema

lint:
	$(DOCKER_LINTER_CMD) prototool lint schema