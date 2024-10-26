.PHONY: all build download

all: download build

download:
	go get -d ./...

build:
	go build -o ./bin/music_service

rebuild:
	go mod tidy
	make build

format:
	go fmt ./...

clean:
	rm -rf ./bin