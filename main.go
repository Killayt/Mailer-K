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

	// getting maket
	var PathToMaket []byte
	fmt.Print("Set the full PATH to html MAKET here: ")
	fmt.Fscan(os.Stdin, &PathToMaket)
	maket := load.LoadMaket(string(PathToMaket))

	fmt.Println("\n" + "[SUCCESS] Maket was found !")

	// making message
	message := gomail.NewMessage()
	message.SetHeader("From", dialer.Username)

	sub, err := config.LoadSubscribers("./config/subsctibers.json")
	if err != nil {
		log.Panic(err)
	}

	message.SetHeader("To", sub.Mail)

	// generate subject
	subject := load.CreateSubject()
	message.SetHeader("Subject", string(subject)) //generate subject
	message.SetBody("text/html", string(maket))

	if err := dialer.DialAndSend(message); err != nil {
		log.Panic(err)
	}
	log.Printf("message was send to %v to the person %v", sub.Mail, sub.Name)
}
