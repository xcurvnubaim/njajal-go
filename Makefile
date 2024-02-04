ifneq (,$(wildcard ./.env))
    include .env
    export
endif
GOPATH=$(HOME)/go
PATH := $(PATH):$(GOPATH)/bin

migrate:
	migrate -path db/migrations -database "$(DB_URL)" -verbose up

migrate-down:
	migrate -path db/migrations -database "$(DB_URL)" -verbose down

new_migration:
	migrate create -ext sql -dir db/migrations -seq $(name)

db_docs:
	dbdocs build doc/db.dbml

db_schema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml

sqlc:
	sqlc generate

test:
	go test -v -cover -short ./...

server:
	go run main.go

dev:
	air

print:
	echo $1 

.PHONY: migrate migrate-down new_migration db_docs db_schema sqlc test server