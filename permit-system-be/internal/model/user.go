package model

type User struct {
	BaseModel

	UnitID *string

	FullName string `gorm:"type:varchar(255)"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`

	IsActive bool `gorm:"default:true"`

	Unit Unit `gorm:"foreignKey:unit_id"`

	Roles []Role `gorm:"many2many:user_roles"`
}
