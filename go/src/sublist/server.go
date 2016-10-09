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

	nodes, err := server.db.GetNodes()
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}
	defaultJSON, err := nodes.ToBytes()
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}

	// Route matching
	switch {
	case anyArgGiven.MatchString(string(ctx.Path())):
		fmt.Fprintf(ctx, "Hi there, I really love %s!\n", ctx.Path()[1:])
	case root.MatchString(string(ctx.Path())):
		fmt.Fprintf(ctx, "%s\n", defaultJSON)
	default:
		ctx.Error("not found", fasthttp.StatusNotFound)
	}
}
