package dto

import "time"

type PermitLicenseResponse struct {
	ID                    string  `json:"id"`
	DocumentName          string  `json:"document_name"`
	Description           string  `json:"description"`
	FileURL               string  `json:"file_url"`
	FileSize              int64   `json:"file_size"`
	ExpiredAt             string  `json:"expired_at"`
	Status                string  `json:"status"`
	MasterDocument        string  `json:"master_document"`
	UploadedBy            string  `json:"uploaded_by"`
	Unit                  string  `json:"unit"`
	RelatedPrevDocumentID *string `json:"related_prev_document_id,omitempty"`
}

type PermitLicenseDetailResponse struct {
	ID string `json:"id"`

	DocumentName string `json:"document_name"`
	Description  string `json:"description"`

	FileURL  string `json:"file_url"`
	FileSize int64  `json:"file_size"`

	ExpiredAt  time.Time  `json:"expired_at"`
	ApprovedAt *time.Time `json:"approved_at"`

	RejectedReason string `json:"rejected_reason"`

	IsActive bool `json:"is_active"`

	IsExtend *bool `json:"is_extend"`

	MasterDocument MasterDocumentResponse `json:"master_document"`

	User UserResponse `json:"user"`

	Unit UnitResponse `json:"unit"`

	CurrentStatus ApprovalStatusResponse `json:"current_status"`

	RelatedPrevDocument *RelatedPermitResponse `json:"related_prev_document"`

	RelatedNextDocument *RelatedPermitResponse `json:"related_next_document"`

	ApprovalHistories []ApprovalHistoryResponse `json:"approval_histories"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserResponse struct {
	ID string `json:"id"`

	FullName string `json:"full_name"`

	Email string `json:"email"`

	Unit UnitResponse `json:"unit"`
}

type ApprovalStatusResponse struct {
	ID string `json:"id"`

	Code string `json:"code"`

	Name string `json:"name"`
}

type RelatedPermitResponse struct {
	ID string `json:"id"`

	DocumentName string `json:"document_name"`

	ExpiredAt time.Time `json:"expired_at"`
}

type ApprovalHistoryResponse struct {
	ID string `json:"id"`

	Notes string `json:"notes"`

	CreatedAt time.Time `json:"created_at"`

	Approver UserResponse `json:"approver"`

	Status ApprovalStatusResponse `json:"status"`
}
