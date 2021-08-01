postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root todoapp

dropdb:
	docker exec -it postgres12 dropdb todoapp

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/todoapp?sslmode=disable" -verbose up

test:
	go test -v -cover ./...

run:
	go run main.go

mock:
	mockgen -destination db/mock/querier.go -package mockdb github.com/adityaladwa/todoapp/db/sqlc Querier

.PHONEY :postgres createdb dropdb migrateup test run mock