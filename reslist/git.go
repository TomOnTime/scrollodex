package main

func NewGit(c GitConfig) (Databaser, error) {
	db := &gitHandle{}
	return db
}

type gitHandle struct {
	url string
}

type GitConfig struct {
	url string
}

func CategoryStore(id int, data dex.Category) error {}
func CategoryList() ([]dex.Category, error)         {}
