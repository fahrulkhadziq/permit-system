package dto

type PaginationRequest struct {
	Page   int    `query:"page"`
	Limit  int    `query:"limit"`
	Search string `query:"search"`
}

type PaginationResponse struct {
	Page       int         `json:"page"`
	Limit      int         `json:"limit"`
	TotalRows  int64       `json:"total_rows"`
	TotalPages int         `json:"total_pages"`
	Data       interface{} `json:"data"`
}
