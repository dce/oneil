.PHONY: all build clean

all: build

build:
	go build -o bin/apriltag cmd/apriltag.go

clean:
	rm -f bin/apriltag
