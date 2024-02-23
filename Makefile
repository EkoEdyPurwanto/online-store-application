# Load environment variables from app.env
include app.env

# Migrate database up
migrate-up:
	migrate -path db/migrations -database "postgresql://$(APP_DB_USER):$(APP_DB_PASSWORD)@$(APP_DB_HOST):$(APP_DB_PORT)/$(APP_DB_NAME)?sslmode=disable" -verbose up

# Migrate database down
migrate-down:
	migrate -path db/migrations -database "postgresql://$(APP_DB_USER):$(APP_DB_PASSWORD)@$(APP_DB_HOST):$(APP_DB_PORT)/$(APP_DB_NAME)?sslmode=disable" -verbose down

# Declare phony targets
.PHONY: migrate-up migrate-down
