#
default: test

clean:
	rm -rf bin/*; rm -rf pkg/*

build:
	go install

test:
	go test

package:
		CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/app && docker build . -t adbourne/go-archetype-kafka-processor:latest