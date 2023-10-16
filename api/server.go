package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/pawarpranav83/golang-chat/db/sqlc"
)

// Server serves http requests
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// This creates a new http server and setup routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// Routes
	// The createUser is a method since we require access to the store db to process the request
	router.POST("/users", server.createUser)
	router.GET("/users/:id", server.getUser)
	router.GET("/users", server.listUsers)
	router.DELETE("/users/:id", server.deleteUser)
	router.PATCH("/users/:id", server.updateUser)

	// Add the router with all the routes attached to server.router
	server.router = router
	return server
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
