package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/vule96/sn-be/api"
	db "github.com/vule96/sn-be/db/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:P@ssword123@localhost:5432/social_network?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to database ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server ", err)
	}
}
