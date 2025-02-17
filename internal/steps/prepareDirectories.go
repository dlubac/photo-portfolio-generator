package steps

import (
	"dlubac/photo-portfolio-generator/internal/utilities"
)

func PrepareDirectories() error {
	directoriesToCreate := []string{"output", "output/galleries", "output/photo-reel"}

	err := utilities.DeleteDirectory("output")
	if err != nil {
		return err
	}

	for _, directory := range directoriesToCreate {
		err = utilities.CreateDirectory(directory)
		if err != nil {
			return err
		}
	}

	return nil
}
