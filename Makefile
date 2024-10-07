run:
	@go run cmd/main.go

test:
	@go test -v ./...

test-auth:
	@go test -v ./services/aut
	
migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@, $(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down 