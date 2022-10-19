package entity

type Comment struct {
	ID        string `json:"id" gorm:"primaryKey"`
	UserID    string `json:"user_id"`
	PhotoID   string `json:"photo_id"`
	Message   string `json:"message" gorm:"not null"`
	User      *User  `gorm:"foreignKey:UserID"`
	Photo     *Photo `gorm:"foreignKey:PhotoID"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}
