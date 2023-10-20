package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/pawarpranav83/golang-chat/api"
	db "github.com/pawarpranav83/golang-chat/db/sqlc"
	"github.com/pawarpranav83/golang-chat/db/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load configurations: ", err)
	}

	// Connecting to db
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}
}

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/pawarpranav83/golang-chat/pkg/websocket"
// )


// func main() {
// 	fmt.Println("Chat Project")

// 	// setupRoutes()

// 	http.ListenAndServe(":9000", nil)
// }
