package main

import (
	"fmt"
)

func main() {
	fmt.Println("Spinning up the database")
	db := NewDatabase()

	fmt.Println("Creating a server")
	server := NewServer(db)

	fmt.Println("Here goes nothing!")
	server.Start()
}
