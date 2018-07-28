CURRENT_DIR = $(shell pwd)
DOCKER_VERSION := $(shell docker --version 2>/dev/null)
DOCKER_COMPOSE_VERSION := $(shell docker-compose --version 2>/dev/null)
PORT = 8000

all:
ifndef DOCKER_VERSION
    $(error "command docker is not available, please install Docker")
endif
ifndef DOCKER_COMPOSE_VERSION
    $(error "command docker-compose is not available, please install Docker")
endif

build-linux:
	GOOS=linux GOARCH=amd64 go build

build-docker-auth-server: build-auth-portal-prod
	GOOS=linux GOARCH=amd64 go build -o docker/auth-server/auth-server ./server/
	cd docker/auth-server && docker build -t openpsd/auth-server:latest .
	rm -f docker/auth-server/auth-server

build-auth-portal-prod:
	cd auth-portal && yarn install && yarn run build && rm -rf ../docker/auth-server/web && mv dist ../docker/auth-server/web

run-docker:
	docker run -p $(PORT):8000 openpsd/auth-server
