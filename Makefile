.PHONY: lint fmt check update

DOCKER_WORKDIR=$(shell pwd)
DOCKER_USER=$(shell id -u)
DOCKER_GROUP=$(shell id -g)
DOCKER_LINTER_CMD=docker run -v $(DOCKER_WORKDIR):$(DOCKER_WORKDIR) -w $(DOCKER_WORKDIR) risla8/prototool:latest

update: check lint fmt generate

fmt:
	$(DOCKER_LINTER_CMD) prototool format schema --overwrite

generate:
	$(DOCKER_LINTER_CMD) prototool generate schema

lint:
	$(DOCKER_LINTER_CMD) prototool lint schema

check:
	$(DOCKER_LINTER_CMD) prototool break check --git-branch master