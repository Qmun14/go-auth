package services

import (
	"fmt"
	"log"

	"github.com/Qmun14/jwtAuth/mail"
	"github.com/Qmun14/jwtAuth/utils"
)

func CreateVerifyEmail(emailID int64, secretCode string, userEmail string, userName string) error {
	config, err := utils.LoadConfig("..")
	if err != nil {
		log.Fatal("tidak bisa memuat config:", err)
	}
	mailer := mail.NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)

	subject := "Welcome to User Management App"

	verifyUrl := fmt.Sprintf("http://%s/api/verify_email?email_id=%d&secret_code=%s", config.ServerAddress, emailID, secretCode)
	content := fmt.Sprintf(`Hello %s,<br/>
	Thank you for registering with us!<br/>
	Please <a href="%s">click here</a> to verify your email address.<br/>
	`, userName, verifyUrl)
	to := []string{userEmail}

	err = mailer.SendEmail(subject, content, to, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to send verify email: %w", err)
	}
	return nil
}
