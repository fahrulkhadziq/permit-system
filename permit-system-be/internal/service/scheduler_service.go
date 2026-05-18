package service

import (
	"fmt"
	"os"
	"permit-license/internal/helper"
	"permit-license/internal/repository"
	"time"
)

type SchedulerService struct {
	Repo         repository.SchedulerRepository
	EmailService EmailService
	UserRepo     repository.UserRepository
}

func (s *SchedulerService) SendExpiredReminder() {

	targetDate := time.Now().AddDate(0, 0, 30) // 30 days from now

	permits, err := s.Repo.FindExpiringDocuments(targetDate)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, permit := range permits {

		headUsers, _ := s.UserRepo.FindByRoleAndUnit("HEAD_UNIT", permit.UnitID)
		headEmails := helper.ExtractEmails(headUsers)
		userEmail := helper.UserEmail(&permit.User)

		emails := helper.MergeEmails(
			headEmails,
			userEmail,
		)

		url := fmt.Sprintf(
			"%s/documents/%s",
			os.Getenv("APP_URL"),
			permit.ID.String(),
		)

		body := helper.ExpirationReminderEmail(
			permit.DocumentName,
			url,
			permit.ExpiredAt.Format(
				"2006-01-02",
			),
		)

		s.EmailService.SendAsync(
			emails,
			"Document Expired Reminder",
			body,
		)
	}
}

func (s *SchedulerService) AutoMarkExpiredDocuments() {

	permits, err := s.Repo.FindExpiredDocuments()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var ids []string
	for _, permit := range permits {
		ids = append(ids, permit.ID.String())
	}

	if len(ids) == 0 {
		return
	}

	err = s.Repo.MarkExpired(ids)
	if err != nil {
		fmt.Println(err.Error())
	}
}
