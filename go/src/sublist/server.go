package main

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

type Server struct {
	db *Database
}

func NewServer(db *Database) *Server {
	return &Server{
		db: db,
	}
}

// Start runs a server
func (server *Server) Start() {
	fmt.Println("Running a server on localhost port 8080")
	fasthttp.ListenAndServe(":8080", server.Route)
}

// Route Routes API Requests to the appropriate handler function
func (server *Server) Route(ctx *fasthttp.RequestCtx) {

	// Route matching
	path := string(ctx.Path())
	for _, endpoint := range Endpoints {
		if endpoint.Regexp.MatchString(path) {
			endpoint.Handler(server, ctx)
			return
		}
	}

	// Default case
	ctx.Error(fmt.Sprintf("Invalid path: '%s'", path), fasthttp.StatusNotFound)
}
