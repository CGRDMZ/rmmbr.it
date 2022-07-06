.PHONY: dev migrate-up migrate-down create-migration install-deps

dev:
	@go run main.go

migrate-up:
	@migrate -database "postgres://postgres:postgres@localhost:5432/rmmbrit?sslmode=disable" -path db/migrations up

migrate-down:
	@migrate -database "postgres://postgres:postgres@localhost:5432/rmmbrit?sslmode=disable" -path db/migrations down

create-migration:
	@scripts/create-migration.sh

install-deps:
	./scripts/install-gomigrate.sh

