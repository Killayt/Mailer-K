package main

import (
	"fmt"
	"log"
	"mailerK/service/load"
	"os"

	"gopkg.in/gomail.v2"
)

func main() {
	// registration dialer (from who we will send)
	dialer := load.Registration()
	fmt.Println("\n[SUCCESS] Dialer declorated by " + dialer.Username + " on host " + dialer.Host + "\n")

	var PathToMaket []byte
	fmt.Print("Set the full PATH to html MAKET here: ")
	fmt.Fscan(os.Stdin, &PathToMaket)
	maket := load.LoadMaket(string(PathToMaket))

	fmt.Println("\n" + "[SUCCESS] Maket was found !")

	// sending
	message := gomail.NewMessage()
	message.SetHeader("From", dialer.Username)
	message.SetHeader("To", dialer.Username)

	// generate subject
	subject := load.CreateSubject()

	message.SetHeader("Subject", subject) //generate subject
	message.SetBody("text/html", string(maket))

	if err := dialer.DialAndSend(message); err != nil {
		log.Panic("[ERROR]"+"not send", err)
	}
}
