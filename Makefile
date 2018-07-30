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

build-all: build-auth-server build-auth-portal-prod
	echo "build all"

build-auth-server:
	GOOS=linux GOARCH=amd64 go build -o docker/auth-server/auth-server ./server/

build-auth-portal-prod:
	cd auth-portal && yarn install && yarn run build && rm -rf ../docker/auth-server/web && mv dist ../docker/auth-server/web

run: stop
	cd docker && HYDRA_VERSION=v1.0.0-beta.4 docker-compose up --build -d && docker-compose logs -f

stop:
	cd docker && HYDRA_VERSION=v1.0.0-beta.4 docker-compose kill
	cd docker && HYDRA_VERSION=v1.0.0-beta.4 docker-compose rm -f

clean:
	rm -f docker/auth-server/auth-server
	rm -f docker/auth-server/web