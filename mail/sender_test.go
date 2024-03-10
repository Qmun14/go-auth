package mail

import (
	"testing"

	"github.com/Qmun14/jwtAuth/utils"
	"github.com/stretchr/testify/assert"
)

func TestSendEmail(t *testing.T) {
	config, err := utils.LoadConfig("..")
	assert.NoError(t, err)

	sender := NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)

	subject := "A test email"
	content := `
	<h1>Hello World</h1>
	<p>This is a test message from <a href="https://mamun-portfolio.vercel.app/"> Ma'mun Ramdhan</a></p>
	`
	to := []string{"mamunramdhan@gmail.com"}
	bcc := []string{"bcc@test.com"}
	cc := []string{"ccd@test.com"}
	attachFiles := []string{"../README.md"}

	err = sender.SendEmail(subject, content, to, cc, bcc, attachFiles)
	assert.NoError(t, err)

}
