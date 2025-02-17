package steps

import (
	"dlubac/photo-portfolio-generator/internal/utilities"
	"path/filepath"
)

func CopyStaticFiles() error {
	files := []string{"styles.css", "favicon.ico", "iAWriterDuoS-Regular.woff2", "spotlight.bundle.js"}

	for _, file := range files {
		err := utilities.CopyFile(filepath.Join("templates", file), filepath.Join("output", file))
		if err != nil {
			return err
		}
	}

	return nil

}
