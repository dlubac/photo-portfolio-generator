# Photo Portfolio Generator

This application is a static photography portfolio generator. Sites generated from it have a minimal amount of html/css, and no javascript.

## Usage
- Create an `images/` directory in the project root
- Create a directory within `images/` for each gallery
- Add images to gallery directories
    - Images should be `.jpg` files no more than 3000px on the long edge
    - Images should follow the naming convention `<gallery>_n.jpg`, where `n` is an integer that controls the ordering of images within a gallery
- Add `cover.jpg` to each gallery
    - Recommended sizing is 600x840 for half-width portrait images and 1200x600 for full-width landscape images
- Copy `config.yml.example` to `config.yml` and populate it
    - Comments are present in the example that explain specific values
- Run `go run cmd/portfoliogenerator/main.go` or `go build -o portfolioGenerator cmd/portfoliogenerator/main.go && ./portfolioGenerator` to run the generator