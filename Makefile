SHELL=/bin/bash
# make 指令
MAKE=make
# Go parameters
GO_CMD=go
GO_MOD=$(GO_CMD) mod
GO_BUILD=$(GO_CMD) build
GO_CLEAN=$(GO_CMD) clean
GO_TEST=$(GO_CMD) test
GO_GET=$(GO_CMD) get
GO_VET=$(GO_CMD) vet
GO_RUN=$(GO_CMD) run
#測試
COVER_PROFILE=cover.out

#docker
DOCKER_CMD=docker
DOCKER_BUILD=$(DOCKER_CMD) build
DOCKER_PUSH=$(DOCKER_CMD) push
DOCKER_RUN=$(DOCKER_CMD) run
DOCKER_IMAGE_NAME=request-limit-server
DOCKER_FULL_IMAGE=$(DOCKER_IMAGE_NAME)

#docker-compose
DOCKER_COMPOSE_CMD=docker-compose
DOCKER_COMPOSE_RUN=$(DOCKER_COMPOSE_CMD) up
DOCKER_COMPOSE_RUN_BACK=$(DOCKER_COMPOSE_RUN) -d
DOCKER_COMPOSE_DOWN=$(DOCKER_COMPOSE_CMD) down
#打包
BINARY_NAME=app
BINARY_UNIX=$(BINARY_NAME)_unix
VERSION=1.0.0
#版本控制
SUB_VERSION_FILE=./version_control/BUILD

#其他指令
ALL_PATH=./...

all: deps run
deps:
	$(GO_MOD) download
test:
	$(GO_TEST) -v $(ALL_PATH) -cover
test_nocache:
	$(GO_TEST) -v $(ALL_PATH) -cover -count=1
build:
	$(GO_BUILD) -o $(BINARY_NAME)
run:
	$(GO_RUN) main.go
clean:
	$(GO_CLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
pack:
	tar -cvzf $(BINARY_NAME)-v$(VERSION).tar.gz $(BINARY_NAME) ./config
docker_build:
	@echo "開始打包 Docker Image - $(DOCKER_FULL_IMAGE)"
	$(DOCKER_BUILD) -t $(DOCKER_FULL_IMAGE) .
docker_run:
	@echo "開始執行 Docker Image - $(DOCKER_FULL_IMAGE)"
	$(DOCKER_RUN) -p 8888:8888 --name limit-request-server $(DOCKER_FULL_IMAGE)
docker_compose_run:
	@echo "開始執行 Docker Compose - $(DOCKER_FULL_IMAGE)"
	$(DOCKER_COMPOSE_RUN_BACK)

docker: docker_build docker_run
docker_compose: docker_build docker_compose_run