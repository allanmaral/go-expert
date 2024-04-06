include .env

migrate-create:
	migrate create -ext sql -dir sql/migrations $(name)

migrate:
	migrate -path sql/migrations -database $(DATABASE_URL) up

migrate-down:
	migrate -path sql/migrations -database $(DATABASE_URL) down

.PHONY: migrate migrate-create migrate-down

