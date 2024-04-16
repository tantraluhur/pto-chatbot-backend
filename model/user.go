package models

type User struct {
	ID       int64  `gorm:"primaryKey; autoIncrement"`
	Name     string `gorm:"not null"`
	Email    string `gorm:"not null"`
	Password []byte `gorm:"not null"`
}
