.PHONY: dev migrate-up migrate-down create-migration install-deps

dev:
	@go run main.go

migrate-up:
	@migrate -database "postgres://postgres:postgres@localhost:5432/rmmbrit?sslmode=disable" -path db/migrations up

migrate-down:
	@migrate -database "postgres://postgres:postgres@localhost:5432/rmmbrit?sslmode=disable" -path db/migrations down

create-migration:
	@test -n "$(mig-name)" || (echo ">> migration name is not set Eg: 'make mig-name=something-something create-migration'" ; exit 1)
	@migrate create -ext sql -dir db/migrations -seq $(mig-name)

install-deps:
	./scripts/install-gomigrate.sh

