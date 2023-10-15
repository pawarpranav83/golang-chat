package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Check gorilla's docs
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Upgrades http to websocket conn
func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	// For checking client (origin of websocket conn)
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return conn, nil
}
