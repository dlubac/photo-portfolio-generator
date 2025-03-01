package utilities

import (
	"dlubac/photo-portfolio-generator/internal/structs"
	"sort"
)

func SortImages(images []structs.GalleryImage, oldestFirst bool) {
	if oldestFirst {
		sort.Slice(images, func(i, j int) bool {
			return images[i].CreateTimestamp.Before(images[j].CreateTimestamp)
		})
		return
	}

	sort.Slice(images, func(i, j int) bool {
		return images[i].CreateTimestamp.After(images[j].CreateTimestamp)
	})
}
