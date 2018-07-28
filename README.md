# OpenPSD auth server

OpenPSD authorization including oauth2 capabilities provided by [hydra](https://www.ory.sh/).

Work in progress!

## requirements

docker
docker-compose
yarn
go

## build
Use `make build-docker-auth-server` to build the docker container.

## run
Use `make run-docker` to run the docker container. The server will be available at `http://localhost:8000`.

## TODO
- Setup docker-compose to combine with hydra

