include .envrc
MIGRATIONS_PATH = ./cmd/migrate/migrations

.PHONY: migration
migration:
	@migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(filter-out $@,$(MAKECMDGOALS))

.PHONY: migrate-up
migrate-up:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(AUTH_DB_ADDR) up

.PHONY: migrate-down
migrate-down:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(AUTH_DB_ADDR) down 1

.PHONY: migrate-down-all
migrate-down-all:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(AUTH_DB_ADDR) down

.PHONY: seed
seed: 
	@go run cmd/migrate/seed/main.go

.PHONY: gen-docs
gen-docs:
	@swag init -g ./api/main.go -d cmd,internal,docs/swagger && swag fmt