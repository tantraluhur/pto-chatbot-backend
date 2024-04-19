package models

type AccessToken struct {
	ID          int64  `gorm:"primaryKey; autoIncrement"`
	AccessToken string `json:"access_token" gorm:"not null"`
}

type RefreshToken struct {
	ID           int64  `gorm:"primaryKey; autoIncrement"`
	RefreshToken string `json:"refresh_token" gorm:"not null"`
}
