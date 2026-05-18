package service

import (
	"errors"
	"fmt"
	"os"
	"permit-license/internal/constants"
	"permit-license/internal/dto"
	"permit-license/internal/helper"
	"permit-license/internal/model"
	"permit-license/internal/repository"
	"time"

	"gorm.io/gorm"
)

type ApprovalService struct {
	Repo         repository.PermitLicenseRepository
	UserRepo     repository.UserRepository
	EmailService EmailService
}

func (s *ApprovalService) getStatusByCode(code string) (*model.ApprovalStatus, error) {

	return s.Repo.FindStatusByCode(code)
}

func (s *ApprovalService) createApprovalHistory(tx *gorm.DB, permitID string, userID string, statusID string, notes string) error {

	history := model.ApprovalHistory{
		PermitLicenseID: permitID,
		ApproverID:      userID,
		StatusID:        statusID,
		Notes:           notes,
	}

	return s.Repo.CreateApprovalHistoryTx(
		tx,
		&history,
	)
}

func (s *ApprovalService) sendWaitingDirectorEmail(
	permit *model.PermitLicense,
) {

	directors, _ :=
		s.UserRepo.FindByRole(
			constants.RoleDirector,
		)

	emails :=
		helper.ExtractEmails(
			directors,
		)

	url := fmt.Sprintf(
		"%s/documents/%s",
		os.Getenv("APP_URL"),
		permit.ID.String(),
	)

	body := helper.WaitingApprovalEmail(
		permit.DocumentName,
		url,
	)

	s.EmailService.SendAsync(
		emails,
		"Document Waiting for Approval",
		body,
	)
}

func (s *ApprovalService) sendApprovedEmail(
	permit *model.PermitLicense,
) {

	headUsers, _ :=
		s.UserRepo.FindByRoleAndUnit(
			constants.RoleHeadUnit,
			permit.UnitID,
		)

	headEmails :=
		helper.ExtractEmails(headUsers)

	uploaderEmail :=
		helper.UserEmail(
			&permit.User,
		)

	emails :=
		helper.MergeEmails(
			headEmails,
			uploaderEmail,
		)

	url := fmt.Sprintf(
		"%s/documents/%s",
		os.Getenv("APP_URL"),
		permit.ID.String(),
	)

	body := helper.ApprovedEmail(
		permit.DocumentName,
		url,
	)

	s.EmailService.SendAsync(
		emails,
		"Document Approved",
		body,
	)
}

func (s *ApprovalService) sendRejectedEmail(
	permit *model.PermitLicense,
	reason string,
) {

	headUsers, _ :=
		s.UserRepo.FindByRoleAndUnit(
			constants.RoleHeadUnit,
			permit.UnitID,
		)

	headEmails :=
		helper.ExtractEmails(headUsers)

	uploaderEmail :=
		helper.UserEmail(
			&permit.User,
		)

	emails :=
		helper.MergeEmails(
			headEmails,
			uploaderEmail,
		)

	url := fmt.Sprintf(
		"%s/documents/%s",
		os.Getenv("APP_URL"),
		permit.ID.String(),
	)

	body := helper.RejectedEmail(
		permit.DocumentName,
		url,
		reason,
	)

	s.EmailService.SendAsync(
		emails,
		"Document Rejected",
		body,
	)
}

func (s *ApprovalService) Approve(permitID string, userID string, role string, userUnitID string, req dto.ApprovalRequest) error {

	return repository.WithTransaction(
		func(tx *gorm.DB) error {
			permit, err := s.Repo.FindByIdForUpdate(tx, permitID)
			if err != nil {
				return err
			}

			var nextStatusCode string

			switch role {
			case constants.RoleHeadUnit:
				if permit.CurrentStatus.Code != constants.StatusWaitingApproval {
					return errors.New("document not waiting for head approval")
				}

				if permit.UnitID != userUnitID {
					return errors.New("user not authorized to approve this document")
				}

				nextStatusCode = constants.StatusWaitingDirectorApproval

			case constants.RoleDirector:
				if permit.CurrentStatus.Code != constants.StatusWaitingDirectorApproval {
					return errors.New("document not waiting for director approval")
				}
				nextStatusCode = constants.StatusApproved
			default:
				return errors.New("role not allowed to approve")
			}

			nextStatus, err := s.getStatusByCode(nextStatusCode)
			if err != nil {
				return err
			}

			permit.CurrentStatusID = nextStatus.ID.String()

			if nextStatusCode == constants.StatusApproved {
				now := time.Now()
				permit.ApprovedAt = &now
			}

			err = s.Repo.UpdateTx(tx, permit.ID.String(), permit)
			if err != nil {
				return err
			}

			err = s.createApprovalHistory(tx, permit.ID.String(), userID, nextStatus.ID.String(), req.Notes)
			if err != nil {
				return err
			}

			switch nextStatusCode {
			case constants.StatusWaitingDirectorApproval:
				s.sendWaitingDirectorEmail(permit)
			case constants.StatusApproved:
				s.sendApprovedEmail(permit)
			}

			return nil
		})

}

func (s *ApprovalService) Reject(permitID string, userID string, role string, userUnitID string, req dto.ApprovalRequest) error {

	return repository.WithTransaction(
		func(tx *gorm.DB) error {
			permit, err := s.Repo.FindByIdForUpdate(tx, permitID)
			if err != nil {
				return err
			}

			switch role {

			case constants.RoleHeadUnit:
				if permit.CurrentStatus.Code != constants.StatusWaitingApproval {
					return errors.New("document not waiting for head approval")
				}
				if permit.UnitID != userUnitID {
					return errors.New("user not authorized to approve this document")
				}

			case constants.RoleDirector:
				if permit.CurrentStatus.Code != constants.StatusWaitingDirectorApproval {
					return errors.New("document not waiting for director approval")
				}

			default:
				return errors.New("role not allowed to reject")
			}

			rejectedStatus, err := s.getStatusByCode(constants.StatusRejected)
			if err != nil {
				return err
			}

			permit.CurrentStatusID = rejectedStatus.ID.String()
			permit.RejectedReason = req.Notes

			err = s.Repo.UpdateTx(tx, permit.ID.String(), permit)
			if err != nil {
				return err
			}

			err = s.createApprovalHistory(tx, permit.ID.String(), userID, rejectedStatus.ID.String(), req.Notes)
			if err != nil {
				return err
			}

			s.sendRejectedEmail(permit, req.Notes)

			return nil
		})

}
