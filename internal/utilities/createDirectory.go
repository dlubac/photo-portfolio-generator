package utilities

import (
	"log"
	"os"
)

func CreateDirectory(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.Mkdir(path, 0755)
		if err != nil {
			log.Printf("Unable to create output directory: %v", err)
			return err
		}
	}

	log.Printf("Created directory: %v", path)

	return nil
}
