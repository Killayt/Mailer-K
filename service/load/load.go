package load

import (
	"log"
	"os"
)

func GetHtmlMaket(path string) []byte {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Panic("File not found\n", err)
	}

	return file
}
