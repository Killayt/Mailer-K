package load

import (
	"fmt"
	"log"
	"os"
)

type Subject []byte

func CreateSubject() Subject {
	var subject string
	fmt.Println("Enter subject: ")
	fmt.Fscan(os.Stdin, &subject)

	return Subject(subject)
}

func LoadMaket(path string) (file []byte) {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Panic("File not found", err)
	}

	return file
}
