package model

//AuthRequest request type for auth
type AuthRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

//AuthResponse response for auth
type AuthResponse struct {
	User
	Token string `json:"token"`
}
