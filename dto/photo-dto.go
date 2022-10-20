package dto

type RequestPhoto struct {
	ID       string `json:"id,omitempty"`
	Caption  string `json:"caption,omitempty" validate:"required"`
	PhotoURL string `json:"photo_url" validate:"required,url"`
	Title    string `json:"title" validate:"required"`
	UserID   string `json:"user_id,omitempty"`
}

type ResponseCreatePhoto struct {
	ID        string `json:"id,omitempty"`
	Title     string `json:"title"`
	Caption   string `json:"caption,omitempty"`
	PhotoURL  string `json:"photo_url"`
	CreatedAt int64  `json:"created_at,omitempty"`
}

type ResponseGetPhoto struct {
	ID       string `json:"id,omitempty"`
	Title    string `json:"title"`
	Caption  string `json:"caption,omitempty"`
	PhotoURL string `json:"photo_url"`
	User     struct {
		Username string `json:"username"`
		Email    string `json:"email"`
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
