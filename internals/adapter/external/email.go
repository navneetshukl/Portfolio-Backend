package email

import (
	"context"
	"fmt"
	"portfolio/internals/config"
	"strconv"

	"gopkg.in/mail.v2"
)

type MailPayLoad struct {
	Subject string `json:"subject"`
	Body    string `json:"body"`
	Mail string `json:"email"`
}

type Mail struct {
	conf *config.Config
}

type MailSvc interface {
	SendMail(ctx context.Context, req *MailPayLoad) error
}

func NewMailSvc(conf *config.Config) MailSvc {
	return &Mail{
		conf: conf,
	}
}

func (m *Mail) SendMail(ctx context.Context, req *MailPayLoad) error {

	from := m.conf.EmailConfig.FROM_EMAIL_ADDRESS
	appPassword := m.conf.EmailConfig.EMAIL_API_PASSWORD
	to := m.conf.EmailConfig.TO_EMAIL_ADDRESS
	smtpHost := m.conf.EmailConfig.SMTP_HOST
	smtpPort := m.conf.EmailConfig.SMTP_PORT

	port, err := strconv.Atoi(smtpPort)
	if err != nil {
		return err
	}

	req.Body=fmt.Sprintf("%s || FROM : %s",req.Body,req.Mail)

	// Create a new message
	message := mail.NewMessage()
	message.SetHeader("From", from)
	message.SetHeader("To", to)
	message.SetHeader("Subject", req.Subject)
	message.SetBody("text/plain", req.Body)

	// Set up the SMTP dialer
	dialer := mail.NewDialer(smtpHost, port, from, appPassword)

	// Send the email
	if err := dialer.DialAndSend(message); err != nil {
		return err
	}
	fmt.Println("Email sent successfully!")
	return nil
}
