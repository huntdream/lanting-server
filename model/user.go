package model

//User model
type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username" binding:"required"`
	Name     string `json:"name"`
	Avatar   string `json:"avatar"`
	Email    string `form:"email" json:"email"`
	Bio      string `json:"bio"`
	Password string `json:"password" binding:"required"`
}

//UserResult model
type UserResult struct {
	ID       int64  `json:"id"`
	Username string `json:"username" binding:"required"`
	Name     string `json:"name" `
	Avatar   string `json:"avatar"`
	Email    string `form:"email" json:"email" `
	Bio      string `json:"bio" `
}
