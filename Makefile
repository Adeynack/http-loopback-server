default: build

build:
	go build -o bin/httplbs cmd/httplbs/*.go

run: build
	bin/httplbs
