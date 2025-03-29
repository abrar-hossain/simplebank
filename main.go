package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/abrar-mashuk/simplebank/api"
	db "github.com/abrar-mashuk/simplebank/db/sqlc"
	"github.com/abrar-mashuk/simplebank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}

}
