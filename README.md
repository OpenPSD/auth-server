# OpenPSD auth server

OpenPSD oauth2 enabled authorization server based on [dex](https://github.com/coreos/dex).

## requirements

docker

## build
Use `make build` to build the docker container.

## configure

Take a look at the `config/config.yaml` file for configuration.

## run
Use `make run` to run the docker container. By default this will create a `dex.db` sqlite database in the `./data` directory. 

The server will be available at `http://localhost:5556/dex`.

