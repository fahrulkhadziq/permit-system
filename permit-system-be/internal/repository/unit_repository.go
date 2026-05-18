package repository

import (
	"permit-license/config"
	"permit-license/internal/dto"
	"permit-license/internal/helper"
	"permit-license/internal/model"
)

type UnitRepository struct{}

func (r *UnitRepository) Create(unit *model.Unit) error {
	return config.DB.Create(unit).Error
}

func (r *UnitRepository) FindAll(
	params dto.QueryParams,
) ([]model.Unit, int64, error) {

	var units []model.Unit
	var totalRows int64

	page, limit :=
		helper.NormalizePagination(
			params.Page,
			params.Limit,
		)

	query := config.DB.
		Model(&model.Unit{}).
		Where("is_active = ?", true)

	query = helper.ApplySearch(
		query,
		params.Search,
		[]string{
			"name",
		},
	)

	query.Count(&totalRows)

	query = helper.ApplySorting(
		query,
		params.Sort,
		params.Order,
	)

	query = helper.ApplyPagination(
		query,
		page,
		limit,
	)

	err := query.Find(&units).Error

	return units, totalRows, err
}

func (r *UnitRepository) FindByID(id string) (*model.Unit, error) {

	var unit model.Unit

	err := config.DB.
		First(&unit, "id = ?", id).
		Error

	return &unit, err
}

func (r *UnitRepository) Update(id string, unit *model.Unit) error {

	return config.DB.
		Model(&model.Unit{}).
		Where("id = ?", id).
		Updates(unit).
		Error
}

func (r *UnitRepository) Delete(id string) error {
	return config.DB.
		Model(&model.Unit{}).
		Where("id = ?", id).
		Update("is_active", false).
		Error
}
