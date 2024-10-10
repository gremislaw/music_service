.PHONY: all test build docker proto

all: build test

build:
	go mod download
	go build -o ./bin/music_service

test:
	go test ./...

docker:
	sudo docker compose up -d --build
	sudo docker logs -tf music_service

rebuild:
	go mod tidy
	make build
	
proto:
	protoc --go_out=. --go_opt=paths=import \
    --go-grpc_out=. --go-grpc_opt=paths=import \
    proto/main.proto
	cd api && GOPROXY=direct go mod init github.com/gremislaw/music_service/api
	cd api && GOPROXY=direct go mod tidy