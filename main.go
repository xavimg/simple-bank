package main

import (
	"database/sql"
	"log"

	"simple-bank/api"
	"simple-bank/config"
	db "simple-bank/db/sqlc"

	_ "github.com/lib/pq"
)

func main() {
	config, err := config.LoadENVConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	if err := server.StartRESTServer(config.AddrServer); err != nil {
		log.Fatal("cannot start server")
		return
	}

	return
}
