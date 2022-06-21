package model

//User model
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username" binding:"required"`
	Name     string `json:"name"`
	Email    string `form:"email" json:"email"`
	Bio      string `json:"bio"`
	Password string `json:"password" binding:"required"`
}

//UserRequest model
type UserRequest struct {
	ID       int    `json:"id"`
	Username string `json:"username" binding:"required"`
	Name     string `json:"name"`
	Email    string `form:"email" json:"email"`
	Bio      string `json:"bio"`
	Password string `json:"password" binding:"required"`
}

//UserResponse model
type UserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username" binding:"required"`
	Name     string `json:"name" `
	Email    string `form:"email" json:"email" `
	Bio      string `json:"bio" `
}
