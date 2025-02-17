package main

import (
	"dlubac/photo-portfolio-generator/internal/steps"
	"dlubac/photo-portfolio-generator/internal/structs"
	"html/template"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// TODO
	// - domain name cli param
	// - site title cli param
	//
	metadata := structs.SiteMetadata{
		Domain:      "dlubac.com",
		Title:       "Matt Dlubac Photography",
		Description: "Outdoor&#32;photography&#32;by&#32;Matt&#32;Dlubac",
	}

	galleriesPath := filepath.Join("content", "galleries") + string(filepath.Separator)
	matches, err := filepath.Glob(galleriesPath + "*")
	if err != nil {
		log.Fatalf("Error searching for galleries: %v", err)
	}

	if len(matches) == 0 {
		log.Fatal("No galleries found")
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

	tmpl, err := template.ParseFiles("templates/homepage.html")
	if err != nil {
		log.Fatal(err)
	}

	homepage, err := os.Create("output/index.html")
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.Execute(homepage, structs.Homepage{Galleries: galleries, Metadata: metadata})
	if err != nil {
		log.Fatal(err)
	}
}
