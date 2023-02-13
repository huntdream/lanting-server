package model

//User model
type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username" binding:"required"`
	Name     string `json:"name,omitempty"`
	Avatar   string `json:"avatar"`
	Email    string `form:"email,omitempty" json:"email"`
	Bio      string `json:"bio"`
	Password string `json:"password,omitempty"`
}
