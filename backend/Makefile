APP_NAME=bookstore
CMD=./cmd/api
DB_CONTAINER=bookstore_postgres
DB_URL=postgres://bookstore_user:bookstore_password@localhost:5432/bookstore?sslmode=disable

.PHONY: infra run build test fmt vet tidy migrate seed logs

infra:
	docker compose -f ../infra/docker-compose.yml up -d

run:
	go run $(CMD)

build:
	go build -o bin/$(APP_NAME) $(CMD)

test:
	go test ./...

fmt:
	go fmt ./...

vet:
	go vet ./...

tidy:
	go mod tidy

migrate:
	docker exec -i $(DB_CONTAINER) psql -U bookstore_user -d bookstore < migrations/init.sql

seed:
	docker exec -i $(DB_CONTAINER) psql -U bookstore_user -d bookstore < migrations/seed.sql

logs:
	docker logs -f $(DB_CONTAINER)