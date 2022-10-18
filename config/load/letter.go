package load

import (
	"fmt"
	"os"
)

type Subject []byte

func CreateSubject() Subject {
	var subject string
	fmt.Println("Enter subject: ")
	fmt.Fscan(os.Stdin, &subject)

	return Subject(subject)
}
