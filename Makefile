#
default: build test

clean:
	rm -rf bin/app

build:
	./scripts/build-with-toolchain.sh

test:
	./scripts/test-recursively.sh

package:
	GOOS=linux go build -o bin/app; docker build . -t adbourne/go-archetype-kafka:latest