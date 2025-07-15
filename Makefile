.PHONY: run-db-service build-db-service run-api-service build-api-service tidy migrate

run-db-service:
	go run ./cmd/db-service

build-db-service:
	go build -o bin/db-service ./cmd/db-service

run-api-service:
	go run ./cmd/api-service

build-api-service:
	go build -o bin/api-service ./cmd/api-service

tidy:
	go mod tidy

MIGRATE_DB_URL=postgres://user:password@localhost:5432/checklist?sslmode=disable

migrate-up:
	migrate -path migrations -database "$(MIGRATE_DB_URL)" up

migrate-down:
	migrate -path migrations -database "$(MIGRATE_DB_URL)" down

migrate-force-clean:
	migrate -path migrations -database "$(MIGRATE_DB_URL)" force 0
	migrate -path migrations -database "$(MIGRATE_DB_URL)" drop

up:
	docker-compose up -d

down:
	docker-compose down