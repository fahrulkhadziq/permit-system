package service

import (
	"fmt"
	"os"
	"permit-license/config"

	"github.com/go-mail/mail/v2"
)

type EmailService struct{}

func (s *EmailService) SendEmail(to []string, subject string, body string) error {
	m := mail.NewMessage()

	m.SetHeader("From", os.Getenv("MAIL_FROM"))
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	dialer := config.GetMailer()
	return dialer.DialAndSend(m)
}

func (s *EmailService) SendAsync(to []string, subject string, body string) {

	go func() {
		err := s.SendEmail(to, subject, body)

		if err != nil {
			fmt.Println("EMAIL ERROR: ", err.Error())
		}
	}()
}
