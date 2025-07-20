package core

import "context"

type SendEmail struct {
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

type EmailUseCase interface {
	SendEmail(ctx context.Context, req *SendEmail) error
}
