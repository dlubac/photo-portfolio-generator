package structs

import "time"

type SiteMetadata struct {
	Domain      string
	Title       string
	Description string
}

type Config struct {
	OutputDirectory    string               `yaml:"outputDirectory"`
	Domain             string               `yaml:"domain"`
	SiteTitle          string               `yaml:"siteTitle"`
	SiteDescription    string               `yaml:"siteDescription"`
	Galleries          []Gallery            `yaml:"galleries"`
	FooterLinks        []FooterLink         `yaml:"footerLinks"`
	CameraModelMapping []CameraModelMapping `yaml:"cameraModelMapping"`
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
	//CoverImageAltText string
	Images   []GalleryImage
	Metadata SiteMetadata
}

type GalleryImage struct {
	Image           string
	Thumbnail       string
	ThumbnailSmall  string
	CreateTimestamp time.Time
}

type HomePageGallery struct {
	Gallery           Gallery
	IndexPath         string
	CoverImagePath    string
	CoverImageAltText string
}

type FooterLink struct {
	Name string `yaml:"name"`
	Url  string `yaml:"url"`
}

type GalleryPage struct {
	PageTitle       string
	PageDescription string
	Url             string
	Domain          string
	Images          []GalleryPageImage
	FooterLinks     []FooterLink
}

type GalleryPageImage struct {
	Url      string
	Metadata string
	AltText  string
	Width    int64
	Height   int64
}

type HomePage struct {
	PageTitle         string
	PageDescription   string
	Domain            string
	HomePageGalleries []HomePageGallery
	FooterLinks       []FooterLink
}

type ImageMetadata struct {
	CameraModel  string
	LensModel    string
	FocalLength  string
	ShutterSpeed string
	Aperture     string
	ISO          string
	Width        int64
	Height       int64
}

type CameraModelMapping struct {
	Model string `yaml:"model"`
	Name  string `yaml:"name"`
}

type SortedImage struct {
	Name  string
	Index int
}
