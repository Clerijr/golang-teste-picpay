include .env

run:
	docker compose up

migrate-up:
	migrate -path ./db/migrations -database ${POSTGRES_MIGRATION_URL} -verbose up

migrate-down:
	migrate -path ./db/migrations -database ${POSTGRES_MIGRATION_URL} -verbose down
