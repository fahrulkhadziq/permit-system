package model

type ApprovalStatus struct {
	BaseModel

	Code string `gorm:"unique;not null"`
	Name string `gorm:"not null"`
}
