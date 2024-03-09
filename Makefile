postgres:
	docker run --name postgres15 -p 5431:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine

createdb: 
	docker exec -it postgres15 createdb --username=root --owner=root auth_jwt

dropdb: 
	docker exec -it postgres15 dropdb auth_jwt

migrateup: 
	migrate -path db/migration -database "postgresql://root:secret@localhost:5431/auth_jwt?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5431/auth_jwt?sslmode=disable" -verbose down

sqlc:
	sqlc generate

dev: 
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown dev