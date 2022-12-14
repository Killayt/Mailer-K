package load

import (
	"crypto/tls"
	"fmt"
	"os"

	"gopkg.in/gomail.v2"
)

func Registration() *gomail.Dialer {
	fmt.Println("Let`s register to service :)")

	var YouMail string
	fmt.Println("Enter email sandler: ")
	fmt.Fscan(os.Stdin, &YouMail)

	var Passwd string
	fmt.Println("Enter password to email sandler: ")
	fmt.Fscan(os.Stdin, &Passwd)

	Dialer := gomail.NewDialer("smtp.gmail.com", 587, YouMail, Passwd)
	Dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	return Dialer

}
