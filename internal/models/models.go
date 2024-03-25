package models

type File struct {
	Path string
	Type string
}

type Author struct {
	Name string
}

type Book struct {
	ID        string
	Title     string
	Authors   []Author
	Tags      []string
	Resources []File
}
