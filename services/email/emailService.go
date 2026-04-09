package email

import (
	"belajar-go/models"
	"belajar-go/utils"
)

type EmailService interface {
	SendEmail(to, subject, body string)
}

type EmailServiceImpl struct{}

func NewEmailService() *EmailServiceImpl {
	return &EmailServiceImpl{}
}

func (s *EmailServiceImpl) SendEmail(to, subject, body string) {
	job := models.EmailJob{
		To:      to,
		Subject: subject,
		Body:    body,
	}

	utils.EnqueueEmail(job)
}
