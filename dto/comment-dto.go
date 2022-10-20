package dto

type RequestComment struct {
	ID      string `json:"id"`
	Message string `json:"message" validate:"required"`
	UserID  string `json:"user_id" validate:"required"`
	PhotoID string `json:"photo_id" validate:"required"`
}

type RequestCommentUpdate struct {
	ID      string `json:"id"`
	UserID  string `json:"user_id" validate:"required"`
	Message string `json:"message" validate:"required"`
}

type ResponseCreateComment struct {
	ID      string `json:"id"`
	Message string `json:"message"`
	PhotoID string `json:"photoID"`
	UserID  string `json:"user_id"`
}

type ResponseGetComment struct {
	ID      string `json:"id"`
	Message string `json:"message"`
	UserID  string `json:"user_id"`
	PhotoID string `json:"photo_id"`
	Photo   struct {
		ID       string `json:"id"`
		Caption  string `json:"caption"`
		PhotoURL string `json:"photo_url"`
		Title    string `json:"title"`
		UserID   string `json:"user_id"`
	} `json:"photo"`
	User struct {
		ID       string `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
	} `json:"user"`
}
