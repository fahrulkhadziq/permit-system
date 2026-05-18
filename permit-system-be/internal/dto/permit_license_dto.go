package dto

type CreatePermitLicenseRequest struct {
	MaterDocumentID       string  `form:"master_document_id" validate:"required"`
	DocumentName          string  `form:"document_name" validate:"required"`
	Description           string  `form:"description" validate:"required"`
	ExpiredAt             string  `form:"expired_at" validate:"required"`
	RelatedPrevDocumentID *string `form:"related_prev_document_id"`
}
