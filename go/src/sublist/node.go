package main

import (
	"encoding/json"
)

type Node struct {
	Description string
	Children    []*Node
}

func NewNode(description string) *Node {
	return &Node{
		Description: description,
		// Children implicitly initialized to nil slice
	}
}

func (parent *Node) AddChild(child *Node) {
	parent.Children = append(parent.Children, child)
}

func (node *Node) ToBytes() ([]byte, error) {
	bytes, err := json.Marshal(node)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func NewNodeFromBytes(bytes []byte) *Node {
	var node Node
	json.Unmarshal(bytes, &node)
	return &node
}

func MakeTestSubList() *Node {
	root := NewNode("Root Node")

	kitchen := NewNode("Kitchen")
	root.AddChild(kitchen)

	thing1 := NewNode("Silverware")
	thing2 := NewNode("Cake pans")
	thing3 := NewNode("Drying rack")
	kitchen.AddChild(thing1)
	kitchen.AddChild(thing2)
	kitchen.AddChild(thing3)

	bedroom := NewNode("Bedroom")
	root.AddChild(bedroom)

	return root
}
