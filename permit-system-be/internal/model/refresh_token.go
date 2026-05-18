package model

import "time"

type RefreshToken struct {
	BaseModel

	UserID string

	Token string `gorm:"type:text"`

	ExpiredAt time.Time

	IsRevoked bool `gorm:"default:false"`

	User User `gorm:"foreignKey:UserID"`
}
