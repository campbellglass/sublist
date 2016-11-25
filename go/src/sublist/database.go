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
	return root, nil
}
