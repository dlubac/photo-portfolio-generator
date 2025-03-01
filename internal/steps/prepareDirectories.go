package steps

import (
	"dlubac/photo-portfolio-generator/internal/utilities"
	"log"
)

func PrepareDirectories() {
	directoriesToCreate := []string{"output", "output/galleries", "output/photo-reel"}

	err := utilities.DeleteDirectory("output")
	if err != nil {
		log.Fatalf("Error deleting output directory: %v", err)
	}

	for _, directory := range directoriesToCreate {
		err = utilities.CreateDirectory(directory)
		if err != nil {
			log.Fatalf("Error creating directory: %v", err)
		}
	}
}
