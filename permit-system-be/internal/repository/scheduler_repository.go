package repository

import (
	"permit-license/config"
	"permit-license/internal/model"
	"time"
)

type SchedulerRepository struct{}

func (r *SchedulerRepository) FindExpiringDocuments() ([]model.PermitLicense, error) {

	var permits []model.PermitLicense

	err := config.DB.
		Preload("User").
		Where(`
			current_status_id = ?
			AND is_extend IS NULL
			AND is_active = ?
		`,
			"202e64d1-656b-4173-b6f3-095536888a17",
			true,
		).
		Find(&permits).Error

	return permits, err
}

func (r *SchedulerRepository) FindExpiredDocuments() ([]model.PermitLicense, error) {
	var permits []model.PermitLicense

	now := time.Now()

	err := config.DB.
		Where("expired_at < ?", now).
		Where("is_active = ?", true).
		Find(&permits).Error

	return permits, err
}
