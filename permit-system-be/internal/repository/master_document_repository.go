package repository

import (
	"permit-license/config"
	"permit-license/internal/dto"
	"permit-license/internal/helper"
	"permit-license/internal/model"
)

type MasterDocumentRepository struct{}

func (r *MasterDocumentRepository) CreateMasterDocument(masterDoc *model.MasterDocument) error {

	return config.DB.Create(masterDoc).Error
}

func (r *MasterDocumentRepository) FindAll(params dto.QueryParams) ([]model.MasterDocument, int64, error) {

	var masterDocs []model.MasterDocument
	var totalRows int64

	page, limit := helper.NormalizePagination(
		params.Page,
		params.Limit,
	)

	query := config.DB.Model(&model.MasterDocument{})

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

	err := query.Find(&masterDocs).Error
	return masterDocs, totalRows, err
}

func (r *MasterDocumentRepository) FindByID(id string) (*model.MasterDocument, error) {
	var masterDoc model.MasterDocument

	err := config.DB.Where("id = ?", id).First(&masterDoc).Error
	if err != nil {
		return nil, err
	}
	return &masterDoc, err
}

func (r *MasterDocumentRepository) Update(id string, data *model.MasterDocument) error {
	return config.DB.
		Model(&model.MasterDocument{}).
		Where("id = ?", id).
		Updates(data).Error
}

func (r *MasterDocumentRepository) Delete(
	id string,
) error {

	return config.DB.
		Model(&model.MasterDocument{}).
		Where("id = ?", id).
		Update("is_active", false).
		Error
}
