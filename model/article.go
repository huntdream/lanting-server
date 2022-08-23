package model

//Article model
type Article struct {
	ID       int64  `json:"id"`
	Title    string `json:"title"`
	AuthorId int64  `json:"authorId"`
	Content  string `json:"content"`
	Excerpt  string `json:"excerpt"`
	CanEdit  bool   `json:"canEdit"`
	Author   struct {
		ID       int64  `json:"id"`
		Username string `json:"username"`
	} `json:"author"`
	Base
}

// FeedItem Feed Item
type FeedItem struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	Excerpt string `json:"excerpt"`
}
