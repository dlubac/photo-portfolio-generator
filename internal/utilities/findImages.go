package utilities

import (
	"dlubac/photo-portfolio-generator/internal/structs"
	"log"
	"path/filepath"
	"strconv"
	"strings"
)

func FindImages(path string) ([]string, error) {
	log.Printf("Finding images in: %v", path)
	files, err := filepath.Glob(path + "/*.jpg")
	if err != nil {
		return nil, err
	}

	for i, file := range files {
		if file == path+"/cover.jpg" {
			files = append(files[:i], files[i+1:]...)
			break
		}
	}

	sortedFiles := []structs.SortedImage{}
	for _, file := range files {
		index, _ := strconv.Atoi(strings.Replace(file[len(file)-6:len(file)-4], "_", "", -1))
		sortedFiles = append(sortedFiles, structs.SortedImage{Name: file, Index: index})
	}

	for i := 0; i < len(sortedFiles); i++ {
		for j := i + 1; j < len(sortedFiles); j++ {
			if sortedFiles[i].Index > sortedFiles[j].Index {
				temp := sortedFiles[i]
				sortedFiles[i] = sortedFiles[j]
				sortedFiles[j] = temp
			}
		}
	}

	returnFiles := []string{}
	for _, file := range sortedFiles {
		returnFiles = append(returnFiles, file.Name)
	}

	return returnFiles, nil
}
