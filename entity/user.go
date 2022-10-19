package entity

type User struct {
	ID        string `json:"id" gorm:"primaryKey;not null"`
	Username  string `json:"username" gorm:"unique;not null"`
	Email     string `json:"email" gorm:"unique;not null"`
	Age       int    `json:"age" gorm:"not null"`
	Password  string `json:"password" gorm:"not null"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}
