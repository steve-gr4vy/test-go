.PHONY: test build clean docker-test server pre-commit

test:
	go test -v ./...

deps:
	go mod tidy
	go mod download

clean:
	rm -rf ./build

build: clean
	go build -o build/main src/main.go

docker-test:
	docker compose run --rm app make test

server:
	docker compose up