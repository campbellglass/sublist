package main

type Database struct {
	root *Node
}

func NewDatabase() *Database {
	database := Database{
		root: MakeTestSubList(),
	}
	return &database
}

func (db *Database) GetNodes() (*Node, error) {
	return db.root, nil
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
