package book

type Author struct {
	Id   string
	Name string
}

type Tag struct {
	Id    int8
	Label string
}
type Book struct {
	Id      string
	Title   string
	Tags    []Tag
	Authors []Author
}
