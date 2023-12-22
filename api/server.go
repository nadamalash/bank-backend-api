package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/nadamalash/bank-backend/db/sqlc"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	store  *db.Store   // store is a pointer to the database store
	router *gin.Engine // router is to route the requests to the right handler
}

// NewServer creates a new HTTP server and setup routing.
func NewServer(store *db.Store) *Server {

	server := &Server{store: store}
	router := gin.Default()
	// add routes to router
	router.POST("/accounts", server.createAccount) //server.createAccount is a handler function
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)
	//	router.POST("/transfers", server.createTransfer)

	server.router = router
	return server
}

// Start runs the HTTP server on a specific address.
// This is a blocking function, so it will return an error if the server fails to start
// Otherwise, it will block the current goroutine and keep the server running
// The server will be running until we kill the process to listen to the requests
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// errorResponse is a helper function to return an error response in JSON format
// The response body will be like this:
//
//	{
//		"error": "detailed error message"
//	}
//
// gin.H is a shortcut for map[string]interface{}
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
