migrateup:
	migrate -path db/migration -database "postgresql://adityaladwa:secret@localhost:5432/todoapp?sslmode=disable" -verbose up

test:
	go test -v -cover ./...

run:
	go run main.go

.PHONEY :migrateup test run