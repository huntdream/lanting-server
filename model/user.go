package model

//User model
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username" binding:"required"`
	Name     string `json:"name" gorm:"default:'null'"`
	Email    string `form:"email" json:"email" gorm:"default:'null'"`
	Bio      string `json:"bio" gorm:"default:'null'"`
	Password string `json:"password" binding:"required"`
}

//UserRequest model
type UserRequest struct {
	ID       int    `json:"id"`
	Username string `json:"username" binding:"required"`
	Name     string `json:"name" gorm:"default:'null'"`
	Email    string `form:"email" json:"email" gorm:"default:'null'"`
	Bio      string `json:"bio" gorm:"default:'null'"`
	Password string `json:"password" binding:"required"`
}

//UserResponse model
type UserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username" binding:"required"`
	Name     string `json:"name" gorm:"default:'null'"`
	Email    string `form:"email" json:"email" gorm:"default:'null'"`
	Bio      string `json:"bio" gorm:"default:'null'"`
}
