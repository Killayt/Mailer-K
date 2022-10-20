package load

import (
	"fmt"
	"log"
	"os"
)

func CreateSubject() string {
	var subject string
	fmt.Println("Enter subject: ")
	fmt.Fscan(os.Stdin, &subject)

	return subject
}

func LoadTemplate(path string) (string, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		log.Println("File not found", err)
		LoadTemplate(path)
	}

	return string(bytes), nil
}
