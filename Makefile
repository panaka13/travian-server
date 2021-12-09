.PHONY: clean build all

all: clean tidy vendor build

tidy:
	go mod tidy

vendor:
	go mod vendor

clean:
	rm -rf build
	rm -rf vendor

build:
	go build -o build/main server/main.go
