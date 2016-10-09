package main

import (
	"fmt"
	"regexp"

	"github.com/valyala/fasthttp"
)

type Server struct {
}

func NewServer() *Server {
	fmt.Println("Creating a server")
	server := Server{}
	return &server
}

// Start runs a server
func (*Server) Start() {
	fmt.Println("Starting a server")
	fasthttp.ListenAndServe(":8080", router)
}

// router Routes API Requests to the appropriate handler function
func router(ctx *fasthttp.RequestCtx) {
	// Route regexp
	var root = regexp.MustCompile(`^\/$`)
	var anyArgGiven = regexp.MustCompile(`^\/.+$`)

	// Route matching
	switch {
	case anyArgGiven.MatchString(string(ctx.Path())):
		fmt.Fprintf(ctx, "Hi there, I really love %s!\n", ctx.Path()[1:])
	case root.MatchString(string(ctx.Path())):
		fmt.Fprintf(ctx, "I've got to love something...\n")
	default:
		ctx.Error("not found", fasthttp.StatusNotFound)
	}
}
