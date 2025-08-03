build:
	@go build -o bin/go-go cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/go-go