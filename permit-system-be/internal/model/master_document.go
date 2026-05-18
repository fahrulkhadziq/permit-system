package model

type MasterDocument struct {
	BaseModel

	Code        string `gorm:"unique;not null"`
	Name        string `gorm:"not null"`
	Description string

	IsActive bool `gorm:"default:true"`
}
