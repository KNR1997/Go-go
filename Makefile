build:
	@go build -o bin/go-go cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/go-go

docker-up:
	@docker-compose up -d

docker-down:
	@docker-compose down

migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down