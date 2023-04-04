package utils

import (
	"crypto/tls"
	"fmt"
	"os"

	"github.com/mr-emerald-wolf/mailer-go/initializers"
	"gopkg.in/gomail.v2"
)

type SendEmailRequest struct {
	Subject string `json:"subject" validate:"required"`
	Body    string `json:"body" validate:"required"`
}

func init() {
	initializers.LoadEnvVariables()
}

func SendMail(s string, b string) error {
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", os.Getenv("SENDER_MAIL"))

	// Set E-Mail receivers
	m.SetHeader("To", os.Getenv("RECIEVER_MAIL"))

	// Set E-Mail subject
	m.SetHeader("Subject", s)

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", b)

	// Settings for SMTP server
	d := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("SENDER_MAIL"), os.Getenv("SENDER_PASS"))

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	err := d.DialAndSend(m)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
