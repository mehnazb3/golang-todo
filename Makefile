createdb:
	createdb --username=postgres --owner=postgres golang-interview-challenge
migrateup:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/golang-interview-challenge?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/golang-interview-challenge?sslmode=disable" -verbose down

.PHONY: migrateup migratedown