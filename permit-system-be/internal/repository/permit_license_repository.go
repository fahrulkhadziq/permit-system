package repository

import (
	"permit-license/config"
	"permit-license/internal/dto"
	"permit-license/internal/helper"
	"permit-license/internal/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PermitLicenseRepository struct{}

func (r *PermitLicenseRepository) CreatePermitLicense(permit *model.PermitLicense) error {

	return config.DB.Create(permit).Error
}

func (r *PermitLicenseRepository) CreateApprovalHistory(history *model.ApprovalHistory) error {

	return config.DB.Create(history).Error
}

func (r *PermitLicenseRepository) FindStatusByCode(code string) (*model.ApprovalStatus, error) {

	var status model.ApprovalStatus

	err := config.DB.Where("code = ?", code).First(&status).Error

	return &status, err
}

func (r *PermitLicenseRepository) FindAll(params dto.QueryParams) ([]model.PermitLicense, int64, error) {

	var permits []model.PermitLicense
	var totalRows int64

	page, limit := helper.NormalizePagination(
		params.Page,
		params.Limit,
	)

	query := config.DB.
		Model(&model.PermitLicense{}).
		Preload("MasterDocument").
		Preload("CurrentStatus").
		Preload("User").
		Preload("Unit")

	query = helper.ApplySearch(
		query,
		params.Search,
		[]string{
			"document_name",
			"description",
		},
	)

	// UNIT FILTER
	if params.UnitID != "" {

		query = query.Where(
			"unit_id = ?",
			params.UnitID,
		)
	}

	// UPLOADER FILTER
	if params.UploadedBy != "" {

		query = query.Where(
			"uploaded_by = ?",
			params.UploadedBy,
		)
	}

	// STATUS FILTER
	if params.StatusCode != "" {

		query = query.Joins(
			"JOIN approval_statuses ON approval_statuses.id = permit_licenses.current_status_id",
		).Where(
			"approval_statuses.code = ?",
			params.StatusCode,
		)
	}

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

	err := query.Find(&permits).Error

	return permits, totalRows, err
}

func (r *PermitLicenseRepository) FindByID(id string) (*model.PermitLicense, error) {

	var permit model.PermitLicense

	err := config.DB.
		Preload("MasterDocument").
		Preload("CurrentStatus").
		Preload("User").
		Preload("User.Unit").
		Preload("Unit").
		Preload("RelatedPrevDocument").
		Preload("ApprovalHistories").
		Preload("ApprovalHistories.Approver").
		Preload("ApprovalHistories.Approver.Unit").
		Preload("ApprovalHistories.Status").
		First(&permit, "id = ?", id).Error

	return &permit, err
}

// func (r *PermitLicenseRepository) Update(id string, permit *model.PermitLicense) error {

// 	return config.DB.Model(&model.PermitLicense{}).Where("id = ?", id).Updates(permit).Error
// }

func (r *PermitLicenseRepository) FindByIdFull(id string) (*model.PermitLicense, error) {

	var permit model.PermitLicense

	err := config.DB.
		Preload("CurrentStatus").
		Preload("Unit").
		First(&permit, "id = ?", id).Error

	return &permit, err
}

func (r *PermitLicenseRepository) FindByIdForUpdate(tx *gorm.DB, id string) (*model.PermitLicense, error) {

	var permit model.PermitLicense

	err := tx.Clauses(clause.Locking{
		Strength: "UPDATE",
	}).
		Preload("CurrentStatus").
		Preload("Unit").
		Preload("User").
		Preload("MasterDocument").
		First(&permit, "id = ?", id).Error

	return &permit, err
}

func (r *PermitLicenseRepository) UpdateTx(tx *gorm.DB, id string, permit *model.PermitLicense) error {
	return tx.
		Model(&model.PermitLicense{}).
		Where("id = ?", id).
		Updates(permit).Error
}

func (r *PermitLicenseRepository) CreateApprovalHistoryTx(tx *gorm.DB, history *model.ApprovalHistory) error {
	return tx.Create(history).Error
}
