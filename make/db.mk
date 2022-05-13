POSTGRES_USER	?= postgres
POSTGRES_PASS	?= pass
POSTGRES_HOST	?= localhost
POSTGRES_PORT	?= 5432
POSTGRES_DBNAME ?= mai_backend
POSTGRES_URL 	:= "postgres://$(POSTGRES_USER):$(POSTGRES_PASS)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DBNAME)?sslmode=disable"

SQLBOILER_TOOL 		:= github.com/volatiletech/sqlboiler/v4
SQLBOILER_PSQL_TOOL := github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql
SQLBOILER_CONFIG	:= ./api-server/internal/db/sqlboiler.toml

.PHONY: generate-sqlboiler
generate-sqlboiler: | $(BUILD_CACHE_DIR)
	@$(GO) build -o $(BUILD_CACHE_DIR) $(SQLBOILER_PSQL_TOOL)
	@$(GO) run $(SQLBOILER_TOOL) \
		--config $(SQLBOILER_CONFIG) \
		-p dbmodel \
		$(BUILD_CACHE_DIR)/sqlboiler-psql

MIGRATION_TOOL := github.com/golang-migrate/migrate/v4/cmd/migrate
MIGRATIONS_DIR := ./api-server/internal/db/migrations

define run-migration-tool =
	@$(GO) run -tags postgres $(MIGRATION_TOOL) \
		-source file://$(MIGRATIONS_DIR) \
		-database $(POSTGRES_URL)
endef

.PHONY: migrate-up
migrate-up:
	$(run-migration-tool) up

.PHONY: migrate-down
migrate-down:
	$(run-migration-tool) down -all

.PHONY: drop-migrations
drop-migrations:
	$(run-migration-tool) drop -f
