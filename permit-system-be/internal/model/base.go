package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (b *BaseModel) BeforeCreate(tx *gorm.DB) error {
	b.ID = uuid.New()
	return nil
}
