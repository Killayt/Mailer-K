package main

import (
	"log"
	"net/smtp"

	"mailerK/mail/"
)

var toSend = mail.NewTarget([]string{"surnamename615@gmail.com"})
var letter = mail.NewLetter("Тема", "тело", toSend.Target)

func sendMailSimple(subject string, body string, to []string) {
	auth := smtp.PlainAuth(
		"",
		"mailerk62@gmail.com",
		"lohzgrkpegwsffol",
		"smtp.gmail.com",
	)

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"mailerk62@gmail.com",
		[]string{"mailerk62@gmail.com"},
		[]byte(toSend),
	)

	if err != nil {
		log.Panic("Something wrong\n", err.Error())
	}
}

func main() {
	sendMailSimple("another subject", "another body here", []string{"surnamename615@gmail.com"})
}
