package dto

type CreateUnitRequest struct {
	Name string `json:"name" validate:"required"`
}

type UpdateUnitRequest struct {
	Name     string `json:"name" validate:"required"`
	IsActive bool   `json:"is_active"`
}

type UnitResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}
