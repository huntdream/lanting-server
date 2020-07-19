package model

//AuthRequest request type for auth
type AuthRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

//AuthResponse response for auth
type AuthResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Bio      string `json:"bio"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}
