package config

import (
	"crypto/tls"
	"log"

	"gopkg.in/gomail.v2"
)

func NewMail() *Mail {
	d := gomail.NewDialer("sandbox.smtp.mailtrap.io", 587, "62ae93b062182d", "102fc868c74199")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	return &Mail{d: d}
}

type Mail struct {
	d *gomail.Dialer
}

func (m *Mail) SendMail(email, body string) {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", "adit@mail.com")
	mailer.SetHeader("To", email)
	mailer.SetHeader("Subject", "Account Verification")
	mailer.SetBody("text/html", body)
	err := m.d.DialAndSend(mailer)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Print("Mail sent")
}
