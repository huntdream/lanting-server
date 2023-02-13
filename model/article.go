package model

//Article model
type Article struct {
	ID         int64  `json:"id"`
	Title      string `json:"title"`
	AuthorId   int64  `json:"authorId"`
	Content    string `json:"content"`
	Excerpt    string `json:"excerpt"`
	CanEdit    bool   `json:"canEdit"`
	Visibility int64  `json:"visibility"`
	Author     User   `json:"author"`
	Base
}

// FeedItem Feed Item
type FeedItem struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	Excerpt string `json:"excerpt"`
}
