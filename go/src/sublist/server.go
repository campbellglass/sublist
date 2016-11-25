package main

import (
	"fmt"
	"regexp"

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
	var root = regexp.MustCompile(`^\/$`)
	var anyArgGiven = regexp.MustCompile(`^\/.+$`)
	var nodes = regexp.MustCompile(`^\/nodes\/?$`)

	// Route matching
	path := string(ctx.Path())
	switch {
	case root.MatchString(path):
		server.RootHandler(ctx)
	case nodes.MatchString(path):
		server.GetNodesHandler(ctx)
	case anyArgGiven.MatchString(path):
		fmt.Fprintf(ctx, "Hi there, I really love %s!\n", ctx.Path()[1:])
	default:
		ctx.Error("not found", fasthttp.StatusNotFound)
	}
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
