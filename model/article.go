package model

// Article model
type Article struct {
	ID         int64  `json:"id"`
	Title      string `json:"title"`
	AuthorId   int64  `json:"authorId"`
	Content    string `json:"content"`
	Text       string `json:"text,omitempty"`
	Excerpt    string `json:"excerpt"`
	CanEdit    bool   `json:"canEdit"`
	Visibility int64  `json:"visibility"`
	Author     User   `json:"author" binding:"-"`
	Base
}

type ArticleNode struct {
	Type     string
	Children []ArticleNode
	Src      string
}

type ArticleContent struct {
	Root ArticleNode
}

type ArticleMedia struct {
	Src  string
	Type string
}
