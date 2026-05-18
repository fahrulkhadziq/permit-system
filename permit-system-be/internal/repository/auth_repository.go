package repository

import (
	"permit-license/config"
	"permit-license/internal/model"
)

type AuthRepository struct{}

func (r *AuthRepository) CreateUser(user *model.User) error {
	return config.DB.Create(user).Error
}
