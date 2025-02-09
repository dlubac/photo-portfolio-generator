package utilities

import (
	"fmt"
	"github.com/evanoberholster/imagemeta"
	"github.com/evanoberholster/imagemeta/exif2"
	"os"
	"strings"
)

func ParseImageExif(path string) (exif2.Exif, error) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	var exif exif2.Exif

	if strings.HasSuffix(path, ".avif") {
		exif, err = imagemeta.DecodeHeif(f)
		if err != nil {
			return exif2.Exif{}, err
		}
	}

	if strings.HasSuffix(path, ".jpg") || strings.HasSuffix(path, ".jpeg") {
		exif, err = imagemeta.DecodeJPEG(f)
		if err != nil {
			return exif2.Exif{}, err
		}
	}

	return exif, nil
}
