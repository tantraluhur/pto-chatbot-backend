package models

type AccessToken struct {
	ID          int64  `gorm:"primaryKey; autoIncrement"`
	AccessToken string `json:"access_token" gorm:"not null"`
}
