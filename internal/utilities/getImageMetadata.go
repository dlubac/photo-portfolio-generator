package utilities

import (
	"dlubac/photo-portfolio-generator/internal/structs"
	"strconv"
	"strings"
)

func GetImageMetadata(imagePath string, config *structs.Config) (structs.ImageMetadata, error) {
	focalLength, err := GetExifTag(imagePath, "FocalLengthIn35mmFilm")
	if err == nil {
		focalLength = focalLength + "mm"
	}

	shutterSpeed, err := GetExifTag(imagePath, "ExposureTime")
	if err == nil {
		numerator, denominator := strings.Split(TrimQuotes(shutterSpeed), "/")[0], strings.Split(TrimQuotes(shutterSpeed), "/")[1]
		if numerator >= denominator {
			shutterSpeed = numerator + "s"
		} else {
			shutterSpeed = "1/" + denominator + "s"
		}
	}

	lensModel, err := GetExifTag(imagePath, "LensModel")
	if err == nil {
		lensModel = TrimQuotes(lensModel)
	}

	cameraModel, err := GetExifTag(imagePath, "Model")
	if err == nil {
		mapping := config.CameraModelMapping
		for _, mapping := range mapping {
			if mapping.Model == TrimQuotes(cameraModel) {
				cameraModel = mapping.Name
				break
			}
		}
	}

	aperture, err := GetExifTag(imagePath, "FNumber")
	if err == nil {
		cleanAperture := TrimQuotes(aperture)
		numerator, _ := strconv.Atoi(strings.Split(cleanAperture, "/")[0])
		denominator, _ := strconv.Atoi(strings.Split(cleanAperture, "/")[1])

		fStop := float64(numerator) / float64(denominator)
		aperture = "f/" + strings.Replace(strconv.FormatFloat(fStop, 'f', 1, 64), ".0", "", 1)
	}

	iso, err := GetExifTag(imagePath, "ISOSpeedRatings")
	if err == nil {
		iso = "ISO " + iso
	}

	return structs.ImageMetadata{
		CameraModel:  cameraModel,
		LensModel:    lensModel,
		FocalLength:  focalLength,
		ShutterSpeed: shutterSpeed,
		Aperture:     aperture,
		ISO:          iso,
	}, nil
}
