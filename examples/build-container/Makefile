CWD := $(abspath $(patsubst %/,%,$(dir $(abspath $(lastword $(MAKEFILE_LIST))))))

# used for local development
all: build

build:
	g++ -static src/main.cpp -o out/static

clean:
	rm out/static

# used for containerized development
build-builder-image:
	docker build -t builder:latest -f Dockerfile.builder .

containerized-build-app:
	docker run --rm -v $(CWD):/app builder:latest make

build-application-image:
	docker build -t static:latest -f Dockerfile.application .

run-application-image:
	docker run --rm static:latest 

.PHONY: all build build-builder-image containerized-build-app build-application-image run-application-image