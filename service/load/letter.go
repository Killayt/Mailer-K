package load

import (
	"fmt"
	"os"
)

func CreateSubject() string {
	var subject string
	fmt.Println("Enter subject: ")
	fmt.Fscan(os.Stdin, &subject)

	return subject
}
