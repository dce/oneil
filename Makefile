.PHONY: all build clean

all: build

build:
	go build -o bin/oneil cmd/oneil.go

clean:
	rm -f bin/oneil
