package database

import (
	"permit-license/config"
	"permit-license/internal/model"
)

func Seed() {
	roles := []model.Role{
		{Name: "USER_UNIT"},
		{Name: "HEAD_UNIT"},
		{Name: "DIRECTOR"},
	}

	for _, role := range roles {
		config.DB.FirstOrCreate(&role, model.Role{Name: role.Name})
	}

	statuses := []model.ApprovalStatus{
		{
			Code: "WAITING_HEAD_APPROVAL",
			Name: "Waiting Head Approval",
		},
		{
			Code: "WAITING_DIRECTOR_APPROVAL",
			Name: "Waiting Director Approval",
		},
		{
			Code: "APPROVED",
			Name: "Approved",
		},
		{
			Code: "REJECTED",
			Name: "Rejected",
		},
		{
			Code: "EXPIRED",
			Name: "Expired",
		},
		{
			Code: "NOT_EXTENDED",
			Name: "Not Extended",
		},
	}
	for _, status := range statuses {
		config.DB.FirstOrCreate(&status, model.ApprovalStatus{Code: status.Code})
	}
}
