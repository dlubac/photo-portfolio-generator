package main

import (
	"dlubac/photo-portfolio-generator/internal/structs"
	"dlubac/photo-portfolio-generator/internal/utilities"
	"html/template"
	"log"
	"os"
	"path/filepath"
)

func main() {
	//var galleries []structs.Gallery

	galleriesPath := filepath.Join("content", "galleries") + string(filepath.Separator)
	matches, err := filepath.Glob(galleriesPath + "*")
	if err != nil {
		log.Fatal(err)
	}

	var galleries []structs.Gallery
	for _, match := range matches {
		info, err := os.Stat(match)
		if err != nil {
			log.Fatal(err)
		}

		if info.IsDir() {
			gallery, err := utilities.BuildGallery(match)
			if err != nil {
				log.Fatal(err)
			}

			galleries = append(galleries, gallery)
		}
	}

	tmpl, err := template.ParseFiles("templates/homepage2.html")
	if err != nil {
		log.Fatal(err)
	}

	homepage, err := os.Create("output/homepage2.html")
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.Execute(homepage, structs.Homepage{Galleries: galleries})
	if err != nil {
		log.Fatal(err)
	}

	//for _, name := range galleryNames {
	//	gallery, err := utilities.BuildGallery(name)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	fmt.Println(gallery)
	//}
}
