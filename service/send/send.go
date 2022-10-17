package sendler

import (
	"crypto/tls"

	"gopkg.in/gomail.v2"
)

type Dialer struct {
}

func Send() {
	d := gomail.NewDialer("smtp.gmail.com", 587, "mailerk62@gmail.com", "lohzgrkpegwsffol")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	m := gomail.NewMessage()
	m.SetHeaders(map[string][]string{
		"From":    {m.FormatAddress("mailerk62@gmail.com", "mailerk62")},
		"To":      {"mailerk62@gmail.com"},
		"Subject": {"Hello"},
	})
	m.SetAddressHeader("Cc", "surnamename615@gmail.com", "Maksim")
	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	// Send emails using d.
}
