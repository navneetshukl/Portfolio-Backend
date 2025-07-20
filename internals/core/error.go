package core

import "errors"

var (
	ReqBodyNotPresent  error = errors.New("request body not present")
	ErrorSendingMail   error = errors.New("error in sending the mail")
	InvalidRequestBody error = errors.New("request body invalid")
)
