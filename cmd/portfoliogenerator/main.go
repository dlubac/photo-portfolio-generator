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

	steps.PrepareDirectories()
	steps.CopyStaticFiles()

	var galleries []structs.Gallery
	for _, match := range matches {
		info, _ := os.Stat(match)
		if info.IsDir() {
			galleries = append(galleries, steps.BuildGallery(match, metadata))
		}
	}

	utilities.BuildTemplate(
		"templates/homepage.html",
		"output/index.html",
		structs.Homepage{PhotoReelPreviews: steps.BuildPhotoReel(metadata), Galleries: galleries, Metadata: metadata})
}
