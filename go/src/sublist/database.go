package main

type Database struct{}

func NewDatabase() *Database {
	return &Database{}
}

func (db *Database) GetNodes() (*Node, error) {
	return MakeTestSubList(), nil
}
