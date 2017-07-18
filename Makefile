#
default: build test

clean:
	rm -rf bin/*; rm -rf pkg/*

build:
	./scripts/build-with-toolchain.sh

test:
	./scripts/test-recursively.sh

package:
	GOOS=linux go build -o bin/app; docker build . -t adbourne/go-archetype-kafka:latest