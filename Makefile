install:
	go get -v ./...
	go mod tidy

clean:
	rm -rf ./build

build: install clean
	go build -v -o ./build/hours

lint:
	go fmt ./...

test: lint
	go test -v ./...