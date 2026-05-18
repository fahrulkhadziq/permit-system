package repository

import (
	"permit-license/config"
	"permit-license/internal/model"
	"time"
)

type SchedulerRepository struct{}

func (r *SchedulerRepository) FindExpiringDocuments(targetDate time.Time) ([]model.PermitLicense, error) {
	var permits []model.PermitLicense

	err := config.DB.
		Preload("User").
		Preload("Unit").
		Where("DATE(expired_at) = DATE(?)", targetDate).
		Where("is_active = ?", true).
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

func (r *SchedulerRepository) MarkExpired(ids []string) error {
	return config.DB.
		Model(&model.PermitLicense{}).
		Where("id IN ?", ids).
		Update("is_active", false).Error
}
