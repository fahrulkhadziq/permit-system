package repository

import (
	"permit-license/config"
	"permit-license/internal/constants"
	"permit-license/internal/dto"
	"permit-license/internal/model"
	"time"

	"gorm.io/gorm"
)

type DashboardRepository struct{}

func (r *DashboardRepository) baseQuery(unitID *string) *gorm.DB {

	query := config.DB.Model(
		&model.PermitLicense{},
	)

	if unitID != nil {
		query = query.Where("unit_id = ?", *unitID)
	}

	return query
}

func (r *DashboardRepository) GetStatistics(unitID *string) (dto.DashboardResponse, error) {
	var response dto.DashboardResponse

	now := time.Now()

	baseQuery := r.baseQuery(unitID)
	baseQuery.Count(&response.TotalDocuments)

	baseQuery.
		Where(
			"current_status_id IN (?)",
			config.DB.
				Model(&model.ApprovalStatus{}).
				Select("id").
				Where("code = ?", constants.StatusApproved),
		).
		Where("expired_at >= ?", now).
		Count(&response.ActiveDocuments)

	baseQuery.
		Where("expired_at < ?", now).
		Count(&response.ExpiredDocuments)

	baseQuery.
		Where(
			"current_status_id IN (?)",
			config.DB.
				Model(&model.ApprovalStatus{}).
				Select("id").
				Where("code IN ?",
					[]string{
						constants.StatusWaitingApproval,
						constants.StatusWaitingDirectorApproval,
					}),
		).
		Count(&response.PendingApprovals)

	baseQuery.
		Where(
			"current_status_id IN (?)",
			config.DB.
				Model(&model.ApprovalStatus{}).
				Select("id").
				Where("code = ?", constants.StatusApproved),
		).
		Count(&response.ApprovedDocuments)

	baseQuery.
		Where(
			"current_status_id IN (?)",
			config.DB.
				Model(&model.ApprovalStatus{}).
				Select("id").
				Where("code = ?", constants.StatusRejected),
		).
		Count(&response.RejectedDocuments)

	// NOT EXTENDED

	subQuery := config.DB.
		Model(&model.PermitLicense{}).
		Select("related_prev_document_id").
		Where("related_prev_document_id IS NOT NULL")

	baseQuery.
		Where("expired_at < ?", now).
		Where("id NOT IN (?)", subQuery).
		Count(&response.NotExtendedDocuments)

	return response, nil

}
