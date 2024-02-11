package utilities

import (
	"log"
	"os"
)

func DeleteDirectory(path string) error {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		err = os.RemoveAll(path)
		if err != nil {
			log.Printf("Unable to delete directory: %v", err)
			return err
		}
	}

	log.Printf("Deleted directory: %v", path)

	return nil
}
