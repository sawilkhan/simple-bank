package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/sawilkhan/simple-bank/api"
	db "github.com/sawilkhan/simple-bank/db/sqlc"
)


const (
	dbDriver = "postgres"
	dbSource = "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main(){
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil{
		log.Fatal("cannot establish connection with the database", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil{
		log.Fatal("Cannot start server: ", err)
	}
}