package steps

import (
	"dlubac/photo-portfolio-generator/internal/structs"
	"dlubac/photo-portfolio-generator/internal/utilities"
	"log"
	"path/filepath"
	"strings"
)

func buildPhotoReelImage(image string) structs.GalleryImage {
	exif, err := utilities.ParseImageExif(image)
	if err != nil {
		log.Fatalf("Error parsing photo reel image: %v", err)
	}

	thumbnailPath := utilities.AppendToFile(image, "_thumb")

	for _, file := range []string{image, thumbnailPath} {
		err := utilities.CopyFile(file, strings.Replace(file, "content", "output", 1))
		if err != nil {
			log.Fatalf("Error copying image: %v", err)
		}
	}

	return structs.GalleryImage{
		Image:           utilities.GetFileNameFromPath(image),
		Thumbnail:       utilities.GetFileNameFromPath(thumbnailPath),
		CreateTimestamp: exif.DateTimeOriginal()}
}

func BuildPhotoReel(metadata structs.SiteMetadata) []structs.GalleryImage {
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
		photoReelImages = append(photoReelImages, buildPhotoReelImage(path))
	}

	if len(photoReelImages) < 3 {
		log.Fatalf("Photo reel must contain at least 3 images.")
	}

	utilities.SortImages(photoReelImages, false)
	utilities.BuildTemplate(
		filepath.Join("templates", "photo-reel.html"),
		filepath.Join("output", "photo-reel", "index.html"),
		structs.PhotoReel{Images: photoReelImages, Metadata: metadata})

	return photoReelImages[:3]
}
