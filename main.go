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
	template, err := load.LoadTemplate(string(PathToMaket))
	if err != nil {
		log.Println("Template not found")
	}

	fmt.Println("\n" + "[SUCCESS] Maket was found !")

	// creating subject
	subject := load.CreateSubject()

	// making message
	s, err := dialer.Dial()
	if err != nil {
		panic(err)
	}

	// var wg sync.WaitGroup

	m := gomail.NewMessage()

	for _, subscriber := range sub.Subscribers {

		m.SetHeader("From", dialer.Username)
		m.SetAddressHeader("To", subscriber.Email, subscriber.Name)
		m.SetHeader("Subject", string(subject))

		m.SetBody("text/html", template)

		if err := gomail.Send(s, m); err != nil {
			log.Printf("Could not send email to %q: %v", subscriber.Email, err)
		}
		m.Reset()
		log.Printf("message was send to %v to the person %v (phone: %v)", subscriber.Email, subscriber.Name, subscriber.Phone)

	}

}
