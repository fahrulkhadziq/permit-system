package config

import (
	"os"
	"strconv"

	"github.com/go-mail/mail/v2"
)

func GetMailer() *mail.Dialer {

	port, _ := strconv.Atoi(os.Getenv("MAIL_PORT"))

	return mail.NewDialer(
		os.Getenv("MAIL_HOST"),
		port,
		os.Getenv("MAIL_USERNAME"),
		os.Getenv("MAIL_PASSWORD"),
	)
}
