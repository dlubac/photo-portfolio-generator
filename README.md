# Photo Portfolio Generator

This utility creates a minimal photography portfolio site. It has three kinds of pages:
- **Gallery**: a collection of images shown in chronological order
- **Photo Reel**: an Instagram-esque page showing images in a reverse-chronological grid
- **Home Page**: shows a preview of the first three images in the photo reel and a listing of galleries

## Image Format/Sizing Recommendations
[Spotlight](https://github.com/nextapps-de/spotlight) is used to show full-screen images and thumbnail previews.
All images should be .jpg or .avif and have a creation timestamp in EXIF data.

Lightroom Classic templates for all image formats can be found in [lightroom_templates](lightroom_templates/)

### Photo Reel
Full-size and thumbnail are variants are needed for photo reel images.

- **Full Size**: 2000px on the long edge.
- **Thumbnail**: 350px on the long edge for landscape oriented images. 350px on the short edge for portrait. Should have a `_thumb.jpg|avif` suffix.

### Gallery
#### Cover Image
Cover images must be a 4x5 ratio and at least 500px on the long edge. The first image in a gallery with `_cover` in its
name will be used as the cover image on the homepage.

#### Gallery Images
Full size, thumbnail, and small thumbnail variants are needed for gallery images.

- **Full Size**: 3000px on the long edge.
- **Thumbnail**: 750px on the long edge. Should have a `_thumb.jpg|avif` suffix.
- **Small Thumbnail**: 350px on the long edge. Should have a `_thumb_s.jpg|avif` suffix.


## Usage
1. Create a `content/` directory in the project root with the below structure. Images do not need to follow a photo1,
photo2, photo3, etc... naming scheme.
```
.
└── content/
    ├── galleries/
    │   ├── gallery1/
    │   │   ├── photo1.jpg
    │   │   ├── photo1_thumb.jpg
    │   │   ├── photo1_thumb_s.jpg
    │   │   ├── photo2.jpg
    │   │   ├── photo2_thumb.jpg
    │   │   ├── photo2_thumb_s.jpg
    │   │   ├── ...
    │   │   └── gallery1_cover.jpg
    │   └── gallery2/
    │       ├── photo1.jpg
    │       ├── photo1_thumb.jpg
    │       ├── photo1_thumb_s.jpg
    │       ├── photo2.jpg
    │       ├── photo2_thumb.jpg
    │       ├── photo2_thumb_s.jpg
    │       ├── ...
    │       └── gallery2_cover.jpg
    └── photo-reel/
        ├── photo1.jpg
        ├── photo1_thumb.jpg
        ├── photo2.jpg
        ├── photo2_thumb.jpg
        ├── photo3.jpg
        └── photo3_thumb.jpg
```
2. Run `./run.sh build && ./portfolioGenerator` to run the generator. It is designed to crash if images do not meet
naming or format conventions.
3. Generated site will be stored in `/output`

