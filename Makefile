install:
	go get -v ./...
	go mod tidy

.PHONY: build
build: install
	go build -v -o ./build/hours

lint:
	go fmt ./...

test: lint
	go test -v ./...