include .env

postgres:
	docker run --name postgres16 -p 5432:5432 -e POSTGRES_USER=${POSTGRES_USER} -e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} -d postgres:16-alpine

createdb:
	docker exec -it postgres16 createdb --username=${POSTGRES_USER} --owner=${POSTGRES_USER} ${DB_NAME}

dropdb:
	docker exec -it postgres16 dropdb ${DB_NAME}

migrate-up:
	migrate -path ./db/migrations -database ${POSTGRES_MIGRATION_URL} -verbose up

migrate-down:
	migrate -path ./db/migrations -database ${POSTGRES_MIGRATION_URL} -verbose down

.PHONY: postgres createdb dropdb