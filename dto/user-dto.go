package dto

type RequestRegister struct {
	ID       string `json:"id,omitempty"`
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Age      int    `json:"age" validate:"required,min=8"`
}

type Response struct {
	ID       uint   `json:"id,omitempty"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

type RequestLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type ResponseLogin struct {
	Token string `json:"token"`
}
