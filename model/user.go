package model

// User model
type User struct {
	ID       int64   `json:"id"`
	Username string  `json:"username" binding:"required"`
	Name     string  `json:"name,omitempty"`
	Avatar   string  `json:"avatar,omitempty"`
	Email    *string `json:"email,omitempty"`
	Bio      string  `json:"bio,omitempty"`
	Password string  `json:"password,omitempty"`
}

// UserInfo model
type UserInfo struct {
	ID       int64   `json:"id"`
	Username string  `json:"username"`
	Name     string  `json:"name,omitempty"`
	Avatar   string  `json:"avatar,omitempty"`
	Email    *string `json:"email,omitempty"`
	Bio      string  `json:"bio,omitempty"`
}
