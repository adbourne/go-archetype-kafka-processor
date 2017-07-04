#
default: test

clean:
	rm -rf bin/*; rm -rf pkg/*

build:
	go install

test:
	go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...

package:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/app && docker build . -t adbourne/go-archetype-kafka-processor:latest