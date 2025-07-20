package usecase

import (
	"context"
	"log"
	email "portfolio/internals/adapter/external"
	"portfolio/internals/config"
	"portfolio/internals/core"
)

type EmailUseCaseImpl struct {
	conf    *config.Config
	mailSvc email.MailSvc
}

func NewEmailUseCase(c *config.Config, mailSvc email.MailSvc) core.EmailUseCase {
	return &EmailUseCaseImpl{
		conf:    c,
		mailSvc: mailSvc,
	}
}

func (m *EmailUseCaseImpl) SendEmail(ctx context.Context, req *core.SendEmail) error {
	if req == nil {
		return core.ReqBodyNotPresent
	}
	if req.Subject == "" || req.Body == "" {
		return core.InvalidRequestBody
	}

	emailReq := &email.MailPayLoad{
		Subject: req.Subject,
		Body:    req.Body,
	}

	err := m.mailSvc.SendMail(ctx, emailReq)
	if err != nil {
		log.Println("error in sending the mail ", err)
		return core.ErrorSendingMail
	}
	return nil
}
