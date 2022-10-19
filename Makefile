dep:
	go mod tidy
	go mod download

run: dep
	go run main.go

docker-build:

docker-run: docker-build

build:
	go build -o main.app

test: build
	go test ./...

linter: test
	golangci-lint run ./...

test-all: linter
