package entity

type SocialMedia struct {
	ID             string `json:"id" gorm:"primaryKey"`
	Name           string `json:"name" gorm:"not null"`
	SocialMediaUrl string `json:"social_media_url" gorm:"unique;not null"`
	UserID         string `json:"user_id"`
	User           User   `json:"user"`
	CreatedAt      int64  `json:"created_at"`
	UpdatedAt      int64  `json:"updated_at"`
}
