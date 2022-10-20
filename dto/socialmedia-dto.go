package dto

type RequestSocialMedia struct {
	ID             string `json:"id"`
	Name           string `json:"name" validate:"required"`
	SocialMediaUrl string `json:"social_media_url" validate:"required"`
	UserID         string `json:"user_id" validate:"required"`
}

type ResponseCreateSocialMedia struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
	UserID         string `json:"user_id"`
	CreatedAt      int64  `json:"created_at,omitempty"`
	UpdatedAt      int64  `json:"updated_at,omitempty"`
}

type ResponseGetSocialMedias struct {
	ID             string `json:"id"`
	CreatedAt      int64  `json:"created_at,omitempty"`
	UpdatedAt      int64  `json:"updated_at,omitempty"`
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
	UserID         string `json:"user_id"`
	User           struct {
		ID              string `json:"id"`
		Username        string `json:"username"`
		ProfileImageUrl string `json:"profile_image_url"`
	} `json:"user"`
}
