package service

import (
	"errors"
	"permit-license/internal/dto"
	"permit-license/internal/model"
	"permit-license/internal/repository"
)

type UnitService struct {
	Repo     repository.UnitRepository
	UserRepo repository.UserRepository
}

func (s *UnitService) Create(req dto.CreateUnitRequest) error {

	unit := model.Unit{
		Name:     req.Name,
		IsActive: true,
	}

	return s.Repo.Create(&unit)
}

func (s *UnitService) FindAll(params dto.QueryParams) ([]model.Unit, int64, error) {
	return s.Repo.FindAll(params)
}

func (s *UnitService) FindByID(id string) (*model.Unit, error) {
	return s.Repo.FindByID(id)
}

func (s *UnitService) Update(id string, req dto.UpdateUnitRequest) error {

	unit, err := s.Repo.FindByID(id)

	if err != nil {
		return err
	}

	unit.Name = req.Name
	unit.IsActive = req.IsActive

	return s.Repo.Update(id, unit)
}

func (s *UnitService) Delete(id string) error {

	totalUsers, err :=
		s.UserRepo.CountByUnitID(id)

	if err != nil {
		return err
	}

	if totalUsers > 0 {

		return errors.New(
			"unit still used by users",
		)
	}

	return s.Repo.Delete(id)
}
