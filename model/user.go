package models

type User struct {
	ID       int64  `json:"id" gorm:"primaryKey; autoIncrement"`
	Name     string `json:"name" gorm:"not null"`
	Email    string `json:"email" gorm:"not null"`
	Password []byte `json:"-" gorm:"not null"`
}
