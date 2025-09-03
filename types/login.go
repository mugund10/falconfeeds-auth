package types

// Request body for Login
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Response body for Login
type LoginResponse struct {
	Token string `json:"token"`
}
