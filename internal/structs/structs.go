package structs

import "time"

type SiteMetadata struct {
	Domain      string
	Title       string
	Description string
}

type Homepage struct {
	PhotoReelPreviews []GalleryImage
	Galleries         []Gallery
	Metadata          SiteMetadata
}

type PhotoReel struct {
	Images   []GalleryImage
	Metadata SiteMetadata
}

type Gallery struct {
	Path           string
	Name           string
	HTMLPath       string
	CoverImagePath string
	Images         []GalleryImage
	Metadata       SiteMetadata
}

type GalleryImage struct {
	Image           string
	Thumbnail       string
	ThumbnailSmall  string
	CreateTimestamp time.Time
}
