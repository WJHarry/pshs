.PHONY: build clean install

all: build

build:
	@go build .

clean:
	@go clean -i

install:
	@go install
