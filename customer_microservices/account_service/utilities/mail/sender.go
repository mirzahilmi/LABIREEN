package mail

import (
	"bytes"
	"fmt"
	"log"
	"net/smtp"

	"github.com/jordan-wright/email"
)

const (
	smtpAuthAddress   = "smtp.gmail.com"
	smtpServerAddress = "smtp.gmail.com:587"
)

type EmailSender interface {
	SendEmail(data *EmailData) error
}

type GmailSender struct {
	name              string
	fromEmailAddress  string
	fromEmailPassword string
}

type EmailData struct {
	Email   []string
	URL     string
	Name    string
	Subject string
}

func NewGmailSender(name string, email string, pwd string) EmailSender {
	return &GmailSender{
		name:              name,
		fromEmailAddress:  email,
		fromEmailPassword: pwd,
	}
}

func (sender *GmailSender) SendEmail(data *EmailData) error {
	var body bytes.Buffer

	template, err := ParseTemplateDir("templates")
	if err != nil {
		log.Fatal("Could not parse template", err)
	}

	template.ExecuteTemplate(&body, "verificationCode.html", &data)

	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", sender.name, sender.fromEmailAddress)
	e.Subject = data.Subject
	e.HTML = []byte(body.String())
	e.To = data.Email

	smtpAuth := smtp.PlainAuth("", sender.fromEmailAddress, sender.fromEmailPassword, smtpAuthAddress)
	return e.Send(smtpServerAddress, smtpAuth)
}
