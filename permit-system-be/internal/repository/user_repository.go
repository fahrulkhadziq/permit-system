package repository

import (
	"permit-license/config"
	"permit-license/internal/model"
)

type UserRepository struct{}

func (r *UserRepository) CreateUserRole(userRole *model.UserRole) error {
	return config.DB.Create(userRole).Error
}

func (r *UserRepository) FindUserByEmail(email string) (*model.User, error) {
	var user model.User

	err := config.DB.Where("email = ?", email).First(&user).Error

	return &user, err
}

func (r *UserRepository) GetUserRole(userID string) (*model.UserRole, error) {
	var userRole model.UserRole

	err := config.DB.Preload("Role").
		Where("user_id", userID).
		First(&userRole).Error

	return &userRole, err
}

func (r *UserRepository) FindByRole(roleCode string) ([]model.User, error) {
	var users []model.User

	err := config.DB.Joins(
		"JOIN user_roles ur  ON ur.user_id = users.id",
	).Joins(
		"JOIN roles r ON r.id = ur.role_id",
	).
		Where("r.code = ?", roleCode).
		Find(&users).Error

	return users, err
}

func (r *UserRepository) FindByRoleAndUnit(roleCode, unitID string) ([]model.User, error) {
	var users []model.User

	err := config.DB.Joins(
		"JOIN user_roles ur  ON ur.user_id = users.id",
	).Joins(
		"JOIN roles r ON r.id = ur.role_id",
	).
		Where("r.code = ? AND users.unit_id = ?", roleCode, unitID).
		Find(&users).Error

	return users, err
}

func (r *UserRepository) CountByUnitID(
	unitID string,
) (int64, error) {

	var total int64

	err := config.DB.
		Model(&model.User{}).
		Where("unit_id = ?", unitID).
		Count(&total).
		Error

	return total, err
}
