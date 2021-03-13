install:
	go get -t ./...

clean:
	rm -rf ./build

build: install clean
	go build -v -o ./build/hours

test:
	go test -v ./...