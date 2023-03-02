include infrastructure/local/.env

GOOSE_RUN := go run github.com/pressly/goose/cmd/goose@latest
JET_RUN:= go run github.com/go-jet/jet/v2/cmd/jet@latest
POSTGRES_URI = "postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL_MODE)"


.PHONY:up
up:
	docker compose -f ./infrastructure/local/docker-compose.yaml up

.PHONY: add-migration
add-migration:
	$(GOOSE_RUN) -dir db/migrations/goose postgres $(POSTGRES_URI) create new-migration sql

.PHONY:migrate-up
migrate-up:
	$(GOOSE_RUN) -dir db/migrations/goose postgres $(POSTGRES_URI) up

.PHONY:gen-db
gen-db:
	$(JET_RUN) -dsn=$(POSTGRES_URI) -schema=public -ignore-tables goose_db_version -path=db/gen
