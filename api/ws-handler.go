package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/pawarpranav83/golang-chat/token"
	"github.com/pawarpranav83/golang-chat/ws"
)

type createRoomRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Creating Room
func (server *Server) InitiateRoom(ctx *gin.Context) {
	var req createRoomRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	server.hub.Rooms[req.ID] = &ws.Room{
		ID:      req.ID,
		Name:    req.Name,
		Clients: make(map[int64]*ws.Client),
	}
	fmt.Println(server.hub, server.hub.Rooms[req.ID])

	ctx.JSON(http.StatusOK, req)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// Later edit so that only fronted is allowed, see #11:47
	CheckOrigin: func(r *http.Request) bool { return true },
}

// User Joining Room
func (server *Server) JoinRoom(ctx *gin.Context) {
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	roomId := ctx.Param("roomId")
	user := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	client := &ws.Client{
		Conn:     conn,
		Message:  make(chan *ws.Message, 10),
		ID:       user.ID,
		RoomId:   roomId,
		Username: user.Username,
	}

	msg := &ws.Message{
		Content:  "New User Joined",
		RoomId:   roomId,
		Username: user.Username,
	}

	// Register new client throught the register channel
	server.hub.Register <- client

	// Broadcast the message
	server.hub.Broadcast <- msg

	// writeMessage()
	go client.WriteMessage()

	// readMessage()
	client.ReadMessage(server.hub)
}

type getRoomsResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Get all rooms created till now
func (server *Server) getRooms(ctx *gin.Context) {
	rooms := make([]getRoomsResponse, 0)

	for _, r := range server.hub.Rooms {
		rooms = append(rooms, getRoomsResponse{
			ID:   r.ID,
			Name: r.Name,
		})
	}

	ctx.JSON(http.StatusOK, rooms)
}

type getClientsResponse struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
}

// Get all clients in the particular room
func (server *Server) getClients(ctx *gin.Context) {
	var clients []getClientsResponse
	roomId := ctx.Param("roomId")

	if _, ok := server.hub.Rooms[roomId]; !ok {
		clients = make([]getClientsResponse, 0)
		ctx.JSON(http.StatusOK, clients)
	}

	for _, cl := range server.hub.Rooms[roomId].Clients {
		clients = append(clients, getClientsResponse{
			ID:       cl.ID,
			Username: cl.Username,
		})
	}
	ctx.JSON(http.StatusOK, clients)
}

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/pawarpranav83/golang-chat/websocket"
// )

// func serveWs(pool *websocket.Pool, w gin.ResponseWriter, r *http.Request) {
// 	fmt.Println("websocket endpoint reached")

// 	conn, err := websocket.Upgrade(w, r)

// 	if err != nil {
// 		// Doubt
// 		fmt.Fprintf(w, "%+v\n", err)
// 	}

// 	client := &websocket.Client{
// 		Conn: conn,
// 		Pool: pool,
// 	}

// 	// Since we create a new client, we have to register that client
// 	pool.Register <- client
// 	client.Read()
// }

// // type joinRoomRequest struct {
// // 	RoomId int64 `uri:"roomId" binding:"required,min=1"`
// // }

// func (server *Server) getWebsocket(ctx *gin.Context, pool *websocket.Pool) {
// 	// var req joinRoomRequest
// 	// if err := ctx.ShouldBindUri(&req); err != nil {
// 	// 	ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 	// 	return
// 	// }

// 	// user, _ := server.store.GetUserbyUsername(ctx, ctx.MustGet(authorizationPayloadKey).(*token.Payload).Username)

// 	// _, err := server.store.GetRoomuser(ctx, db.GetRoomuserParams{
// 	// 	RoomID: req.RoomId,
// 	// 	UserID: user.ID,
// 	// })

// 	// if err != nil {
// 	// 	ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 	// 	return
// 	// }

// 	w, r := ctx.Writer, ctx.Request
// 	serveWs(pool, w, r)
// }
