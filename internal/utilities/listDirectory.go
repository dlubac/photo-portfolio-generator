package utilities

import (
	"log"
	"os"
	"regexp"
	"slices"
)

func ListDirectory(path string, regexPattern string, onlyDirectories bool) ([]string, error) {
	var results []string

	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	if onlyDirectories {
		for i, entry := range entries {
			if !entry.IsDir() {
				slices.Delete(entries, i, i+1)
			}
		}
	}

	for _, entry := range entries {
		if entry == nil {
			continue
		}

		if regexPattern == "" {
			results = append(results, entry.Name())
			continue
		}

		matched, err := regexp.MatchString(regexPattern, entry.Name())
		if err != nil {
			log.Printf("Error parsing regex: %s", err)
		}
		if matched {
			results = append(results, entry.Name())
		}
	}

	return results, nil
}
