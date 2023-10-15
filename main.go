package main

import (
	"fmt"
	"net/http"

	"github.com/pawarpranav83/golang-chat/pkg/websocket"
)

func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("websocket endpoint reached")

	conn, err := websocket.Upgrade(w, r)

	if err != nil {
		// Doubt
		fmt.Fprintf(w, "%+v\n", err)
	}

	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}

	// Since we create a new client, we have to register that client
	pool.Register <- client
	client.Read()
}

func setupRoutes() {
	// Creates a new pool, and initializes all the channels and maps.
	pool := websocket.NewPool()

	// ServeMux - Router
	// Socket functions are usually executed in go routines.
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})
}

func main() {
	fmt.Println("Chat Project")

	setupRoutes()

	http.ListenAndServe(":9000", nil)
}
