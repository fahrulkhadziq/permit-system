package helper

import (
	"permit-license/internal/dto"
	"permit-license/internal/model"
)

func ToPermitLicenseResponse(permit model.PermitLicense) dto.PermitLicenseResponse {
	return dto.PermitLicenseResponse{
		ID: permit.ID.String(),

		DocumentName:          permit.DocumentName,
		Description:           permit.Description,
		FileURL:               permit.FileURL,
		FileSize:              permit.FileSize,
		ExpiredAt:             permit.ExpiredAt.Format("2006-01-02"),
		Status:                permit.CurrentStatus.Name,
		MasterDocument:        permit.MasterDocument.Name,
		UploadedBy:            permit.User.FullName,
		Unit:                  permit.Unit.Name,
		RelatedPrevDocumentID: permit.RelatedPrevDocumentID,
	}
}

func ToPermitLicenseDetailsResponse(permit *model.PermitLicense) dto.PermitLicenseDetailResponse {
	response := dto.PermitLicenseDetailResponse{
		ID: permit.ID.String(),

		DocumentName:   permit.DocumentName,
		Description:    permit.Description,
		FileURL:        permit.FileURL,
		FileSize:       permit.FileSize,
		ExpiredAt:      permit.ExpiredAt,
		ApprovedAt:     permit.ApprovedAt,
		RejectedReason: permit.RejectedReason,
		IsActive:       permit.IsActive,
		IsExtend:       *permit.IsExtend,
		MasterDocument: dto.MasterDocumentResponse{
			ID:          permit.MasterDocument.ID.String(),
			Code:        permit.MasterDocument.Code,
			Name:        permit.MasterDocument.Name,
			Description: permit.MasterDocument.Description,
		},
		User: dto.UserResponse{
			ID:       permit.User.ID.String(),
			FullName: permit.User.FullName,
			Email:    permit.User.Email,
			Unit: dto.UnitResponse{
				ID:   permit.User.Unit.ID.String(),
				Name: permit.User.Unit.Name,
			},
		},
		Unit: dto.UnitResponse{
			ID:   permit.Unit.ID.String(),
			Name: permit.Unit.Name,
		},

		CurrentStatus: dto.ApprovalStatusResponse{
			ID:   permit.CurrentStatus.ID.String(),
			Code: permit.CurrentStatus.Code,
			Name: permit.CurrentStatus.Name,
		},
		CreatedAt: permit.CreatedAt,
		UpdatedAt: permit.UpdatedAt,
	}

	if permit.RelatedPrevDocument != nil {
		response.RelatedPrevDocument =
			&dto.RelatedPermitResponse{
				ID:           permit.RelatedPrevDocument.ID.String(),
				DocumentName: permit.RelatedPrevDocument.DocumentName,
				ExpiredAt:    permit.RelatedPrevDocument.ExpiredAt,
			}
	}

	if permit.RelatedNextDocument != nil {
		response.RelatedNextDocument =
			&dto.RelatedPermitResponse{
				ID:           permit.RelatedNextDocument.ID.String(),
				DocumentName: permit.RelatedNextDocument.DocumentName,
				ExpiredAt:    permit.RelatedNextDocument.ExpiredAt,
			}
	}

	for _, history := range permit.ApprovalHistories {

		response.ApprovalHistories = append(response.ApprovalHistories, dto.ApprovalHistoryResponse{
			ID:        history.ID.String(),
			Notes:     history.Notes,
			CreatedAt: history.CreatedAt,

			Approver: dto.UserResponse{
				ID:       history.Approver.ID.String(),
				FullName: history.Approver.FullName,
				Email:    history.Approver.Email,
				Unit: dto.UnitResponse{
					ID:   history.Approver.Unit.ID.String(),
					Name: history.Approver.Unit.Name,
				},
			},
			Status: dto.ApprovalStatusResponse{
				ID:   history.Status.ID.String(),
				Code: history.Status.Code,
				Name: history.Status.Name,
			},
		})
	}
	return response
}
