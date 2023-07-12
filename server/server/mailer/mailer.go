package mailer

import (
	"fmt"

	"github.com/anthdm/hollywood/actor"
	"github.com/sijibomii/cryptopay/core/models"
	"gopkg.in/gomail.v2"
)

type SendActivationMailMessage struct {
	Payload models.User
}

type Mailer struct {
	SmtpHost     string
	SmtpPort     int
	SmtpUsername string
	SmtpPassword string
}

func sendActivationEmail(m *Mailer, user models.User) {
	// Sender and recipient details
	senderName := "Sijibomi"
	senderEmail := "hello@sijibomi.com"
	subject := "Activate Your Account"

	// Email body
	body := fmt.Sprintf(`Please click the following link to activate your account: <a href=http://localhost:3000/activate/"%s">http://localhost:3000/activate/%s</a>.`, user.Verification_token, user.Verification_token)

	// Create a new message
	message := gomail.NewMessage()
	message.SetHeader("From", fmt.Sprintf("%s <%s>", senderName, senderEmail))
	message.SetHeader("To", user.Email)
	message.SetHeader("Subject", subject)
	message.SetBody("text/html", body)

	// Create a new SMTP dialer
	dialer := gomail.NewDialer(m.SmtpHost, m.SmtpPort, m.SmtpUsername, m.SmtpPassword)

	// Send the email
	if err := dialer.DialAndSend(message); err != nil {
		//fmt.Printf("Mailer error: %s", err.Error())
	}
}

func (m *Mailer) Receive(ctx *actor.Context) {
	switch l := ctx.Message().(type) {
	case actor.Started:
		//fmt.Println("Mailer started")

	case SendActivationMailMessage:
		sendActivationEmail(m, l.Payload)

	default:
		//fmt.Println("UNKNOWN MESSAGE TO MAILER")

	}
}
