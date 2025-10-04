build:
	go build -o bin/ecom cmd/main.go

test:
	go test -v ./...

run: build
	./bin/ecom


MIGRATE := migrate
MIGRATIONS_DIR := cmd/migrate/migrations

.PHONY: migration migrate-up migrate-down

# 用法: make migration name=add_user_table
migration:
	$(MIGRATE) create -ext sql -dir $(MIGRATIONS_DIR) $(name)

migrate-up:
	go run cmd/migrate/main.go up

migrate-down:
	go run cmd/migrate/main.go down