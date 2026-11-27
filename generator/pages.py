"""HTML for every page: home, galleries, archive, posts and 404."""

import html
import os

from .model import Config, Page, count, long_date, short_month, span

LIGHTBOX_CSS = '\n\t<link rel="stylesheet" href="/lightgallery.min.css" />'
SHORT_SHEET = 7  # posts up to this many photos get taller rows
ROW_SHEET = 4  # posts up to this many photos sit in one row


def e(s) -> str:
	return html.escape(str(s), quote=True)


def meta_spans(parts) -> str:
	parts = [x for x in parts if x]
	return " ".join(f'<span class="nw">{"· " if k else ""}{e(x)}</span>' for k, x in enumerate(parts))


class Site:
	def __init__(self, config: Config, galleries: list[Page], posts: list[Page], out: str):
		self.c = config
		self.site = f"https://{config.domain}"
		self.G = galleries
		self.P = posts
		self.out = out

	# ---------- shared chrome ----------

	def head(self, title, desc, path, image=None, extra=""):
		g0 = self.G[0]
		img = image or f"{g0.url}{g0.cover}"
		name = self.c.name
		return f"""<!doctype html>
<html lang="en">
<head>
	<meta charset="utf-8" />
	<meta name="viewport" content="width=device-width, initial-scale=1" />
	<title>{e(title)}</title>
	<meta name="description" content="{e(desc)}" />
	<link rel="canonical" href="{self.site}{path}" />
	<meta property="og:type" content="website" />
	<meta property="og:site_name" content="{name}" />
	<meta property="og:title" content="{e(title)}" />
	<meta property="og:description" content="{e(desc)}" />
	<meta property="og:url" content="{self.site}{path}" />
	<meta property="og:image" content="{self.site}{e(img)}" />
	<meta name="twitter:card" content="summary_large_image" />
	<meta name="theme-color" content="#f6f5f1" />
	<link rel="icon" href="/favicon.ico" />
	<link rel="preload" href="/fonts/iAWriterDuoS-Regular.woff2" as="font" type="font/woff2" crossorigin />{extra}
	<link rel="stylesheet" href="/styles.css" />
</head>
"""

	def header(self, active):
		def item(label, href, key):
			cur = ' aria-current="page"' if key == active else ""
			return f'<a href="{href}"{cur}>{label}</a>'

		return f"""<body>
<a class="skip" href="#main">Skip to photographs</a>
<header class="site-header">
	<a class="site-name" href="/">{self.c.name}</a>
	<nav aria-label="Site">
		{item("Galleries", "/", "galleries")}
		{item("Archive", "/archive/", "archive")}
		<a href="{self.c.instagram}" target="_blank" rel="noopener">Instagram<span class="sr"> (opens in a new tab)</span></a>
	</nav>
</header>
"""

	def footer(self, scripts=False):
		js = ""
		if scripts:
			js = '\n<script src="/lightgallery.min.js" defer></script>\n<script src="/site.js" defer></script>'
		return f"""<footer class="site-footer">
	<p>{e(self.c.description)}</p>
	<p><a href="{self.c.instagram}" target="_blank" rel="noopener">Instagram<span class="sr"> (opens in a new tab)</span></a></p>
</footer>{js}
</body>
</html>
"""

	def write(self, rel, text):
		path = os.path.join(self.out, rel)
		os.makedirs(os.path.dirname(path), exist_ok=True)
		with open(path, "w", encoding="utf-8") as f:
			f.write(text)

	# ---------- components ----------

	def gallery_card(self, g: Page, i: int):
		load = 'loading="lazy"'
		if i < 3:
			load = 'loading="eager"' + (' fetchpriority="high"' if i == 0 else "")
		tid = f"g-{g.slug}"
		return f"""		<li class="card">
			<a href="{g.url}" aria-labelledby="{tid}">
				<img src="{g.url}{e(g.cover)}" width="{g.cover_w}" height="{g.cover_h}" alt="{e(g.cover_alt)}" {load} decoding="async" />
				<div class="card-text">
					<h3 class="card-title" id="{tid}">{e(g.title)}</h3>
					<span class="meta">{meta_spans([count(g.n)])}</span>
				</div>
			</a>
		</li>
"""

	def post_card(self, p: Page, tag="h2"):
		tid = f"p-{p.slug}"
		meta = meta_spans([long_date(p.date) if p.title else "", count(p.n)])
		return f"""		<li class="card card--post">
			<a href="{p.url}" aria-labelledby="{tid}">
				<img src="{p.url}{e(p.cover)}" width="{p.cover_w}" height="{p.cover_h}" alt="{e(p.cover_alt)}" loading="lazy" decoding="async" />
				<div class="card-text">
					<{tag} class="card-title" id="{tid}">{e(p.heading)}</{tag}>
					<span class="meta">{meta}</span>
				</div>
			</a>
		</li>
"""

	def sheet(self, g: Page, place: str):
		mode = "row" if g.n <= ROW_SHEET else ("short" if g.n <= SHORT_SHEET else "")
		out = [f'\t<ol class="sheet{" sheet--" + mode if mode else ""}" id="sheet">']
		row_max = 500 if mode == "short" else 340
		sum_r = sum(p.tw / p.th for p in g.photos)
		show_dates = len({(p.taken.year, p.taken.month) for p in g.photos if p.taken}) > 1
		for i, p in enumerate(g.photos):
			when = short_month(p.taken) if p.taken else ""
			sub = f"<p class='lg-meta'>{e(' · '.join(x for x in [place, when] if x))}</p>"
			ratio = round(p.tw / p.th, 4)
			load = 'loading="eager"' if i < 3 else 'loading="lazy"'
			cands = [(p.thumb, p.tw)] + ([(p.mid, 1200)] if p.mid else []) + [(p.full, p.w)]
			srcset = ", ".join(f"{e(f)} {w}w" for f, w in cands)
			lb_srcset = ", ".join(f"{e(f)} {w}w" for f, w in cands[1:])
			if mode == "row":
				desk = round(min(1264 * ratio / sum_r, ratio * 760))
			else:
				desk = round(min(ratio * row_max * 1.75, 1264))
			sizes = f"(max-width: 640px) calc(100vw - 32px), {desk}px"
			date_cap = f"<span>{e(when)}</span>" if when and show_dates else ""
			out.append(f"""		<li class="frame" style="--r:{ratio}">
			<a class="frame-link" href="{e(p.full)}" data-srcset="{lb_srcset}" data-sizes="100vw" data-sub-html="{e(sub)}">
				<img src="{e(p.thumb)}" srcset="{srcset}" sizes="{sizes}" width="{p.tw}" height="{p.th}" alt="{e(p.alt)}" {load} decoding="async" />
			</a>
			<p class="frame-cap meta" aria-hidden="true"><span>{i + 1:02d}</span>{date_cap}</p>
		</li>""")
		out.append("\t</ol>")
		return "\n".join(out) + "\n"

	# ---------- pages ----------

	def home(self):
		cards = "".join(self.gallery_card(g, i) for i, g in enumerate(self.G))
		recent = "".join(self.post_card(p, "h3") for p in self.P[:3])
		body = f"""<main id="main" class="wrap">
	<h1 class="sr">{self.c.name}</h1>
	<section aria-labelledby="galleries-h" class="home-galleries">
		<h2 class="section-title" id="galleries-h">Galleries</h2>
		<ol class="cards cards--galleries">
{cards}		</ol>
	</section>
	<section aria-labelledby="recent-h" class="recent">
		<h2 class="section-title" id="recent-h">Archive <a class="meta" href="/archive/">All {len(self.P)} posts</a></h2>
		<ol class="cards cards--posts">
{recent}		</ol>
	</section>
</main>
"""
		self.write(
			"index.html",
			self.head(f"{self.c.name} — Photography", self.c.description, "/")
			+ self.header("galleries")
			+ body
			+ self.footer(),
		)

	def gallery_page(self, i: int, g: Page):
		nxt = self.G[(i + 1) % len(self.G)]
		when = span(p.taken for p in g.photos)
		meta = " · ".join(x for x in [count(g.n), when] if x)
		body = f"""<main id="main" class="wrap">
	<div class="page-head">
		<h1>{e(g.title)}</h1>
		<p class="meta">{e(meta)}</p>
	</div>
{self.sheet(g, g.title)}
	<nav class="next" aria-label="Next gallery">
		<a class="next-link" href="{nxt.url}">
			<img src="{nxt.url}{e(nxt.cover)}" width="{nxt.cover_w}" height="{nxt.cover_h}" alt="" loading="lazy" decoding="async" />
			<span>
				<span class="next-title">{e(nxt.title)}</span>
				<span class="meta">Next gallery · {e(count(nxt.n))}</span>
			</span>
		</a>
		<a class="back" href="/">All galleries</a>
	</nav>
</main>
"""
		desc = f"{g.title}: {count(g.n)} by {self.c.name}" + (f", {when}." if when else ".")
		self.write(
			f"galleries/{g.slug}/index.html",
			self.head(f"{g.title} — {self.c.name}", desc, g.url, g.url + g.cover, LIGHTBOX_CSS)
			+ self.header("galleries")
			+ body
			+ self.footer(scripts=True),
		)

	def archive_index(self):
		dates = [p.date for p in self.P]
		when = f"{short_month(min(dates))} – {short_month(max(dates))}"
		cards = "".join(self.post_card(p) for p in self.P)
		body = f"""<main id="main" class="wrap">
	<div class="page-head">
		<h1>Archive</h1>
		<p class="meta">{len(self.P)} posts · {e(when)}</p>
	</div>
	<ol class="cards cards--posts">
{cards}	</ol>
</main>
"""
		p0 = self.P[0]
		self.write(
			"archive/index.html",
			self.head(
				f"Archive — {self.c.name}",
				f"Recent photographs by {self.c.name}, posted {when}.",
				"/archive/",
				p0.url + p0.cover,
			)
			+ self.header("archive")
			+ body
			+ self.footer(),
		)

	def post_page(self, i: int, p: Page):
		newer = self.P[i - 1] if i > 0 else None
		older = self.P[i + 1] if i + 1 < len(self.P) else None
		shown = long_date(p.date)
		meta = " · ".join(x for x in [shown if p.title else "", count(p.n)] if x)
		place = p.title or shown

		def pn(q, label, cls):
			if not q:
				return f'<span class="{cls}"></span>'
			return (
				f'<a class="{cls}" href="{q.url}"><span class="meta">{label}</span>'
				f'<span class="next-title">{e(q.heading)}</span></a>'
			)

		body = f"""<main id="main" class="wrap">
	<div class="page-head">
		<h1>{e(p.heading)}</h1>
		<p class="meta">{e(meta)}</p>
	</div>
{self.sheet(p, place)}
	<nav class="pager" aria-label="More posts">
		{pn(older, "Older", "pager-older")}
		<a class="back" href="/archive/">All posts</a>
		{pn(newer, "Newer", "pager-newer")}
	</nav>
</main>
"""
		desc = f"{place}: {count(p.n)} by {self.c.name}, posted {shown}."
		self.write(
			f"archive/{p.slug}/index.html",
			self.head(f"{p.heading} — {self.c.name}", desc, p.url, p.url + p.cover, LIGHTBOX_CSS)
			+ self.header("archive")
			+ body
			+ self.footer(scripts=True),
		)

	def not_found(self):
		cards = "".join(self.gallery_card(g, 9) for g in self.G[:3])
		body = f"""<main id="main" class="wrap">
	<div class="page-head">
		<h1>Page not found</h1>
		<p class="lede">There’s nothing at this address.</p>
		<p><a href="/">See all galleries</a> or <a href="/archive/">browse the archive</a>.</p>
	</div>
	<h2 class="section-title">A few galleries <a class="meta" href="/">All {len(self.G)} galleries</a></h2>
	<ol class="cards cards--galleries">
{cards}	</ol>
</main>
"""
		head = self.head(f"Page not found — {self.c.name}", self.c.description, "/404.html").replace(
			'<link rel="canonical"', '<meta name="robots" content="noindex" />\n\t<link rel="canonical"'
		)
		self.write("404.html", head + self.header(None) + body + self.footer())

	def build(self) -> int:
		self.home()
		for i, g in enumerate(self.G):
			self.gallery_page(i, g)
		self.archive_index()
		for i, p in enumerate(self.P):
			self.post_page(i, p)
		self.not_found()
		return 3 + len(self.G) + len(self.P)
