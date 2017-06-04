SHELL := /bin/bash
CWD := $(shell pwd)
IMAGE := "jamrizzi/dns-register:latest"
SOME_CONTAINER := $(shell echo some-$(IMAGE) | sed 's/[^a-zA-Z0-9]//g')
DOCKERFILE := $(CWD)/Dockerfile

.PHONY: all
all: clean dependancies build

.PHONY: register
register:
	@go build register.go
	@echo built register

.PHONY: start
start: get realize
	@realize run

.PHONY: get
get:
	@go get
	@echo got dependancies

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
	-@rm -rf ./.realize ./register
	@echo cleaned

.PHONY: dependancies
dependancies: docker realize
	@echo fetched dependancies
.PHONY: docker
docker:
ifeq ($(shell whereis docker), $(shell echo docker:))
	@curl -L https://get.docker.com/ | bash
endif
	@echo fetched docker
.PHONY: realize
realize:
ifeq ($(shell whereis realize), $(shell echo realize:))
	@go get github.com/tockins/realize
endif
	@echo fetched realize
