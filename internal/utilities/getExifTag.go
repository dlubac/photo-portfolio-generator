package utilities

import (
	"os"

	"github.com/rwcarlsen/goexif/exif"
)

func GetExifTag(imagePath string, tagName string) (string, error) {
	img, err := os.Open(imagePath)
	if err != nil {
		return "", err
	}

	exifData, err := exif.Decode(img)
	if err != nil {
		return "", err
	}

	tag, err := exifData.Get(exif.FieldName(tagName))
	if err != nil {
		return "", err
	}

	return tag.String(), nil
}
