package services

import (
	"fmt"
	"github.com/resend/resend-go/v2"
	"kratomTracker/notificationmanager"
)

type EmailNotificationService struct {
	resendAPIKey string
	emails       []string
	fromEmail    string
	client       *resend.Client
}

func NewEmailNotificationService(apiKey, fromEmail string) *EmailNotificationService {
	client := resend.NewClient(apiKey)
	return &EmailNotificationService{
		fromEmail:    fromEmail,
		client:       client,
		resendAPIKey: apiKey,
	}
}

func (service *EmailNotificationService) AddEmail(email string) {
	service.emails = append(service.emails, email)
}

func (service *EmailNotificationService) SendNotification(n notificationmanager.NotificationObject) error {
	for _, email := range service.emails {
		fmt.Println("Sending email to: " + email)
		req := &resend.SendEmailRequest{
			From:        service.fromEmail,
			To:          []string{email},
			Subject:     "New Notification From Kratom Dose Manager",
			Bcc:         nil,
			Cc:          nil,
			ReplyTo:     service.fromEmail,
			Html:        "",
			Text:        n.Content,
			Tags:        nil,
			Attachments: nil,
			Headers:     nil,
			ScheduledAt: "",
		}
		snt, err := service.client.Emails.Send(req)
		if err != nil {
			_ = fmt.Errorf("Error sending email: %s\n", err)
			return err
		}
		fmt.Println("Email sent: ", snt.Id)
	}
	return nil
}
