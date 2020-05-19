# quarks-gora

This app is used in testing of the Quarks Operator.
It's meant to be lightweight and flexible enough to facilitate fast testing on a variety of usecases.
As a BOSH release it should be used as one package in one job. It can also be used in a docker image.

## Features

- http server with ssl
- accepts the following:
  - GET / (returns env)
  - POST / (runs the request body in bash and returns a 200 for exit code 0, 500 otherwise, STDOUD/ERR in response body)

## Run

1. Generate SSL keys: `scripts/gen-keys`
2. Run the HTTPS server `SERVER_KEY=server.key SERVER_CRT=server.crt PORT=4443 go run ./main.go`

## Usage

- Get env `curl -k https://localhost:4443`
- Post a command to be executed (returns 200 on success, 500 on failure): `curl -v -d "exit 1" -k https://localhost:4443`
