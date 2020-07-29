package model

//Article model
type Article struct {
	ID      int64  `json:"id" gorm:"PRIMARY_KEY"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Excerpt string `json:"excerpt" gorm:"-"`
	Base
}
