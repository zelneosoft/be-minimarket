package authentication

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type NewPasswordRequest struct {
	Email       string `json:"email"`
	NewPassword string `json:"new_password"`
}
