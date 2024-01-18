package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/segment3d-app/segment3d-be/api"
	db "github.com/segment3d-app/segment3d-be/db/sqlc"
	"github.com/segment3d-app/segment3d-be/util"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("can't load config")
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("can't connect to database", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(&config, store)
	if err != nil {
		log.Fatal("can't create server", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("can't start server", err)
	}
}
