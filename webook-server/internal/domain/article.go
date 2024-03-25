package domain

type Article struct {
	Title   string
	Content string
	Author
}

type Author struct {
	Id       int64
	UserName string
}
