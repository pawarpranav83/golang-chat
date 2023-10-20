package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	db "github.com/pawarpranav83/golang-chat/db/sqlc"
	"github.com/pawarpranav83/golang-chat/db/util"
	"github.com/pawarpranav83/golang-chat/token"
	"github.com/pawarpranav83/golang-chat/ws"
)

// Server serves http requests
// config in the struct we require the token duration in the config obj later when creating tokens
type Server struct {
	config     util.Config
	store      *db.Store
	tokenMaker token.PasetoMaker
	hub        *ws.Hub
	router     *gin.Engine
}

// This creates a new http server and setup routing
func NewServer(config util.Config, store *db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		store:      store,
		tokenMaker: *tokenMaker,
		config:     config,
	}

	hub := ws.NewHub()
	server.hub = hub
	go hub.Run()

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {

	router := gin.Default()

	// Routes
	// The createUser is a method since we require access to the store db to process the request
	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)
	// router.GET("/ws", func(ctx *gin.Context) { server.getWebsocket(ctx, pool) })

	// Doubt
	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	authRoutes.POST("/ws/createRoom", server.InitiateRoom)
	authRoutes.GET("/ws/joinRoom/:roomId", server.JoinRoom)
	authRoutes.GET("/ws/getRooms", server.getRooms)
	authRoutes.GET("/ws/getClients/:roomId", server.getClients)

	authRoutes.GET("/users/me", server.getMe)
	authRoutes.GET("/users/:id", server.getUser)
	authRoutes.GET("/users", server.listUsers)
	authRoutes.DELETE("/users/:id", server.deleteUser)
	authRoutes.DELETE("/users/me", server.deleteMe)
	authRoutes.PATCH("/users/:id", server.updateUser)
	authRoutes.PATCH("/users/me", server.updateMe)

	authRoutes.POST("/rooms", server.createRoomUser)
	authRoutes.POST("/rooms/:roomId/:userId", server.addUsertoRoom)
	authRoutes.GET("/rooms/:roomId", server.getRoomUser)
	authRoutes.GET("/rooms/me", server.getRoomsofUser)
	authRoutes.DELETE("/rooms/:roomId", server.userLeaveRoom)

	// Add the router with all the routes attached to server.router
	server.router = router
}

// Runs the server on the given address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// func errorResponse is used to convert the error into a key value object, so that Gin can serialize it to JSON
func errorResponse(err error) gin.H {
	// The Error method in on type  error itself, its the only method on type error, it returns the error in string format
	return gin.H{"error": err.Error()}
}
