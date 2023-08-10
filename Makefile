BINARY_NAME=main
.DEFAULT_GOAL := run

build:
	GOARCH=amd64 GOOS=darwin go build -o ./bin/${BINARY_NAME}-darwin ./cmd/app/main.go
	GOARCH=amd64 GOOS=linux go build -o ./bin/${BINARY_NAME}-linux ./cmd/app/main.go
	GOARCH=amd64 GOOS=windows go build -o ./bin/${BINARY_NAME}-windows ./cmd/app/main.go

run: build
	./bin/${BINARY_NAME}-darwin

clean:
	go clean
	rm ./bin/${BINARY_NAME}-darwin
	rm ./bin/${BINARY_NAME}-linux
	rm ./bin/${BINARY_NAME}-windows

test:
	go test ./...

dep:
	go mod download

vet:
	go vet ./internal/database
	go vet ./internal/entity
	go vet ./internal/repository
	go vet ./internal/service

lint:
	golangci-lint run --enable-all
