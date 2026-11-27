"""Build dlubac.com from portfolio.json and the images in content/.

Usage:
  uv run main.py [--config portfolio.json] [--content DIR] [--output DIR]

content/ holds one folder per page, matching the site's URLs:
  content/galleries/<slug>/   full-size images and the gallery cover
  content/archive/<slug>/     full-size images and the post cover

portfolio.json lists every page and, for each, its photos in display order
with their alt text. A photo that is listed but missing, or a page with a
photo that has no alt text, stops the build with a message saying which.

Thumbnails and 1200px versions are made next to the sources when missing
(see generator/images.py). The output folder gets the pages, the referenced
images and the static files; nothing already in it is deleted.
"""

import argparse
import datetime as dt
import json
import os
import shutil
import sys
from glob import glob

from dacite import from_dict

from generator import images
from generator.images import ContentError
from generator.model import Config, Page, Photo, long_date
from generator.pages import Site

STATIC = os.path.join(os.path.dirname(os.path.abspath(__file__)), "static")


def load_page(
	content: str, kind: str, slug: str, cover: str, cover_alt: str, photos, heading, title, date=None
) -> Page:
	folder = os.path.join(content, kind, slug)
	if not os.path.isdir(folder):
		raise ContentError(f"{kind}/{slug}: folder {folder} not found")
	if not photos:
		raise ContentError(f"{kind}/{slug}: no photos listed")
	for name in [cover] + [p.file for p in photos]:
		if not os.path.exists(os.path.join(folder, name)):
			raise ContentError(f"{kind}/{slug}: {name} is listed in portfolio.json but not in {folder}")
	missing_alt = [p.file for p in photos if not p.alt.strip()] + ([] if cover_alt.strip() else ["cover"])
	if missing_alt:
		raise ContentError(f"{kind}/{slug}: alt text missing for {', '.join(missing_alt)}")

	cover_w, cover_h = images.size(os.path.join(folder, cover))
	page = Page(kind, slug, heading, title, cover, cover_w, cover_h, cover_alt, date=date)
	for p in photos:
		full_path = os.path.join(folder, p.file)
		w, h = images.size(full_path)
		thumb = images.thumb_for(folder, p.file)
		tw, th = images.size(os.path.join(folder, thumb))
		mid = images.mid_for(folder, p.file, w)
		page.photos.append(Photo(p.file, thumb, mid, w, h, tw, th, p.alt, images.taken(full_path)))

	derived = ("_thumb", f"_w{images.MID_WIDTH}", "_cover")
	on_disk = {
		os.path.basename(f)
		for f in glob(os.path.join(folder, "*.avif")) + glob(os.path.join(folder, "*.jpg"))
		if not any(t in os.path.basename(f) for t in derived)
	}
	for name in sorted(on_disk - {p.file for p in photos}):
		print(f"  note: {kind}/{slug}/{name} is not in portfolio.json, so it is not on the site")
	return page


def copy_if_changed(src: str, dest: str) -> bool:
	if os.path.exists(dest):
		a, b = os.stat(src), os.stat(dest)
		if a.st_size == b.st_size and int(a.st_mtime) == int(b.st_mtime):
			return False
		if os.path.samefile(src, dest):
			return False
	os.makedirs(os.path.dirname(dest), exist_ok=True)
	shutil.copy2(src, dest)
	return True


def main() -> int:
	parser = argparse.ArgumentParser(description=__doc__.split("\n")[0])
	parser.add_argument("--config", default="portfolio.json")
	parser.add_argument("--content", help="overrides content_dir in the config")
	parser.add_argument("--output", help="overrides output_dir in the config")
	args = parser.parse_args()

	with open(args.config, encoding="utf-8") as f:
		config = from_dict(data_class=Config, data=json.load(f))
	content = args.content or config.content_dir
	out = args.output or config.output_dir

	try:
		galleries = []
		for g in config.galleries:
			print(f"galleries/{g.slug}")
			galleries.append(load_page(content, "galleries", g.slug, g.cover, g.cover_alt, g.photos, g.name, g.name))
		posts = []
		for p in config.archive:
			print(f"archive/{p.slug}")
			date = dt.date.fromisoformat(p.date)
			heading = p.title or long_date(date)
			posts.append(load_page(content, "archive", p.slug, p.cover, p.cover_alt, p.photos, heading, p.title, date))
	except ContentError as err:
		print(f"error: {err}", file=sys.stderr)
		return 1
	if not galleries or not posts:
		print("error: portfolio.json needs at least one gallery and one archive post", file=sys.stderr)
		return 1

	copied = 0
	for page in galleries + posts:
		folder = os.path.join(content, page.kind, page.slug)
		names = {page.cover} | {n for p in page.photos for n in (p.full, p.thumb, p.mid) if n}
		for name in sorted(names):
			copied += copy_if_changed(os.path.join(folder, name), os.path.join(out, page.kind, page.slug, name))
	for src in sorted(glob(os.path.join(STATIC, "**", "*"), recursive=True)):
		if os.path.isfile(src):
			copied += copy_if_changed(src, os.path.join(out, os.path.relpath(src, STATIC)))

	built = Site(config, galleries, posts, out).build()
	print(f"built {built} pages, copied {copied} files into {out}")
	return 0


if __name__ == "__main__":
	sys.exit(main())
