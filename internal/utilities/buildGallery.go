package utilities

import (
	"dlubac/photo-portfolio-generator/internal"
	"dlubac/photo-portfolio-generator/internal/structs"
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"log"
	"path/filepath"
	"strings"
)

func BuildGallery(path string) (structs.Gallery, error) {
	fmt.Printf("Building gallery %s\n", path)

	coverImages, err := filepath.Glob(path + string(filepath.Separator) + "*_cover.*")
	if err != nil {
		return structs.Gallery{}, err
	}
	coverImagePath := coverImages[0]

	var baseImages []string
	files, err := filepath.Glob(path + string(filepath.Separator) + "*")
	if err != nil {
		return structs.Gallery{}, err
	}

	for _, file := range files {
		for _, extension := range internal.ImageFileExtensions {
			if strings.HasSuffix(file, extension) && !strings.Contains(file, "_thumb") && !strings.Contains(file, "_cover.") {
				baseImages = append(baseImages, file)
			}
		}
	}

	var galleryImages []structs.GalleryImage
	for _, image := range baseImages {
		log.Printf("Processing image: %v", image)
		exif, err := ParseImageExif(image)
		if err != nil {
			return structs.Gallery{}, err
		}

		pathParts := strings.Split(image, ".")
		thumbnailPath := pathParts[0] + "_thumb." + pathParts[1]
		thumbnailSmallPath := pathParts[0] + "_thumb_s." + pathParts[1]

		galleryImage := structs.GalleryImage{
			Path: path, ThumbnailPath: thumbnailPath, ThumbnailSmallPath: thumbnailSmallPath, CreateTimestamp: exif.DateTimeOriginal(),
		}

		galleryImages = append(galleryImages, galleryImage)
	}

	return structs.Gallery{
		Path:              path,
		Name:              cases.Title(language.English).String(strings.Split(path, string(filepath.Separator))[2]),
		HTMLPath:          strings.Replace(path, "content"+string(filepath.Separator), "", 1) + string(filepath.Separator) + "index.html",
		CoverImagePath:    coverImagePath,
		CoverImageAltText: "",
		Images:            galleryImages,
	}, nil
}
