package helper

import (
	"permit-license/internal/dto"
	"permit-license/internal/model"
)

func ToUnitResponse(unit model.Unit) dto.UnitResponse {

	return dto.UnitResponse{
		ID: unit.ID.String(),

		Name: unit.Name,

		IsActive: unit.IsActive,
	}
}
