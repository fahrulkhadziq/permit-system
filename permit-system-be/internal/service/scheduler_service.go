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

	permits, err := s.Repo.FindExpiringDocuments()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	today := time.Now()

	for _, permit := range permits {

		daysLeft := int(
			permit.ExpiredAt.Sub(today).Hours() / 24,
		)

		shouldSend := false

		switch daysLeft {

		case 30, 21, 14, 7, 0:
			shouldSend = true
		}

		// EXPIRED => EVERY 7 DAYS
		if daysLeft < 0 {

			expiredDays := daysLeft * -1

			if expiredDays%7 == 0 {

				shouldSend = true
			}
		}

		if !shouldSend {
			continue
		}

		headUsers, _ :=
			s.UserRepo.FindByRoleAndUnit(
				"HEAD_UNIT",
				permit.UnitID,
			)

		headEmails :=
			helper.ExtractEmails(
				headUsers,
			)

		userEmail :=
			helper.UserEmail(
				&permit.User,
			)

		emails :=
			helper.MergeEmails(
				headEmails,
				userEmail,
			)

		url := fmt.Sprintf(
			"%s/documents/%s",
			os.Getenv("APP_URL"),
			permit.ID.String(),
		)

		body :=
			helper.ExpirationReminderEmail(
				permit.DocumentName,
				url,
				permit.ExpiredAt.Format(
					"2006-01-02",
				),
			)

		subject :=
			"Document Expiration Reminder"

		if daysLeft < 0 {

			subject =
				"Document Already Expired"
		}

		s.EmailService.SendAsync(
			emails,
			subject,
			body,
		)
	}
}
