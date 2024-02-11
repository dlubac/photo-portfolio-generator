package utilities

import (
	"dlubac/photo-portfolio-generator/internal/structs"
)

func FormatImageMetadata(metadata structs.ImageMetadata) string {
	return metadata.CameraModel + " / " + metadata.LensModel + " / " + metadata.FocalLength + " / " + metadata.ShutterSpeed + " / " + metadata.Aperture + " / " + metadata.ISO
}
