.PHONY: dev build clean

all: dev

dev: build
	./2do

build: clean
	go get ./...
	go build .

test:
	go test ./...

clean:
	rm -rf 2do
