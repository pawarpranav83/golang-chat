package ws

// import "fmt"

// type Pool struct {
// 	Register   chan *Client
// 	Unregister chan *Client
// 	Clients    map[*Client]bool
// 	Broadcast  chan Message
// }

// func NewPool() *Pool {
// 	return &Pool{
// 		Register:   make(chan *Client),
// 		Unregister: make(chan *Client),
// 		Clients:    make(map[*Client]bool),
// 		Broadcast:  make(chan Message),
// 	}
// }

// func (pool *Pool) Start() {
// 	// Doubt why is for loop used
// 	for {
// 		select {
// 		case client := <-pool.Register:
// 			pool.Clients[client] = true
// 			fmt.Println("size of connection pool:", len(pool.Clients))
// 			for client := range pool.Clients {
// 				fmt.Println(client)
// 				client.Conn.WriteJSON(Message{Type: 1, Body: "New User Joined..."})
// 			}
// 			break

// 		case client := <-pool.Unregister:
// 			delete(pool.Clients, client)
// 			fmt.Println("size of connection pool:", len(pool.Clients))
// 			for client := range pool.Clients {
// 				client.Conn.WriteJSON(Message{Type: 1, Body: "User Disconnected..."})
// 			}
// 			break

// 		case message := <-pool.Broadcast:
// 			fmt.Println("Sending msg to all clients in the pool")
// 			for client := range pool.Clients {
// 				if err := client.Conn.WriteJSON(message); err != nil {
// 					fmt.Println(err)
// 					return
// 				}
// 			}
// 		}
// 	}
// }
