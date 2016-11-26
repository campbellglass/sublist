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
	// Route URI regexp
	// Routes are ordered from most specific to least specific
	nodes := NewEndpoint(`^\/nodes\/?$`, server.GetNodesHandler)
	root := NewEndpoint(`^\/$`, server.RootHandler)

	endpoints := []*Endpoint{
		nodes,
		root,
	}

	// Route matching
	path := string(ctx.Path())
	for _, endpoint := range endpoints {
		if endpoint.Regexp.MatchString(path) {
			endpoint.Handler(ctx)
			return
		}
	}

	// Default case
	ctx.Error(fmt.Sprintf("Invalid path: '%s'", path), fasthttp.StatusNotFound)
}

// RootHandler returns a basic message
func (server *Server) RootHandler(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "This is the root path\n")
}

// GetNodesHandler returns all nodes present
func (server *Server) GetNodesHandler(ctx *fasthttp.RequestCtx) {
	nodes, err := server.db.GetNodes()
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		return
	}
	defaultJSON, err := nodes.ToBytes()
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		return
	}
	fmt.Fprintf(ctx, "%s\n", defaultJSON)
}
