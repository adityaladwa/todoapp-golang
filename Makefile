migrateup:
	migrate -path db/migration -database "postgresql://adityaladwa:secret@localhost:5432/todoapp?sslmode=disable" -verbose up

test:
	go test -v -cover ./...

run:
	go run main.go

mock:
	mockgen -destination db/mock/store.go -package mockdb github.com/adityaladwa/todoapp/db/sqlc Store

.PHONEY :migrateup test run mock