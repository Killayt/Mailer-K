package load

import (
	"log"
	"os"
)

func LoadMaket(path string) (file []byte) {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Panic("File not found", err)
	}

	return file
}
