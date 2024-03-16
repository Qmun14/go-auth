DB_URL=postgresql://root:secret@localhost:5431/auth_jwt?sslmode=disable

postgres:
	docker run --name postgres15 -p 5431:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine

createdb: 
	docker exec -it postgres15 createdb --username=root --owner=root auth_jwt

dropdb: 
	docker exec -it postgres15 dropdb auth_jwt

migrateup: 
	migrate -path db/migration -database "${DB_URL}" -verbose up

migrateup1: 
	migrate -path db/migration -database "${DB_URL}" -verbose up 1

migratedown:
	migrate -path db/migration -database "${DB_URL}" -verbose down

migratedown1:
	migrate -path db/migration -database "${DB_URL}" -verbose down 1

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

migrate_version:
	migrate -path db/migration -database "${DB_URL}" version

migrate_force:
	migrate -path db/migration -database "${DB_URL}" force $(version)

sqlc:
	sqlc generate

dev: 
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown new_migration migrate_version migrate_force dev