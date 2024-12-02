package main

import (
	"database/sql"
	"log"
	"simplebank/api"
	db "simplebank/db/sqlc"
	"simplebank/util"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("can't load configuration files")
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("can't connect to the database:", err)
	}
	defer conn.Close()

	store := db.Newstore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}
