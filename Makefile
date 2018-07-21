CURRENT_DIR = $(shell pwd)
DOCKER_VERSION := $(shell docker --version 2>/dev/null)
DOCKER_COMPOSE_VERSION := $(shell docker-compose --version 2>/dev/null)

all:
ifndef DOCKER_VERSION
    $(error "command docker is not available, please install Docker")
endif
# ifndef DOCKER_COMPOSE_VERSION
#     $(error "command docker-compose is not available, please install Docker")
# endif

build:
	docker build -t openpsd/auth-server:latest .

run:
	docker run -v $(CURRENT_DIR)/data:/data -v $(CURRENT_DIR)/config:/config -p 5556:5556 openpsd/auth-server