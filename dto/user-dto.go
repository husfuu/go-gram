package dto

type RequestRegister struct {
	ID       string `json:"id,omitempty" swaggerignore:"true"`
	Username string `json:"username" validate:"required" example:"keikaruizawa"`
	Email    string `json:"email" validate:"required,email" example:"keikaruizawa@gmail.com"`
	Password string `json:"password" validate:"required,min=6" example:"password"`
	Age      int    `json:"age" validate:"required,min=8" example:"22"`
}

type Response struct {
	ID       uint   `json:"id,omitempty"`
	Username string `json:"username" example:"keikaruizawa"`
	Email    string `json:"email" example:"keikaruizawa@gmail.com"`
	Age      int    `json:"age" example:"22"`
}

type RequestLogin struct {
	Email    string `json:"email" validate:"required,email" example:"keikaruizawa@gmail.com"`
	Password string `json:"password" validate:"required,min=6" example:"password"`
}

type ResponseLogin struct {
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"`
}

// this is for swagger needs
type ExampleRequestUpdate struct {
	Username string `json:"username" example:"keikaruizawa"`
	Email    string `json:"email" example:"keikaruizawa@gmail.com"`
}

type ExampleResponseDelete struct {
	Message string `json:"message" example:"your account has been successfully deleted"`
}
