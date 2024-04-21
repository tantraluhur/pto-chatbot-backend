package models

type AccessToken struct {
	ID          string `json:"id" gorm:"primaryKey"`
	AccessToken string `json:"access_token" gorm:"not null"`
}

type RefreshToken struct {
	ID           string `json:"id" gorm:"primaryKey"`
	RefreshToken string `json:"refresh_token" gorm:"not null"`
}
