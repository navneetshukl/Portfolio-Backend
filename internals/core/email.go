package core

import "context"

type SendEmail struct {
	Subject string `json:"subject"`
	Message string `json:"message"`
	Mail    string `json:"email"`
}

type EmailUseCase interface {
	SendEmail(ctx context.Context, req *SendEmail) error
}
