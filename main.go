package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/adityaladwa/todoapp/api"
	db "github.com/adityaladwa/todoapp/db/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://adityaladwa:secret@localhost:5432/todoapp?sslmode=disable"
	serverAddress = ":8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to database: ", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)

	server.Start(serverAddress)
}
