package entity

type Photo struct {
	ID        string `json:"id" gorm:"primaryKey;not null"`
	Title     string `json:"title" gorm:"not null"`
	Caption   string `json:"caption" gorm:"null"`
	PhotoURL  string `json:"photo_url" gorm:"not null"`
	UserID    string `json:"user_id"`
	User      *User  `gorm:"foreignKey:UserID"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}
