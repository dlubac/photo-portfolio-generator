# Photo Portfolio Generator

Builds [dlubac.com](https://dlubac.com): a homepage of galleries, one page per gallery, a dated archive
of short posts, and a 404 page. Photos open in a lightbox (lightGallery). Requires macOS (image
resizing uses `sips`) and [uv](https://docs.astral.sh/uv/).

```
uv run main.py                                  # content/ -> output/
uv run main.py --output ~/dev/repos/portfolio   # straight into the site repo
```

The build stops with a message if a listed photo is missing or has no alt text. Photos that are in
`content/` but not in `portfolio.json` are reported and left off the site. Nothing in the output
folder is deleted.

## Content

```
content/
├── galleries/<slug>/      full-size photos + <name>_cover.avif
└── archive/<slug>/        full-size photos + <name>_cover.avif   (slug is usually the date)
```

`portfolio.json` lists every page, in order, with each page's photos in display order:

```jsonc
{
	"galleries": [
		{
			"name": "Iceland",
			"slug": "iceland",
			"cover": "iceland_cover.avif",
			"cover_alt": "Striped rhyolite mountains with patches of snow under a broken sky",
			"photos": [{ "file": "_7R51233-Enhanced-NR.avif", "alt": "A waterfall pouring over…" }]
		}
	],
	"archive": [
		// newest first; "title" is optional, untitled posts are headed by their date
		{ "slug": "2026-03-03", "date": "2026-03-03", "title": "Blood Moon", "cover": "…", "cover_alt": "…", "photos": [] }
	]
}
```

Alt text describes what is in the photo, in a short sentence, for people who cannot see it.

Gallery pages show the month range from the photos' EXIF capture dates, and month captions when a
gallery spans more than one month. Posts show the `date` given in `portfolio.json`.

## Images

Export from Lightroom only the **full-size photo** and the **cover**. The generator makes the rest
next to the source, once:

| File | Size | Made by |
| --- | --- | --- |
| `<name>.avif` (or `<name>_full.avif`) | 3000px long edge | Lightroom |
| `<name>_cover.avif` | gallery 1000 × 1250 (4:5), post 1000 × 1000 (1:1) | Lightroom |
| `<name>_thumb.avif` | 750px long edge | generator (an existing `_thumb` or `_thumbnail` file is used as is) |
| `<name>_w1200.avif` | 1200px wide, when the full photo is wider than 1300px | generator |

### Lightroom export settings

| Setting | Full size | Gallery cover | Post cover |
| --- | --- | --- | --- |
| Crop (on a virtual copy) | as edited | 4:5 | 1:1 |
| Image format | AVIF | AVIF | AVIF |
| Quality | 90 | 80 | 80 |
| Color space | Display P3 | Display P3 | Display P3 |
| Bit depth | 10 bits/component | 10 bits/component | 10 bits/component |
| HDR output | off | off | off |
| Resize to fit | Long Edge 3000 px | Long Edge 1250 px | Long Edge 1000 px |
| Don't enlarge | on | on | on |
| Output sharpening | off | Screen, Standard | Screen, Standard |
| Metadata | All, remove location info | All, remove location info | All, remove location info |

Covers are shown up to about 410 CSS pixels wide, so 1000px stays sharp on 2× screens.
