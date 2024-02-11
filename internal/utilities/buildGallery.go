package utilities

import (
	"dlubac/photo-portfolio-generator/internal/structs"
	"html/template"
	"log"
	"os"
	"strings"
)

func BuildGallery(gallery *structs.Gallery, config *structs.Config) error {
	outputFilePath := config.OutputDirectory + "/galleries/" + gallery.Path

	err := CreateDirectory(outputFilePath)
	if err != nil {
		log.Fatalf("Unable to create output directory: %v", err)
	}

	images, err := FindImages("images/" + gallery.Path)
	if err != nil {
		log.Fatalf("Unable to find images in path: %v", err)
	}
	log.Printf("Found %v images in: %v", len(images), "images/"+gallery.Path)

	for _, image := range images {
		err = CopyFile(image, strings.Replace(image, "images/"+gallery.Path, outputFilePath, 1))
		if err != nil {
			log.Fatalf("Unable to copy image: %v", err)
		}
	}

	galleryImages := make([]structs.GalleryPageImage, len(images))
	for i, image := range images {
		log.Printf("Processing image: %v", image)
		metadata, _ := GetImageMetadata(image, config)
		altText, err := GetExifTag(image, "ImageDescription")
		if err != nil {
			altText = ""
		}

		galleryImages[i] = structs.GalleryPageImage{
			Url:      strings.Replace(image, "images/"+gallery.Path, ".", 1),
			Metadata: FormatImageMetadata(metadata),
			AltText:  TrimQuotes(altText),
		}
	}

	galleryPage := &structs.GalleryPage{
		PageTitle:       config.SiteTitle,
		PageDescription: gallery.Description,
		Url:             gallery.Path,
		Domain:          config.Domain,
		Images:          galleryImages,
		FooterLinks:     config.FooterLinks,
	}

	tmpl, err := template.ParseFiles("templates/gallery.html")
	if err != nil {
		return err
	}

	file, err := os.Create(outputFilePath + "/index.html")
	if err != nil {
		log.Printf("Unable to create file: %v", err)
		return err
	}

	err = tmpl.Execute(file, galleryPage)
	if err != nil {
		return err
	}

	log.Printf("Gallery built: %v", gallery.Name)

	return nil
}
