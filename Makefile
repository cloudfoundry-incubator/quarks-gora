all: test build

build:
	go build

test:
	go test

gen-keys: server.crt server.key

server.crt:
	./scripts/gen-keys

server.key:
	./scripts/gen-keys

run: server.crt server.key
	@echo "Running on port 4443"
	SERVER_KEY=server.key SERVER_CRT=server.crt PORT=4443 go run ./main.go