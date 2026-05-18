package dto

type MasterDocumentRequest struct {
	Name        string `json:"name" validate:"required"`
	Code        string `json:"code" validate:"required"`
	Description string `json:"description"`
}

type UpdateMasterDocumentRequest struct {
	Code        string `json:"code"validate:"required"`
	Name        string `json:"name"validate:"required"`
	Description string `json:"description"`
	IsActive    bool   `json:"is_active"`
}

type MasterDocumentResponse struct {
	ID          string `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
