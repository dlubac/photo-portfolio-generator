package utilities

import (
	"io"
	"log"
	"os"
)

func CopyFile(source string, destination string) error {
	sourceFile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	newFile, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer newFile.Close()

	_, err = io.Copy(newFile, sourceFile)
	if err != nil {
		return err
	}

	err = newFile.Sync()
	if err != nil {
		return err
	}

	log.Printf("Copied file: %v to: %v", source, destination)

	return nil
}
