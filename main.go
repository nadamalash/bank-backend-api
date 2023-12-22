package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/nadamalash/bank-backend/api"
	db "github.com/nadamalash/bank-backend/db/sqlc"
	"github.com/nadamalash/bank-backend/util"
)

// Entry point to run the server
func main() {
	config, err := util.LoadConfig(".") // load config from the current directory
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err) //fatal error is a helper function to log the error and exit the program
	}
}
