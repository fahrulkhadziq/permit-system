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

func (r *DashboardRepository) baseQuery() *gorm.DB {

	query := config.DB.Model(
		&model.PermitLicense{}).
		Where("is_active = ?", true)

	return query
}

func (r *DashboardRepository) GetStatistics(unitID *string) (dto.DashboardResponse, error) {
	baseQuery := r.baseQuery()
	if unitID != nil {
		baseQuery = baseQuery.Where(
			"unit_id = ?",
			*unitID,
		)
	}
	return r.buildStatistics(baseQuery)

}

func (r *DashboardRepository) GetStatisticsAll() (dto.DashboardResponse, error) {
	return r.buildStatistics(
		r.baseQuery(),
	)
}

func (r *DashboardRepository) buildStatistics(baseQuery *gorm.DB) (dto.DashboardResponse, error) {

	var response dto.DashboardResponse

	now := time.Now()

	approvedSubQuery := config.DB.
		Model(&model.ApprovalStatus{}).
		Select("id").
		Where(
			"code = ?",
			constants.StatusApproved,
		)

	pendingSubQuery := config.DB.
		Model(&model.ApprovalStatus{}).
		Select("id").
		Where(
			"code IN ?",
			[]string{
				constants.StatusWaitingApproval,
				constants.StatusWaitingDirectorApproval,
			},
		)

	rejectedSubQuery := config.DB.
		Model(&model.ApprovalStatus{}).
		Select("id").
		Where(
			"code = ?",
			constants.StatusRejected,
		)

	// TOTAL DOCUMENTS

	baseQuery.
		Session(&gorm.Session{}).
		Count(
			&response.TotalDocuments,
		)

	// ACTIVE DOCUMENTS

	baseQuery.
		Session(&gorm.Session{}).
		Where(
			"current_status_id IN (?)",
			approvedSubQuery,
		).
		Where(
			"expired_at >= ?",
			now,
		).
		Count(
			&response.ActiveDocuments,
		)

	// EXPIRED DOCUMENTS

	baseQuery.
		Session(&gorm.Session{}).
		Where(
			"expired_at < ?",
			now,
		).
		Count(
			&response.ExpiredDocuments,
		)

	// PENDING APPROVALS

	baseQuery.
		Session(&gorm.Session{}).
		Where(
			"current_status_id IN (?)",
			pendingSubQuery,
		).
		Count(
			&response.PendingApprovals,
		)

	// APPROVED DOCUMENTS

	baseQuery.
		Session(&gorm.Session{}).
		Where(
			"current_status_id IN (?)",
			approvedSubQuery,
		).
		Count(
			&response.ApprovedDocuments,
		)

	// REJECTED DOCUMENTS

	baseQuery.
		Session(&gorm.Session{}).
		Where(
			"current_status_id IN (?)",
			rejectedSubQuery,
		).
		Count(
			&response.RejectedDocuments,
		)

	// NOT EXTENDED

	baseQuery.
		Session(&gorm.Session{}).
		Where(
			"expired_at < ?",
			now,
		).
		Where(
			"is_extend IS NULL",
		).
		Count(
			&response.NotExtendedDocuments,
		)

	return response, nil
}
