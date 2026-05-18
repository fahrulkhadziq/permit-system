package repository

import (
	"permit-license/config"

	"gorm.io/gorm"
)

func WithTransaction(fn func(tx *gorm.DB) error) error {
	tx := config.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err := fn(tx)

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
