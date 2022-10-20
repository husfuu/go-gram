package dto

type RequestComment struct {
	ID      string `json:"id"`
	Message string `json:"message"`
	UserID  string `json:"user_id"`
	PhotoID string `json:"photo_id"`
}

type RequestCommentUpdate struct {
	ID      string `json:"id"`
	UserID  string `json:"user_id"`
	Message string `json:"message"`
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

// {
// 	"data": [
// 	  {
// 		"created_at": "string",
// 		"id": 0,
// 		"message": "string",
// 		"photo": {
// 		  "caption": "string",
// 		  "id": 0,
// 		  "photo_url": "string",
// 		  "title": "string",
// 		  "user_id": 0
// 		},
// 		"photo_id": 0,
// 		"user": {
// 		  "email": "string",
// 		  "id": 0,
// 		  "username": "string"
// 		},
// 		"user_id": 0
// 	  }
// 	],
// 	"status": 0
//   }
