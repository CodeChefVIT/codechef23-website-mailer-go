package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/mr-emerald-wolf/mailer-go/initializers"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type SendEmailRequest struct {
	Subject string `json:"subject" validate:"required"`
	Body    string `json:"body" validate:"required"`
}

func init() {
	initializers.LoadEnvVariables()
}

func SendMail(s string, b string) error {
	from := mail.NewEmail("No Reply @CodeChef-VIT", os.Getenv("FROM_EMAIL"))
	subject := s
	to := mail.NewEmail("CodeChef-VIT", os.Getenv("TO_EMAIL"))
	plainTextContent := b

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, "")
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)

	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}

	return err
}
