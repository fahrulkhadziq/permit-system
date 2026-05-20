package dto

type CreatePermitLicenseRequest struct {
	MaterDocumentID       string  `form:"master_document_id" validate:"required"`
	DocumentName          string  `form:"document_name" validate:"required"`
	Description           string  `form:"description" validate:"required"`
	ExpiredAt             string  `form:"expired_at" validate:"required"`
	RelatedPrevDocumentID *string `form:"related_prev_document_id"`
}

type UpdatePermitLicenseRequest struct {
	MasterDocumentID *string `form:"master_document_id"`
	DocumentName     *string `form:"document_name"`
	Description      *string `form:"description"`
	ExpiredAt        *string `form:"expired_at"`

	IsExtend *bool `form:"is_extend"`
}
