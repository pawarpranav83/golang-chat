package ws

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

// Doubt about sync.Mutex
type Client struct {
	Conn     *websocket.Conn
	Message  chan *Message
	ID       int64
	RoomId   string `json:"roomId"`
	Username string `json:"username"`
}

type Message struct {
	Content  string `json:"content"`
	RoomId   string `json:"roomId"`
	Username string `json:"username"`
}

func (c *Client) WriteMessage() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		message, ok := <-c.Message
		if !ok {
			return
		}

		c.Conn.WriteJSON(message)
	}
}

func (c *Client) ReadMessage(hub *Hub) {
	defer func() {
		hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, p, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		message := &Message{
			Content:  string(p),
			RoomId:   c.RoomId,
			Username: c.Username,
		}

		hub.Broadcast <- message
		fmt.Printf("Message received: %+v\n", message)
	}
}
