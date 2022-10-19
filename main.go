package main

import (
	"fmt"
	"log"
	"mailerK/config"
	"mailerK/config/load"
	"os"

	"gopkg.in/gomail.v2"
)

func main() {

	// registration dialer (from who we will send)
	dialer := load.Registration()
	fmt.Println("\n[SUCCESS] Dialer declorated by " + dialer.Username + " on host " + dialer.Host + "\n")

	// getting base pf subscribers
	sub, err := config.LoadSubscribers("./config/subsctibers.json")
	if err != nil {
		log.Panic(err)
	}

	// getting maket
	var PathToMaket []byte
	fmt.Print("Set the full PATH to html MAKET here: ")
	fmt.Fscan(os.Stdin, &PathToMaket)
	maket := load.LoadMaket(string(PathToMaket))

	fmt.Println("\n" + "[SUCCESS] Maket was found !")

	// making message
	s, err := dialer.Dial()
	if err != nil {
		panic(err)
	}

	m := gomail.NewMessage()
	for _, r := range sub.Subscribers {
		m.SetHeader("From", "no-reply@example.com")
		m.SetAddressHeader("To", r.Email, r.Name)
		m.SetHeader("Subject", "Newsletter #1")
		m.SetBody("text/html", string(maket))

		if err := gomail.Send(s, m); err != nil {
			log.Printf("Could not send email to %q: %v", r.Email, err)
		}
		m.Reset()
		log.Printf("message was send to %v to the person %v (phone: %v)", r.Email, r.Name, r.Phone)
	}

}
