# OpenPSD auth server

OpenPSD authorization including oauth2 capabilities provided by [hydra](https://www.ory.sh/).

Work in progress!

## requirements

- docker
- docker-compose
- yarn
- go

## structure

The applications is made up of different parts.

### portal
The UI of the identity provider is written in `VueJS` and can be found in `auth-portal`

### server
The server written in `go` provides the backend api and will implment the neccessary endpoint to complement the hydra server.

### hydra
[Hydra](https://www.ory.sh/) provides all the neccessary Oauth2 enpoint.


## build
Use `make build-docker-auth-server` to build the docker container.

## run
Use `make run-docker` to run the docker container. The server will be available at `http://localhost:8000`.

## TODO
- Setup docker-compose to combine with hydra

