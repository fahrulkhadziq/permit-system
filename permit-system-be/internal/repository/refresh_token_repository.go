package repository

import (
	"permit-license/config"
	"permit-license/internal/model"
)

type RefreshTokenRepository struct{}

func (r *RefreshTokenRepository) Create(
	data *model.RefreshToken,
) error {

	return config.DB.Create(data).Error
}

func (r *RefreshTokenRepository) FindByToken(
	token string,
) (*model.RefreshToken, error) {

	var data model.RefreshToken

	err := config.DB.
		Preload("User").
		Where("token = ?", token).
		First(&data).
		Error

	return &data, err
}

func (r *RefreshTokenRepository) Revoke(
	token string,
) error {

	return config.DB.
		Model(&model.RefreshToken{}).
		Where("token = ?", token).
		Update("is_revoked", true).
		Error
}
