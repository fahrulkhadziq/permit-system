package helper

import (
	"math"
	"permit-license/internal/dto"
)

func Paginate(page int, limit int, totalRows int64, data interface{}) dto.PaginationResponse {

	totalPages := int(math.Ceil(float64(totalRows) / float64(limit)))

	return dto.PaginationResponse{
		Page:       page,
		Limit:      limit,
		TotalRows:  totalRows,
		TotalPages: totalPages,
		Data:       data,
	}
}

func NormalizePagination(page int, limit int) (int, int) {
	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 10
	}

	if limit > 100 {
		limit = 100
	}

	return page, limit
}
