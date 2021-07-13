package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/adityaladwa/todoapp/api"
	"github.com/adityaladwa/todoapp/config"
	db "github.com/adityaladwa/todoapp/db/sqlc"
)

func main() {
	fmt.Println("Starting main...")
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load configuration", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to database: ", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)

	server.Start(config.ServerAddress)
}
