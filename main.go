package main

import (
	"crypto/tls"

	"mailerK/service/load"

	"gopkg.in/gomail.v2"
)

func main() {
	maket := load.GetHtmlMaket("/home/killayt/projects/Mailer-K/service/makets/fitst/user.html")

	d := gomail.NewDialer("smtp.gmail.com", 587, "mailerk62@gmail.com", "lohzgrkpegwsffol")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	m := gomail.NewMessage()
	m.SetHeaders(map[string][]string{
		"From":    {m.FormatAddress("mailerk62@gmail.com", "mailerk62")},
		"To":      {"mailerk62@gmail.com"},
		"Subject": {"Hello"},
	})
	m.SetAddressHeader("Cc", "surnamename615@gmail.com", "Maksim")
	m.SetBody("text/html", string(maket))

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	// Send emails using d.
}
