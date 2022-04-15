package controller

import (
	"log"
	"os"
	"strconv"

	"github.com/tubes/Art-Auction-Tubes/model"
	"gopkg.in/gomail.v2"
)

const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587

var CONFIG_SENDER_NAME = os.Getenv("CONFIG_SENDER_NAME")
var CONFIG_AUTH_EMAIL = os.Getenv("CONFIG_AUTH_EMAIL")
var CONFIG_AUTH_PASSWORD = os.Getenv("CONFIG_AUTH_PASSWORD")

func sendMail(email model.ListEmail) {
	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD,
	)
	s, err := dialer.Dial()
	if err != nil {
		panic(err)
	}

	mailer := gomail.NewMessage()
	for _, r := range email {
		mailer.SetHeader("From", CONFIG_SENDER_NAME)
		mailer.SetAddressHeader("To", r.Email, r.Name)
		mailer.SetHeader("Subject", "You won the bid !!!")
		mailer.SetBody("text/html", "Hello "+r.Name+" !"+" you won the bid at "+r.Date.String()+" with "+strconv.FormatFloat(r.Etherium, 'f', 2, 64)+" etherium.")

		if err := gomail.Send(s, mailer); err != nil {
			log.Printf("Could not send email to %q: %v", r.Email, err)
		}
		mailer.Reset()
	}
	log.Println("Mail sent!")
}
