SHELL := /bin/bash
ARGS := ""
CWD := $(shell pwd)
IMAGE := "codejamninja/dns-register:latest"
SOME_CONTAINER := $(shell echo some-$(IMAGE) | sed 's/[^a-zA-Z0-9]//g')
DOCKERFILE := $(CWD)/Dockerfile

.PHONY: all
all: clean build

.PHONY: install
start: install
	@go build dns-register.go
	@$$GOBIN/dns-register $(ARGS)

.PHONY: watch
watch: install
	@realize start

.PHONY: install
install: get
.PHONY: get
get:
	@go get

.PHONY: build
build:
	@docker build -t $(IMAGE) -f $(DOCKERFILE) $(CWD)
	@echo built $(IMAGE)

.PHONY: pull
pull:
	@docker pull $(IMAGE)
	@echo pulled $(IMAGE)

.PHONY: push
push:
	@docker push $(IMAGE)
	@echo pushed $(IMAGE)

.PHONY: run
run:
	@docker run --name $(SOME_CONTAINER) --rm $(IMAGE)
	@echo ran $(IMAGE)

.PHONY: ssh
ssh:
	@dockssh $(IMAGE)

.PHONY: essh
essh:
	@dockssh -e $(SOME_CONTAINER)

.PHONY: clean
clean:
	-@rm -rf ./.realize ./dns-register
	@echo cleaned
