package ws

type Room struct {
	ID      string            `json:"id"`
	Name    string            `json:"name"`
	Clients map[int64]*Client `json:"clients"`
}

type Hub struct {
	Rooms      map[string]*Room
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *Message
}

func NewHub() *Hub {
	return &Hub{
		Rooms:      make(map[string]*Room),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *Message, 5),
	}
}

func (hub *Hub) Run() {
	for {
		select {
		case client := <-hub.Register:
			room, ok := hub.Rooms[client.RoomId]
			if ok {
				if _, ok := room.Clients[client.ID]; !ok {
					room.Clients[client.ID] = client
				}
			}

		case client := <-hub.Unregister:
			room, ok := hub.Rooms[client.RoomId]
			if ok {
				if _, ok := room.Clients[client.ID]; ok {
					if len(room.Clients) != 0 {
						hub.Broadcast <- &Message{
							Content:  "user left the chat",
							RoomId:   client.RoomId,
							Username: client.Username,
						}
					}

					// Doubt, how are we deleting *Client value with client.ID
					delete(room.Clients, client.ID)
					close(client.Message)
				}
			}

		case message := <-hub.Broadcast:
			room, ok := hub.Rooms[message.RoomId]
			if ok {
				for _, client := range room.Clients {
					client.Message <- message
				}
			}
		}
	}
}
