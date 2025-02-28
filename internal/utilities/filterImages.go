package utilities

import (
	"dlubac/photo-portfolio-generator/internal"
	"strings"
)

func stringContainsAnySubstring(input string, substrings []string) bool {
	for _, substring := range substrings {
		if strings.Contains(input, substring) {
			return true
		}
	}
	return false
}

func imageHasValidExtension(image string, extensions []string) bool {
	for _, extension := range extensions {
		if strings.HasSuffix(image, extension) {
			return true
		}
	}
	return false
}

func FilterImages(images []string, excludedStrings []string) []string {
	var filteredImages []string

	for _, image := range images {
		if !stringContainsAnySubstring(image, excludedStrings) && imageHasValidExtension(image, internal.ImageFileExtensions) {
			filteredImages = append(filteredImages, image)
		}
	}

	return filteredImages
}
