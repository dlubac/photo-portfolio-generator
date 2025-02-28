package steps

import (
	"dlubac/photo-portfolio-generator/internal/structs"
	"dlubac/photo-portfolio-generator/internal/utilities"
	"log"
	"path/filepath"
	"sort"
	"strings"
)

func BuildPhotoReel(metadata structs.SiteMetadata) ([]structs.GalleryImage, error) {
	log.Println("Building photo reel")

	err := utilities.CreateDirectory(filepath.Join("output", "photo-reel"))
	if err != nil {
		log.Fatalf("Error creating photo reel output directory: %v", err)
	}

	files, err := filepath.Glob(filepath.Join("content", "photo-reel", "*"))
	if err != nil || files == nil {
		log.Fatalf("Error finding photo reel files: %v", err)
	}

	imagePaths := utilities.FilterImages(files, []string{"_thumb", "_cover."})

	var photoReelImages []structs.GalleryImage
	for _, path := range imagePaths {
		exif, err := utilities.ParseImageExif(path)
		if err != nil {
			continue
		}

		thumbnailPath := utilities.AppendToFile(path, "_thumb")

		for _, file := range []string{path, thumbnailPath} {
			err := utilities.CopyFile(file, strings.Replace(file, "content", "output", 1))
			if err != nil {
				return nil, err
			}
		}

		photoReelImages = append(photoReelImages, structs.GalleryImage{
			Image:           utilities.GetFileNameFromPath(path),
			Thumbnail:       utilities.GetFileNameFromPath(thumbnailPath),
			CreateTimestamp: exif.DateTimeOriginal()})

	}

	sort.Slice(photoReelImages, func(i, j int) bool {
		return photoReelImages[i].CreateTimestamp.After(photoReelImages[j].CreateTimestamp)
	})

	err = utilities.BuildTemplate(
		filepath.Join("templates", "photo-reel.html"),
		filepath.Join("output", "photo-reel", "index.html"),
		structs.PhotoReel{Images: photoReelImages, Metadata: metadata})
	if err != nil {
		log.Fatalf("Error building template: %s\n", err)
	}

	return photoReelImages[:3], nil
}
