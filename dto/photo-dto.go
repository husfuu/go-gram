package dto

type RequestPhoto struct {
	ID       string `json:"id,omitempty"`
	Caption  string `json:"caption,omitempty" validate:"required" example:"Looks my pretty waifu"`
	PhotoURL string `json:"photo_url" validate:"required,url"`
	Title    string `json:"title" validate:"required" example:"Waifuu Photo"`
	UserID   string `json:"user_id,omitempty"`
}

type ResponseCreatePhoto struct {
	ID        string `json:"id,omitempty" example:"1"`
	Title     string `json:"title" example:"Waifu photo"`
	Caption   string `json:"caption,omitempty" example:"Looks my pretty waifu"`
	PhotoURL  string `json:"photo_url" example:"https://photos/waifuu.png"`
	CreatedAt int64  `json:"created_at,omitempty"`
}

type ResponseGetPhoto struct {
	ID       string `json:"id,omitempty" example:"1" `
	Title    string `json:"title" example:"Waifu photo"`
	Caption  string `json:"caption,omitempty" example:"Looks my pretty waifu"`
	PhotoURL string `json:"photo_url" example:"https://photos/waifuu.png"`
	User     struct {
		Username string `json:"username" example:"kaikaruizawa" `
		Email    string `json:"email" example:"keikaruizawa@gmail.com"`
	}
	CreatedAt int64 `json:"created_at,omitempty"`
}

type ResponseUpdatePhoto struct {
	ID        string `json:"id,omitempty"`
	Title     string `json:"title"`
	Caption   string `json:"caption,omitempty"`
	PhotoURL  string `json:"photo_url"`
	UserID    string `json:"user_id,omitempty"`
	UpdatedAt int64  `json:"created_at,omitempty"`
}
