package main

import (
	"fmt"
	"net/http"
)

type Server struct {
}

func NewServer() *Server {
	fmt.Println("Creating a server")
	server := Server{}
	return &server
}

// Start runs a server
// I'm following this tutorial:
// https://golang.org/doc/articles/wiki/
func (*Server) Start() {
	fmt.Println("Starting a server")
	http.HandleFunc("/", router)
	http.ListenAndServe(":8080", nil)
}

func router(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
