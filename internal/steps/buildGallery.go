package steps

import (
	"dlubac/photo-portfolio-generator/internal"
	"dlubac/photo-portfolio-generator/internal/structs"
	"dlubac/photo-portfolio-generator/internal/utilities"
	"errors"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"log"
	"path/filepath"
	"sort"
	"strings"
)

func BuildGallery(path string, metadata structs.SiteMetadata) (structs.Gallery, error) {
	name := strings.Split(path, string(filepath.Separator))[2]
	log.Printf("Building gallery %s\n", name)

	outputDirectory := strings.Replace(path, "content", "output", 1)
	err := utilities.CreateDirectory(outputDirectory)

	coverImages, err := filepath.Glob(path + string(filepath.Separator) + "*_cover.*")
	if err != nil {
		return structs.Gallery{}, err
	}

	if len(coverImages) == 0 {
		return structs.Gallery{}, errors.New("Cover image not found for gallery " + path)
	}
	coverImagePath := coverImages[0]
	err = utilities.CopyFile(coverImagePath, filepath.Join(outputDirectory, utilities.GetFileNameFromPath(coverImagePath)))

	var fullSizeImages []string
	files, err := filepath.Glob(path + string(filepath.Separator) + "*")
	if err != nil {
		log.Fatalf("Error globbing gallery %s: %s\n", path, err)
	}

	for _, file := range files {
		for _, extension := range internal.ImageFileExtensions {
			if strings.HasSuffix(file, extension) && !strings.Contains(file, "_thumb") && !strings.Contains(file, "_cover.") {
				fullSizeImages = append(fullSizeImages, file)
			}
		}
	}

	var galleryImages []structs.GalleryImage
	for _, image := range fullSizeImages {
		exif, err := utilities.ParseImageExif(image)
		if err != nil {
			log.Fatalf("Error parsing gallery image %s: %s\n", image, err)
		}

		thumbnailPath := utilities.AppendToFile(image, "_thumb")
		thumbnailSmallPath := utilities.AppendToFile(image, "_thumb_s")

		for _, file := range []string{image, thumbnailPath, thumbnailSmallPath} {
			err := utilities.CopyFile(file, strings.Replace(file, path, outputDirectory, 1))
			if err != nil {
				log.Fatalf("Error copying gallery image %s: %s\n", file, err)
			}
		}

		galleryImage := structs.GalleryImage{
			Image:           utilities.GetFileNameFromPath(image),
			Thumbnail:       utilities.GetFileNameFromPath(thumbnailPath),
			ThumbnailSmall:  utilities.GetFileNameFromPath(thumbnailSmallPath),
			CreateTimestamp: exif.DateTimeOriginal(),
		}
		galleryImages = append(galleryImages, galleryImage)
	}

	sort.Slice(galleryImages, func(i, j int) bool {
		return galleryImages[i].CreateTimestamp.Before(galleryImages[j].CreateTimestamp)
	})

	gallery := structs.Gallery{
		Path:           path,
		Name:           strings.Replace(cases.Title(language.English).String(name), "-", " ", -1),
		HTMLPath:       "galleries/" + strings.ToLower(name) + "/index.html",
		CoverImagePath: strings.Replace(coverImagePath, "content"+string(filepath.Separator), "", 1),
		//CoverImageAltText: "",
		Images:   galleryImages,
		Metadata: metadata,
	}

	err = utilities.BuildTemplate(filepath.Join("templates", "gallery.html"), filepath.Join(outputDirectory, "index.html"), gallery)
	if err != nil {
		log.Fatalf("Error building template: %s\n", err)
	}

	return gallery, nil
}
