package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/sawilkhan/simple-bank/api"
	db "github.com/sawilkhan/simple-bank/db/sqlc"
	"github.com/sawilkhan/simple-bank/util"
)


func main(){

	config, err := util.LoadConfig(".")
	if err != nil{
		log.Fatal("cannot load config")
	}


	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil{
		log.Fatal("cannot establish connection with the database", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil{
		log.Fatal("cannot connect to db:", err)
	}
	err = server.Start(config.ServerAddress)
	if err != nil{
		log.Fatal("Cannot start server: ", err)
	}
}