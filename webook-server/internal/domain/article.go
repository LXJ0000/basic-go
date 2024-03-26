package domain

type Article struct {
	Title    string
	Content  string
	AuthorId int64
}

type Author struct {
	Id       int64
	UserName string
}
