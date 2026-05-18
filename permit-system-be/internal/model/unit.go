package model

type Unit struct {
	BaseModel

	Name string `gorm:"type:varchar(255);not null"`

	IsActive bool `gorm:"default:true"`
}
