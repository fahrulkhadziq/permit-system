package service

import (
	"errors"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"permit-license/internal/constants"
	"permit-license/internal/dto"
	"permit-license/internal/helper"
	"permit-license/internal/model"
	"permit-license/internal/repository"
	"time"

	"github.com/google/uuid"
)

type PermitLicenseService struct {
	Repo         repository.PermitLicenseRepository
	UserRepo     repository.UserRepository
	EmailService EmailService
}

func validatePDF(fileHeader *multipart.FileHeader) error {
	if fileHeader.Size > 25*1024*1024 {
		return errors.New("Max File Size  is 25 MB")
	}

	ext := filepath.Ext(fileHeader.Filename)

	mustFile := []string{".pdf", ".PDF"}

	if !contains(mustFile, ext) {
		return errors.New("File must be PDF")
	}

	return nil
}

func contains(arrayString []string, ext string) bool {
	for _, v := range arrayString {
		if v == ext {
			return true
		}
	}
	return false
}

func SaveFile(fileHeader *multipart.FileHeader) (string, error) {
	src, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	os.MkdirAll("storage/documents", os.ModePerm)
	filename := fmt.Sprintf("%s%s",
		uuid.NewString(),
		filepath.Ext(fileHeader.Filename),
	)

	path := filepath.Join(
		"storage/documents",
		filename,
	)

	dst, err := os.Create(path)
	if err != nil {
		return "", err
	}

	defer dst.Close()

	_, err = dst.ReadFrom(src)
	if err != nil {
		return "", err
	}

	return "/storage/documents/" + filename, err
}

func (s *PermitLicenseService) Create(req dto.CreatePermitLicenseRequest, fileHeader *multipart.FileHeader, userID string, unitID string) error {

	err := validatePDF(fileHeader)
	if err != nil {
		return err
	}

	filePath, err := SaveFile(fileHeader)
	if err != nil {
		return err
	}

	status, err := s.Repo.FindStatusByCode(
		constants.StatusWaitingApproval,
	)
	if err != nil {
		return err
	}

	expiredAt, err := time.Parse(
		"2006-01-02",
		req.ExpiredAt,
	)
	if err != nil {
		return err
	}

	permit := model.PermitLicense{
		MasterDocumentID:      req.MaterDocumentID,
		UploadedBy:            userID,
		UnitID:                unitID,
		CurrentStatusID:       status.ID.String(),
		DocumentName:          req.DocumentName,
		Description:           req.Description,
		FileURL:               filePath,
		FileSize:              fileHeader.Size,
		ExpiredAt:             expiredAt,
		RelatedPrevDocumentID: req.RelatedPrevDocumentID,
	}

	err = s.Repo.CreatePermitLicense(&permit)
	if err != nil {
		return err
	}

	// email
	headUser, _ := s.UserRepo.FindByRole("HEAD_UNIT")
	emails := helper.ExtractEmails(headUser)

	url := fmt.Sprintf(
		"%s/documents/%s",
		os.Getenv("APP_URL"),
		permit.ID.String(),
	)

	body := helper.WaitingApprovalEmail(
		permit.DocumentName,
		url,
	)

	s.EmailService.SendAsync(emails, "New Document Waiting Approval", body)

	// end email

	history := model.ApprovalHistory{
		PermitLicenseID: permit.ID.String(),
		ApproverID:      userID,
		StatusID:        status.ID.String(),
		Notes:           "Document uploaded",
	}

	return s.Repo.CreateApprovalHistory(&history)

}

func (s *PermitLicenseService) FindAll(params dto.QueryParams) ([]model.PermitLicense, int64, error) {
	return s.Repo.FindAll(params)
}

func (s *PermitLicenseService) FindByID(id string) (*model.PermitLicense, error) {
	return s.Repo.FindByID(id)
}
