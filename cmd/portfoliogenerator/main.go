package main

import (
	"dlubac/photo-portfolio-generator/internal/utilities"
	"log"
)

func main() {
	config, err := utilities.ParseConfig("config.yml")
	if err != nil {
		log.Fatalf("Unable to parse config: %v", err)
	}

	err = utilities.DeleteDirectory(config.OutputDirectory)
	if err != nil {
		log.Fatalf("Unable to delete output directory: %v", err)
	}

	err = utilities.CreateDirectory(config.OutputDirectory)
	if err != nil {
		log.Fatalf("Unable to create output directory: %v", err)
	}

	err = utilities.CreateDirectory(config.OutputDirectory + "/galleries")
	if err != nil {
		log.Fatalf("Unable to create galleries directory: %v", err)
	}

	err = utilities.CopyFile("templates/styles.css", config.OutputDirectory+"/styles.css")
	if err != nil {
		log.Fatalf("Unable to copy styles.css: %v", err)
	}

	err = utilities.CopyFile("templates/iAWriterDuoS-Regular.woff2", config.OutputDirectory+"/iAWriterDuoS-Regular.woff2")
	if err != nil {
		log.Fatalf("Unable to copy iAWriterDuoS-Regular.woff2: %v", err)
	}

	err = utilities.CopyFile("templates/favicon.ico", config.OutputDirectory+"/favicon.ico")
	if err != nil {
		log.Fatalf("Unable to copy favicon.ico: %v", err)
	}

	galleries := config.Galleries
	for _, Gallery := range galleries {
		log.Printf("Building gallery: %v", Gallery.Name)
		err = utilities.BuildGallery(&Gallery, config)
		if err != nil {
			log.Fatalf("Unable to build gallery: %v", err)
		}
	}

	err = utilities.BuildHomePage(config)
	if err != nil {
		log.Fatalf("Unable to build home page: %v", err)
	}
}
