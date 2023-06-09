package auth

type SignupData struct {
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type ResponseToken struct {
	Token string `json:"token"`
}