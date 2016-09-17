package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
	fmt.Println("This is where we'll be running and testing the server side code.")
	fmt.Println("Eventually we'll make it nice, but this is good enough for now.")
	fmt.Println("The exciting part is that it runs!")
	root := MakeTestSubList()

	bytes, _ := root.ToBytes()
	fmt.Printf("The marhalled json is as follows:\n'%s'\n", bytes)

	reformed := NewNodeFromBytes(bytes)
	fmt.Printf("The reformed object is as follows:\n'%+v'\n", reformed)
	fmt.Printf("The children are as follows:\n'%+v'\n", reformed.Children)
	fmt.Printf("A child is as follows:\n'%+v'\n", reformed.Children[0])
}
