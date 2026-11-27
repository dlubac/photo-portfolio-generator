"""Reading image facts and deriving the smaller versions the pages use.

Lightroom exports the full-size image and the cover. Everything else is made here,
once, next to the source in content/, so later runs reuse it:

  <name>_thumb.avif   750px on the long edge (the grid's base size)
  <name>_w1200.avif   1200px wide, for sharp grids on high-density screens and
                      for the lightbox on phones. Only made when the full image
                      is wider than MID_MIN_WIDTH.

Older exports named thumbnails <name>_thumbnail.avif; those are used as they are.

Resizing uses macOS sips rather than Pillow: Lightroom marks Display P3 AVIFs with
colour tags instead of an ICC profile, and Pillow drops those tags, so its output
would be read as sRGB and look duller. sips keeps the colour space.
"""

import datetime as dt
import os
import re
import shutil
import subprocess

from PIL import Image

THUMB_LONG_EDGE = 750
THUMB_QUALITY = 80
MID_WIDTH = 1200
MID_MIN_WIDTH = 1300
MID_QUALITY = 75

EXIF_IFD = 0x8769
DATE_TIME_ORIGINAL = 36867


class ContentError(Exception):
	pass


def base_name(full: str) -> str:
	"""photo_full.avif -> photo; photo.avif -> photo"""
	stem = full.rsplit(".", 1)[0]
	return re.sub(r"_full$", "", stem)


def size(path: str) -> tuple[int, int]:
	with Image.open(path) as img:
		return img.size


def taken(path: str) -> dt.date | None:
	with Image.open(path) as img:
		value = img.getexif().get_ifd(EXIF_IFD).get(DATE_TIME_ORIGINAL)
	if not value:
		return None
	try:
		return dt.datetime.strptime(str(value)[:10], "%Y:%m:%d").date()
	except ValueError:
		return None


def _resized(src: str, dest: str, resize: list[str], quality: int) -> None:
	if not shutil.which("sips"):
		raise ContentError(f"{dest} needs to be made, and that needs macOS sips")
	cmd = ["sips", *resize, "-s", "format", "avif", "-s", "formatOptions", str(quality), src, "--out", dest]
	result = subprocess.run(cmd, capture_output=True, text=True)
	if result.returncode != 0 or not os.path.exists(dest):
		raise ContentError(f"could not make {dest}: {result.stderr.strip() or result.stdout.strip()}")
	print(f"  made {dest}")


def thumb_for(folder: str, full: str) -> str:
	base = base_name(full)
	ext = full.rsplit(".", 1)[1]
	for candidate in (f"{base}_thumb.{ext}", f"{base}_thumbnail.{ext}"):
		if os.path.exists(os.path.join(folder, candidate)):
			return candidate
	made = f"{base}_thumb.avif"
	_resized(os.path.join(folder, full), os.path.join(folder, made), ["-Z", str(THUMB_LONG_EDGE)], THUMB_QUALITY)
	return made


def mid_for(folder: str, full: str, width: int) -> str | None:
	if width <= MID_MIN_WIDTH:
		return None
	made = f"{full.rsplit('.', 1)[0]}_w{MID_WIDTH}.avif"
	if not os.path.exists(os.path.join(folder, made)):
		_resized(
			os.path.join(folder, full), os.path.join(folder, made), ["--resampleWidth", str(MID_WIDTH)], MID_QUALITY
		)
	return made
