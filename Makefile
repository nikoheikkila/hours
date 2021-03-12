install:
	go mod vendor

clean:
	rm -rf ./build

build: clean
	go build -o ./build/hours
