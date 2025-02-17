package steps

import (
	"dlubac/photo-portfolio-generator/internal"
	"dlubac/photo-portfolio-generator/internal/structs"
	"dlubac/photo-portfolio-generator/internal/utilities"
	"errors"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"html/template"
	"log"
	"os"
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
		return structs.Gallery{}, err
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
			return structs.Gallery{}, err
		}

		pathParts := strings.Split(image, ".")
		thumbnailPath := pathParts[0] + "_thumb." + pathParts[1]
		thumbnailSmallPath := pathParts[0] + "_thumb_s." + pathParts[1]

		for _, file := range []string{image, thumbnailPath, thumbnailSmallPath} {
			err := utilities.CopyFile(file, strings.Replace(file, path, outputDirectory, 1))
			if err != nil {
				return structs.Gallery{}, err
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
		Name:           cases.Title(language.English).String(name),
		HTMLPath:       "galleries/" + strings.ToLower(name) + "/index.html",
		CoverImagePath: strings.Replace(coverImagePath, "content"+string(filepath.Separator), "", 1),
		//CoverImageAltText: "",
		Images:   galleryImages,
		Metadata: metadata,
	}

	tmpl, err := template.ParseFiles("templates/gallery.html")
	if err != nil {
		return structs.Gallery{}, err
	}

	galleryIndex, err := os.Create(outputDirectory + "/index.html")
	if err != nil {
		return structs.Gallery{}, err
	}

	err = tmpl.Execute(galleryIndex, gallery)

	return gallery, nil
}
