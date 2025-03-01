package main

import (
	"dlubac/photo-portfolio-generator/internal/steps"
	"dlubac/photo-portfolio-generator/internal/structs"
	"dlubac/photo-portfolio-generator/internal/utilities"
	"flag"
	"log"
	"os"
	"path/filepath"
)

func main() {
	domain := flag.String("domain", "", "domain")
	title := flag.String("title", "", "site title")
	description := flag.String("description", "", "site description")
	flag.Parse()

	metadata := structs.SiteMetadata{
		Domain:      *domain,
		Title:       *title,
		Description: *description,
	}

	galleriesPath := filepath.Join("content", "galleries") + string(filepath.Separator)
	matches, err := filepath.Glob(galleriesPath + "*")
	if err != nil || len(matches) == 0 {
		log.Fatalf("Error searching for galleries: %v", err)
	}

	err = steps.PrepareDirectories()
	if err != nil {
		log.Fatal(err)
	}

	err = steps.CopyStaticFiles()
	if err != nil {
		log.Fatal(err)
	}

	var galleries []structs.Gallery
	for _, match := range matches {
		info, _ := os.Stat(match)
		if info.IsDir() {
			gallery, err := steps.BuildGallery(match, metadata)
			if err != nil {
				log.Printf("Error building gallery: %v", err)
			}

			galleries = append(galleries, gallery)
		}
	}

	photoReelPreviews, err := steps.BuildPhotoReel(metadata)
	if err != nil {
		log.Fatal(err)
	}

	err = utilities.BuildTemplate(
		"templates/homepage.html",
		"output/index.html",
		structs.Homepage{PhotoReelPreviews: photoReelPreviews, Galleries: galleries, Metadata: metadata})
	if err != nil {
		log.Fatalf("Error building template: %s\n", err)
	}
}
