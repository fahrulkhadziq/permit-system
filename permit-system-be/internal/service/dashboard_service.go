package service

import (
	"permit-license/internal/dto"
	"permit-license/internal/repository"
)

type DashboardService struct {
	Repo repository.DashboardRepository
}

func (s *DashboardService) GetStatistics(unitID *string) (dto.DashboardResponse, error) {
	return s.Repo.GetStatistics(unitID)
}
