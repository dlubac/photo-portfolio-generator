"""Site configuration (portfolio.json) and the page model built from it."""

import datetime as dt
from dataclasses import dataclass, field

MONTHS = [
	"January",
	"February",
	"March",
	"April",
	"May",
	"June",
	"July",
	"August",
	"September",
	"October",
	"November",
	"December",
]


# ---------- portfolio.json ----------


@dataclass
class PhotoConfig:
	file: str  # full-size image, relative to the page's content folder
	alt: str


@dataclass
class GalleryConfig:
	name: str
	slug: str
	cover: str
	cover_alt: str
	photos: list[PhotoConfig]


@dataclass
class PostConfig:
	slug: str  # folder name under content/archive/ and the URL
	date: str  # the date shown on the page, YYYY-MM-DD
	cover: str
	cover_alt: str
	photos: list[PhotoConfig]
	title: str = ""


@dataclass
class Config:
	domain: str
	name: str
	description: str
	instagram: str
	galleries: list[GalleryConfig]
	archive: list[PostConfig]  # newest first
	content_dir: str = "content"
	output_dir: str = "output"


# ---------- built model ----------


@dataclass
class Photo:
	full: str
	thumb: str
	mid: str | None  # 1200px-wide version, when the full image is wide enough to need one
	w: int
	h: int
	tw: int
	th: int
	alt: str
	taken: dt.date | None


@dataclass
class Page:
	kind: str  # "galleries" or "archive"
	slug: str
	heading: str  # gallery name, or post title, or post date
	title: str  # gallery name or post title ("" for untitled posts)
	cover: str
	cover_w: int
	cover_h: int
	cover_alt: str
	photos: list[Photo] = field(default_factory=list)
	date: dt.date | None = None  # posts only

	@property
	def url(self) -> str:
		return f"/{self.kind}/{self.slug}/"

	@property
	def n(self) -> int:
		return len(self.photos)


# ---------- formatting ----------


def short_month(d: dt.date) -> str:
	return f"{MONTHS[d.month - 1][:3]} {d.year}"


def long_date(d: dt.date) -> str:
	return f"{MONTHS[d.month - 1]} {d.day}, {d.year}"


def span(dates) -> str:
	dates = sorted(d for d in dates if d)
	if not dates:
		return ""
	a, b = dates[0], dates[-1]
	if (a.year, a.month) == (b.year, b.month):
		return short_month(a)
	if a.year == b.year:
		return f"{MONTHS[a.month - 1][:3]}–{MONTHS[b.month - 1][:3]} {a.year}"
	return f"{a.year}–{b.year}"


def count(n: int) -> str:
	return f"{n} photograph{'s' if n != 1 else ''}"
