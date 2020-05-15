# quarks-gora

## Run

1. Generate SSL keys: `scripts/gen-keys`
2. Run the HTTPS server `SERVER_KEY=server.key SERVER_CRT=server.crt PORT=4443 go run ./main.go`

## Usage

- Get env `curl -k https://localhost:4443`
- Post a command to be executed (returns 200 on success, 500 on failure): `curl -v -d "exit 1" -k https://localhost:4443`
