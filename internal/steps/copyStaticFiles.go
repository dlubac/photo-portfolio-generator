package steps

import (
	"dlubac/photo-portfolio-generator/internal/utilities"
	"log"
	"path/filepath"
)

func CopyStaticFiles() {
	files := []string{"styles.css", "favicon.ico", "iAWriterDuoS-Regular.woff2", "spotlight.bundle.js"}

	for _, file := range files {
		err := utilities.CopyFile(filepath.Join("templates", file), filepath.Join("output", file))
		if err != nil {
			log.Fatalf("Error copying static files: %v", err)
		}
	}
}
