package main

import (
	"fmt"
	"regexp"

	"github.com/valyala/fasthttp"
)

// Endpoint infrastructure
type Endpoint struct {
	Regexp  *regexp.Regexp
	Handler func(*Server, *fasthttp.RequestCtx)
}

func NewEndpoint(rgx string,
	handler func(*Server, *fasthttp.RequestCtx)) *Endpoint {

	endpoint := Endpoint{
		Regexp:  regexp.MustCompile(rgx),
		Handler: handler,
	}
	return &endpoint
}

// Actual endpoint creation
// Endpoints are ordered from most specific to least specific
var (
	nodes = NewEndpoint(`^\/nodes\/?$`, GetNodesHandler)
	root  = NewEndpoint(`^\/$`, RootHandler)

	// Order matters here - the first occurring match will be used
	Endpoints = []*Endpoint{
		nodes,
		root,
	}
)

// GetNodesHandler returns all nodes present
func GetNodesHandler(server *Server, ctx *fasthttp.RequestCtx) {
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
	prettyJSON, err := PrettyPrint(defaultJSON)
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		return
	}
	fmt.Fprintf(ctx, "%s\n", prettyJSON)
}

// RootHandler returns a basic message
func RootHandler(server *Server, ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "This is the root path\n")
}
