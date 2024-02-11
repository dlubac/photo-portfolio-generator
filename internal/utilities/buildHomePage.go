package utilities

import (
	"dlubac/photo-portfolio-generator/internal/structs"
	"html/template"
	"log"
	"os"
)

func BuildHomePage(config *structs.Config) error {
	log.Print("Building home page")
	galleries := []structs.HomePageGallery{}
	for _, gallery := range config.Galleries {
		galleries = append(galleries, structs.HomePageGallery{
			Gallery:           gallery,
			IndexPath:         "galleries/" + gallery.Path + "/index.html",
			CoverImagePath:    "galleries/" + gallery.Path + "/cover.jpg",
			CoverImageAltText: gallery.Description,
		})

		err := CopyFile("images/"+gallery.Path+"/cover.jpg", config.OutputDirectory+"/galleries/"+gallery.Path+"/cover.jpg")
		if err != nil {
			log.Fatalf("Unable to copy cover image: %v", err)
		}
	}

	tmpl, err := template.ParseFiles("templates/homepage.html")
	if err != nil {
		return err
	}

	file, err := os.Create(config.OutputDirectory + "/index.html")
	if err != nil {
		log.Printf("Unable to create file: %v", err)
		return err
	}

	err = tmpl.Execute(file, structs.HomePage{
		PageTitle:         config.SiteTitle,
		PageDescription:   config.SiteDescription,
		Domain:            config.Domain,
		HomePageGalleries: galleries,
		FooterLinks:       config.FooterLinks,
	})
	if err != nil {
		log.Printf("Unable to execute template: %v", err)
		return err
	}

	log.Print("Built home page")

	return nil
}
