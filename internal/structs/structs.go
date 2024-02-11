package structs

type Config struct {
	OutputDirectory    string               `yaml:"outputDirectory"`
	Domain             string               `yaml:"domain"`
	SiteTitle          string               `yaml:"siteTitle"`
	SiteDescription    string               `yaml:"siteDescription"`
	Galleries          []Gallery            `yaml:"galleries"`
	FooterLinks        []FooterLink         `yaml:"footerLinks"`
	CameraModelMapping []CameraModelMapping `yaml:"cameraModelMapping"`
}

type Gallery struct {
	Name        string
	Path        string `yaml:"path"`
	Description string `yaml:"description"`
	Class       string `yaml:"class"`
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
}

type CameraModelMapping struct {
	Model string `yaml:"model"`
	Name  string `yaml:"name"`
}

type SortedImage struct {
	Name  string
	Index int
}
