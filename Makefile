.PHONY: all build download

all: download build

download:
	go get -d ./...

build:
	go build -o ./bin/music_service

