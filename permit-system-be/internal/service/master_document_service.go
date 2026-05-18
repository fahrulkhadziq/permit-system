package service

import (
	"permit-license/internal/dto"
	"permit-license/internal/model"
	"permit-license/internal/repository"
)

type MasterDocumentService struct {
	Repo repository.MasterDocumentRepository
}

func (s *MasterDocumentService) Create(req dto.MasterDocumentRequest) error {
	masterDoc := &model.MasterDocument{
		Name:        req.Name,
		Code:        req.Code,
		Description: req.Description,
	}
	return s.Repo.CreateMasterDocument(masterDoc)
}

func (s *MasterDocumentService) FindAll(params dto.QueryParams) ([]dto.MasterDocumentResponse, int64, error) {

	masterDocs, total, err := s.Repo.FindAll(params)
	if err != nil {
		return nil, 0, err
	}

	responses := make([]dto.MasterDocumentResponse, len(masterDocs))
	for i, doc := range masterDocs {
		responses[i] = dto.MasterDocumentResponse{
			ID:          doc.ID.String(),
			Name:        doc.Name,
			Code:        doc.Code,
			Description: doc.Description,
		}
	}

	return responses, total, nil
}

func (s *MasterDocumentService) FindByID(id string) (*dto.MasterDocumentResponse, error) {
	masterDoc, err := s.Repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	response := &dto.MasterDocumentResponse{
		ID:          masterDoc.ID.String(),
		Name:        masterDoc.Name,
		Code:        masterDoc.Code,
		Description: masterDoc.Description,
	}
	return response, nil
}

func (s *MasterDocumentService) Update(
	id string,
	req dto.UpdateMasterDocumentRequest,
) error {

	data, err := s.Repo.FindByID(id)

	if err != nil {
		return err
	}

	data.Code = req.Code
	data.Name = req.Name
	data.Description = req.Description
	data.IsActive = req.IsActive

	return s.Repo.Update(id, data)
}

func (s *MasterDocumentService) Delete(
	id string,
) error {

	return s.Repo.Delete(id)
}
