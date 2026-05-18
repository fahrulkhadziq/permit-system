package dto

type QueryParams struct {
	Page  int `query:"page"`
	Limit int `query:"limit"`

	Search string `query:"search"`

	Sort  string `query:"sort"`
	Order string `query:"order"`

	StatusCode string `query:"status"`
	UnitID     string `query:"unit_id"`
	UploadedBy string `query:"uploaded_by"`
}
